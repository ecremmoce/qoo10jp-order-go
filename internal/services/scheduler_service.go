package services

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"qoo10jp-order-go/internal/models"
	"qoo10jp-order-go/pkg/redis"
	"qoo10jp-order-go/pkg/supabase"
	"time"

	"github.com/google/uuid"
)

const (
	OrderJobQueue        = "shopee_order_queue"
	Qoo10JPOrderJobQueue = "qoo10jp_order_queue"
	SchedulerStateKey    = "qoo10jp_scheduler_state"
	// 기본 수집 간격 (분) - API 호출 실패 시 사용
	DefaultCollectionInterval = 30
)

type SchedulerService struct {
	redisClient    *redis.Client
	supabaseClient *supabase.Client
	orderService   *OrderService
}

func NewSchedulerService(redisClient *redis.Client, supabaseClient *supabase.Client, orderService *OrderService) *SchedulerService {
	return &SchedulerService{
		redisClient:    redisClient,
		supabaseClient: supabaseClient,
		orderService:   orderService,
	}
}

// CreateOrderJob creates a new order collection job and pushes it to Redis queue
func (s *SchedulerService) CreateOrderJob(startDate, endDate time.Time) error {
	job := &models.OrderJob{
		ID:            uuid.New().String(),
		StartDate:     startDate,
		EndDate:       endDate,
		StartDatetime: startDate,
		EndDatetime:   endDate,
		CreatedAt:     time.Now(),
		Priority:      1,
		RetryCount:    0,
		MaxRetries:    3,
	}

	jobJSON, err := job.ToJSON()
	if err != nil {
		return fmt.Errorf("failed to serialize job: %v", err)
	}

	err = s.redisClient.PushToQueue(OrderJobQueue, jobJSON)
	if err != nil {
		return fmt.Errorf("failed to push job to queue: %v", err)
	}

	log.Printf("Created order job %s for period %s to %s",
		job.ID, startDate.Format("2006-01-02"), endDate.Format("2006-01-02"))

	return nil
}

// ProcessNextJob processes the next job from the queue
func (s *SchedulerService) ProcessNextJob() error {
	jobData, err := s.redisClient.PopFromQueue(OrderJobQueue)
	if err != nil {
		return fmt.Errorf("failed to pop job from queue: %v", err)
	}

	if jobData == "" {
		return nil // No jobs in queue
	}

	job, err := models.OrderJobFromJSON(jobData)
	if err != nil {
		return fmt.Errorf("failed to deserialize job: %v", err)
	}

	return s.executeJob(job)
}

// executeJob executes a single order collection job
func (s *SchedulerService) executeJob(job *models.OrderJob) error {
	startTime := time.Now()

	result := &models.JobResult{
		JobID:     job.ID,
		StartTime: startTime,
		CreatedAt: time.Now(),
	}

	log.Printf("Starting job %s: collecting orders from %s to %s",
		job.ID, job.StartDate.Format("2006-01-02"), job.EndDate.Format("2006-01-02"))

	// Execute the order collection
	err := s.orderService.CollectOrders(job.StartDate, job.EndDate)

	endTime := time.Now()
	result.EndTime = endTime
	result.Duration = endTime.Sub(startTime).Milliseconds()

	if err != nil {
		result.Success = false
		result.ErrorMsg = err.Error()
		log.Printf("Job %s failed: %v", job.ID, err)
	} else {
		result.Success = true
		log.Printf("Job %s completed successfully in %dms", job.ID, result.Duration)
	}

	// Save job result to Supabase
	err = s.saveJobResult(result)
	if err != nil {
		log.Printf("Failed to save job result: %v", err)
	}

	// Update scheduler state
	err = s.updateSchedulerState(endTime)
	if err != nil {
		log.Printf("Failed to update scheduler state: %v", err)
	}

	return nil
}

// saveJobResult saves the job execution result to Supabase
func (s *SchedulerService) saveJobResult(result *models.JobResult) error {
	return s.supabaseClient.Insert("job_results", result)
}

// updateSchedulerState updates the scheduler state with the last execution time
func (s *SchedulerService) updateSchedulerState(lastExecution time.Time) error {
	// 동적으로 수집 간격 가져오기
	intervalMinutes := s.GetDynamicCollectionInterval()

	state := &models.SchedulerState{
		LastExecutionTime: lastExecution,
		NextScheduleTime:  lastExecution.Add(time.Duration(intervalMinutes) * time.Minute),
		IsRunning:         false,
		UpdatedAt:         time.Now(),
	}

	stateJSON, err := json.Marshal(state)
	if err != nil {
		return err
	}

	log.Printf("[Qoo10JP Scheduler] Next execution in %d minutes at %s",
		intervalMinutes, state.NextScheduleTime.Format("15:04:05"))

	return s.redisClient.Set(SchedulerStateKey, string(stateJSON), 0)
}

// GetDynamicCollectionInterval fetches the current collection interval from API
func (s *SchedulerService) GetDynamicCollectionInterval() int {
	apiURL := os.Getenv("PLATFORM_API_URL")
	apiKey := os.Getenv("PLATFORM_API_KEY")

	if apiURL == "" {
		log.Printf("[Qoo10JP Scheduler] PLATFORM_API_URL not set, using default interval: %d minutes", DefaultCollectionInterval)
		return DefaultCollectionInterval
	}

	// API 호출하여 현재 수집 간격 조회
	req, err := http.NewRequest("GET", apiURL+"/api/v1/fulfillment/schedules", nil)
	if err != nil {
		log.Printf("[Qoo10JP Scheduler] Failed to create request: %v", err)
		return DefaultCollectionInterval
	}

	if apiKey != "" {
		req.Header.Set("X-API-Key", apiKey)
	}

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("[Qoo10JP Scheduler] Failed to fetch schedule: %v", err)
		return DefaultCollectionInterval
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("[Qoo10JP Scheduler] API returned status %d, using default interval", resp.StatusCode)
		return DefaultCollectionInterval
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("[Qoo10JP Scheduler] Failed to read response: %v", err)
		return DefaultCollectionInterval
	}

	var response struct {
		Success bool `json:"success"`
		Data    struct {
			CurrentIntervalMinutes int `json:"current_interval_minutes"`
		} `json:"data"`
	}

	if err := json.Unmarshal(body, &response); err != nil {
		log.Printf("[Qoo10JP Scheduler] Failed to parse response: %v", err)
		return DefaultCollectionInterval
	}

	if !response.Success || response.Data.CurrentIntervalMinutes <= 0 {
		return DefaultCollectionInterval
	}

	log.Printf("[Qoo10JP Scheduler] Dynamic interval from API: %d minutes", response.Data.CurrentIntervalMinutes)
	return response.Data.CurrentIntervalMinutes
}

// GetSchedulerState retrieves the current scheduler state
func (s *SchedulerService) GetSchedulerState() (*models.SchedulerState, error) {
	stateData, err := s.redisClient.Get(SchedulerStateKey)
	if err != nil {
		// If no state exists, return default state
		return &models.SchedulerState{
			LastExecutionTime: time.Now().Add(-5 * time.Minute),
			NextScheduleTime:  time.Now(),
			IsRunning:         false,
			UpdatedAt:         time.Now(),
		}, nil
	}

	var state models.SchedulerState
	err = json.Unmarshal([]byte(stateData), &state)
	if err != nil {
		return nil, err
	}

	return &state, nil
}

// GetQueueStatus returns the current queue status
func (s *SchedulerService) GetQueueStatus() (int64, error) {
	return s.redisClient.GetQueueLength(OrderJobQueue)
}

func (s *SchedulerService) GetRedisClient() *redis.Client {
	return s.redisClient
}

// ScheduleNextJob creates the next job based on the last execution time
func (s *SchedulerService) ScheduleNextJob() error {
	state, err := s.GetSchedulerState()
	if err != nil {
		return err
	}

	// Calculate the date range for the next job
	// Use last execution time to determine what orders to fetch
	startDate := state.LastExecutionTime.Add(-1 * time.Hour) // Overlap by 1 hour to catch any missed orders
	endDate := time.Now()

	return s.CreateOrderJob(startDate, endDate)
}
