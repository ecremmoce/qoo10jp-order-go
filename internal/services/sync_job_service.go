package services

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"qoo10jp-order-go/internal/models"
	"qoo10jp-order-go/pkg/supabase"

	"github.com/google/uuid"
)

// SyncJobService manages synchronization jobs
type SyncJobService struct {
	supabaseClient *supabase.Client
}

// NewSyncJobService creates a new SyncJobService instance
func NewSyncJobService(supabaseClient *supabase.Client) *SyncJobService {
	return &SyncJobService{
		supabaseClient: supabaseClient,
	}
}

// ========================================
// Job Creation
// ========================================

// CreateJob creates a new sync job
func (s *SyncJobService) CreateJob(sellerID string, startDate, endDate time.Time, jobType string) (*models.Qoo10JPSyncJob, error) {
	if jobType == "" {
		jobType = "order_collection"
	}

	job := &models.Qoo10JPSyncJob{
		ID:                 uuid.New().String(),
		SellerID:           sellerID,
		JobType:            jobType,
		StartDate:          &startDate,
		EndDate:            &endDate,
		Status:             models.SyncJobStatusPending,
		TotalCollected:     0,
		TotalSaved:         0,
		TotalUpdated:       0,
		TotalSkipped:       0,
		TotalFailed:        0,
		ProgressPercentage: 0,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}

	// Supabase에 저장
	err := s.supabaseClient.Insert("qoo10jp_sync_jobs", job)
	if err != nil {
		return nil, fmt.Errorf("failed to create sync job: %w", err)
	}

	log.Printf("📝 SyncJob 생성: %s (seller: %s, type: %s)", job.ID, sellerID, jobType)
	return job, nil
}

// ========================================
// Job Status Management
// ========================================

// CanStartCollection checks if a new collection job can start
// Returns false if there's already an active job for this seller
func (s *SyncJobService) CanStartCollection(sellerID string, jobType string) (bool, *models.Qoo10JPSyncJob, error) {
	if jobType == "" {
		jobType = "order_collection"
	}

	// 활성 작업 확인
	var jobs []models.Qoo10JPSyncJob
	query := fmt.Sprintf("seller_id=eq.%s&job_type=eq.%s&status=in.(pending,running)", sellerID, jobType)

	err := s.supabaseClient.Select("qoo10jp_sync_jobs", query, &jobs)
	if err != nil {
		return false, nil, fmt.Errorf("failed to check active jobs: %w", err)
	}

	if len(jobs) > 0 {
		log.Printf("⚠️ 이미 진행 중인 작업 있음: seller=%s, job=%s, status=%s",
			sellerID, jobs[0].ID, jobs[0].Status)
		return false, &jobs[0], nil
	}

	return true, nil, nil
}

// StartJob marks a job as running
func (s *SyncJobService) StartJob(jobID string) error {
	now := time.Now()
	update := map[string]interface{}{
		"status":     models.SyncJobStatusRunning,
		"started_at": now,
		"updated_at": now,
	}

	err := s.supabaseClient.Update("qoo10jp_sync_jobs", fmt.Sprintf("id=eq.%s", jobID), update)
	if err != nil {
		return fmt.Errorf("failed to start job: %w", err)
	}

	log.Printf("▶️ SyncJob 시작: %s", jobID)
	return nil
}

// CompleteJob marks a job as completed
func (s *SyncJobService) CompleteJob(jobID string, result *models.Qoo10JPCollectionResult) error {
	now := time.Now()

	// Duration 계산
	var durationMs int64
	if result != nil {
		durationMs = result.DurationMs
	}

	// 결과 요약 JSON
	var resultSummary json.RawMessage
	if result != nil {
		if data, err := json.Marshal(result); err == nil {
			resultSummary = data
		}
	}

	update := map[string]interface{}{
		"status":              models.SyncJobStatusCompleted,
		"completed_at":        now,
		"updated_at":          now,
		"progress_percentage": 100,
		"duration_ms":         durationMs,
	}

	if result != nil {
		update["total_collected"] = result.TotalCollected
		update["total_saved"] = result.TotalSaved
		update["total_updated"] = result.TotalUpdated
		update["total_skipped"] = result.TotalSkipped
		update["total_failed"] = result.TotalFailed
	}

	if resultSummary != nil {
		update["result_summary"] = resultSummary
	}

	err := s.supabaseClient.Update("qoo10jp_sync_jobs", fmt.Sprintf("id=eq.%s", jobID), update)
	if err != nil {
		return fmt.Errorf("failed to complete job: %w", err)
	}

	log.Printf("✅ SyncJob 완료: %s (collected: %d, saved: %d, duration: %dms)",
		jobID,
		result.TotalCollected,
		result.TotalSaved,
		durationMs)
	return nil
}

// MarkJobFailed marks a job as failed
func (s *SyncJobService) MarkJobFailed(jobID string, errMsg string) error {
	now := time.Now()
	update := map[string]interface{}{
		"status":        models.SyncJobStatusFailed,
		"completed_at":  now,
		"updated_at":    now,
		"error_message": errMsg,
	}

	err := s.supabaseClient.Update("qoo10jp_sync_jobs", fmt.Sprintf("id=eq.%s", jobID), update)
	if err != nil {
		return fmt.Errorf("failed to mark job as failed: %w", err)
	}

	log.Printf("❌ SyncJob 실패: %s - %s", jobID, errMsg)
	return nil
}

// CancelJob marks a job as cancelled
func (s *SyncJobService) CancelJob(jobID string, reason string) error {
	now := time.Now()
	update := map[string]interface{}{
		"status":        models.SyncJobStatusCancelled,
		"completed_at":  now,
		"updated_at":    now,
		"error_message": reason,
	}

	err := s.supabaseClient.Update("qoo10jp_sync_jobs", fmt.Sprintf("id=eq.%s", jobID), update)
	if err != nil {
		return fmt.Errorf("failed to cancel job: %w", err)
	}

	log.Printf("🚫 SyncJob 취소: %s - %s", jobID, reason)
	return nil
}

// ========================================
// Progress Tracking
// ========================================

// UpdateJobProgress updates the progress of a job
func (s *SyncJobService) UpdateJobProgress(jobID string, collected, saved, updated, skipped, failed int, percentage float64) error {
	update := map[string]interface{}{
		"total_collected":     collected,
		"total_saved":         saved,
		"total_updated":       updated,
		"total_skipped":       skipped,
		"total_failed":        failed,
		"progress_percentage": percentage,
		"updated_at":          time.Now(),
	}

	err := s.supabaseClient.Update("qoo10jp_sync_jobs", fmt.Sprintf("id=eq.%s", jobID), update)
	if err != nil {
		return fmt.Errorf("failed to update job progress: %w", err)
	}

	return nil
}

// IncrementCollected increments the collected count
func (s *SyncJobService) IncrementCollected(jobID string, count int) error {
	// 먼저 현재 값을 조회
	var jobs []models.Qoo10JPSyncJob
	err := s.supabaseClient.Select("qoo10jp_sync_jobs", fmt.Sprintf("id=eq.%s", jobID), &jobs)
	if err != nil || len(jobs) == 0 {
		return fmt.Errorf("failed to get job: %w", err)
	}

	update := map[string]interface{}{
		"total_collected": jobs[0].TotalCollected + count,
		"updated_at":      time.Now(),
	}

	return s.supabaseClient.Update("qoo10jp_sync_jobs", fmt.Sprintf("id=eq.%s", jobID), update)
}

// ========================================
// Job Queries
// ========================================

// GetJob retrieves a job by ID
func (s *SyncJobService) GetJob(jobID string) (*models.Qoo10JPSyncJob, error) {
	var jobs []models.Qoo10JPSyncJob
	err := s.supabaseClient.Select("qoo10jp_sync_jobs", fmt.Sprintf("id=eq.%s", jobID), &jobs)
	if err != nil {
		return nil, fmt.Errorf("failed to get job: %w", err)
	}
	if len(jobs) == 0 {
		return nil, fmt.Errorf("job not found: %s", jobID)
	}
	return &jobs[0], nil
}

// GetRecentJobs retrieves recent jobs for a seller
func (s *SyncJobService) GetRecentJobs(sellerID string, limit int) ([]models.Qoo10JPSyncJob, error) {
	if limit <= 0 {
		limit = 10
	}

	var jobs []models.Qoo10JPSyncJob
	query := fmt.Sprintf("seller_id=eq.%s&order=created_at.desc&limit=%d", sellerID, limit)

	err := s.supabaseClient.Select("qoo10jp_sync_jobs", query, &jobs)
	if err != nil {
		return nil, fmt.Errorf("failed to get recent jobs: %w", err)
	}

	return jobs, nil
}

// GetActiveJobs retrieves all active (pending/running) jobs
func (s *SyncJobService) GetActiveJobs() ([]models.Qoo10JPSyncJob, error) {
	var jobs []models.Qoo10JPSyncJob
	query := "status=in.(pending,running)&order=created_at.desc"

	err := s.supabaseClient.Select("qoo10jp_sync_jobs", query, &jobs)
	if err != nil {
		return nil, fmt.Errorf("failed to get active jobs: %w", err)
	}

	return jobs, nil
}

// GetJobStatistics retrieves job statistics for a seller
func (s *SyncJobService) GetJobStatistics(sellerID string, days int) (*SyncJobStatistics, error) {
	if days <= 0 {
		days = 7
	}

	startDate := time.Now().AddDate(0, 0, -days)

	var jobs []models.Qoo10JPSyncJob
	query := fmt.Sprintf("seller_id=eq.%s&created_at=gte.%s&order=created_at.desc",
		sellerID, startDate.Format(time.RFC3339))

	err := s.supabaseClient.Select("qoo10jp_sync_jobs", query, &jobs)
	if err != nil {
		return nil, fmt.Errorf("failed to get job statistics: %w", err)
	}

	stats := &SyncJobStatistics{
		SellerID:   sellerID,
		Period:     days,
		TotalJobs:  len(jobs),
		StartDate:  startDate,
		EndDate:    time.Now(),
	}

	var totalDuration int64
	var completedCount int

	for _, job := range jobs {
		switch job.Status {
		case models.SyncJobStatusCompleted:
			stats.SuccessfulJobs++
			completedCount++
			totalDuration += job.DurationMs
		case models.SyncJobStatusFailed:
			stats.FailedJobs++
		case models.SyncJobStatusCancelled:
			stats.CancelledJobs++
		case models.SyncJobStatusPending, models.SyncJobStatusRunning:
			stats.ActiveJobs++
		}
		stats.TotalCollected += job.TotalCollected
		stats.TotalSaved += job.TotalSaved
	}

	if completedCount > 0 {
		stats.AvgDurationMs = totalDuration / int64(completedCount)
	}

	return stats, nil
}

// ========================================
// Stale Job Cleanup
// ========================================

// CleanupStaleJobs marks running jobs older than timeout as failed
func (s *SyncJobService) CleanupStaleJobs(timeoutMinutes int) (int, error) {
	if timeoutMinutes <= 0 {
		timeoutMinutes = 60 // 기본 1시간
	}

	cutoffTime := time.Now().Add(-time.Duration(timeoutMinutes) * time.Minute)

	// 오래된 실행 중 작업 조회
	var staleJobs []models.Qoo10JPSyncJob
	query := fmt.Sprintf("status=eq.running&started_at=lt.%s", cutoffTime.Format(time.RFC3339))

	err := s.supabaseClient.Select("qoo10jp_sync_jobs", query, &staleJobs)
	if err != nil {
		return 0, fmt.Errorf("failed to find stale jobs: %w", err)
	}

	cleanedCount := 0
	for _, job := range staleJobs {
		err := s.MarkJobFailed(job.ID, fmt.Sprintf("Job timed out after %d minutes", timeoutMinutes))
		if err != nil {
			log.Printf("⚠️ Failed to cleanup stale job %s: %v", job.ID, err)
			continue
		}
		cleanedCount++
	}

	if cleanedCount > 0 {
		log.Printf("🧹 Cleaned up %d stale jobs", cleanedCount)
	}

	return cleanedCount, nil
}

// ========================================
// Statistics Types
// ========================================

// SyncJobStatistics represents job statistics
type SyncJobStatistics struct {
	SellerID       string    `json:"seller_id"`
	Period         int       `json:"period_days"`
	StartDate      time.Time `json:"start_date"`
	EndDate        time.Time `json:"end_date"`
	TotalJobs      int       `json:"total_jobs"`
	SuccessfulJobs int       `json:"successful_jobs"`
	FailedJobs     int       `json:"failed_jobs"`
	CancelledJobs  int       `json:"cancelled_jobs"`
	ActiveJobs     int       `json:"active_jobs"`
	TotalCollected int       `json:"total_collected"`
	TotalSaved     int       `json:"total_saved"`
	AvgDurationMs  int64     `json:"avg_duration_ms"`
}
