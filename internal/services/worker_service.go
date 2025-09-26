package services

import (
	"context"
	"fmt"
	"log"
	"qoo10jp-order-go/internal/models"
	"sync"
	"time"
)

type WorkerService struct {
	schedulerService *SchedulerService
	workerCount      int
	ctx              context.Context
	cancel           context.CancelFunc
	wg               sync.WaitGroup
	isRunning        bool
	mu               sync.RWMutex
}

func NewWorkerService(schedulerService *SchedulerService, workerCount int) *WorkerService {
	ctx, cancel := context.WithCancel(context.Background())
	
	return &WorkerService{
		schedulerService: schedulerService,
		workerCount:      workerCount,
		ctx:              ctx,
		cancel:           cancel,
		isRunning:        false,
	}
}

// Start starts the worker service with specified number of workers
func (w *WorkerService) Start() {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.isRunning {
		log.Println("Worker service is already running")
		return
	}

	w.isRunning = true
	log.Printf("Starting worker service with %d workers", w.workerCount)

	// Start worker goroutines
	for i := 0; i < w.workerCount; i++ {
		w.wg.Add(1)
		go w.worker(i + 1)
	}

	log.Println("Worker service started successfully")
}

// Stop stops the worker service gracefully
func (w *WorkerService) Stop() {
	w.mu.Lock()
	defer w.mu.Unlock()

	if !w.isRunning {
		log.Println("Worker service is not running")
		return
	}

	log.Println("Stopping worker service...")
	
	// Cancel context to signal workers to stop
	w.cancel()
	
	// Wait for all workers to finish
	log.Println("Waiting for workers to finish...")
	w.wg.Wait()
	
	// Reset state
	w.isRunning = false
	
	// Create new context for future starts
	w.ctx, w.cancel = context.WithCancel(context.Background())
	
	log.Println("Worker service stopped successfully")
}

// IsRunning returns whether the worker service is currently running
func (w *WorkerService) IsRunning() bool {
	w.mu.RLock()
	defer w.mu.RUnlock()
	return w.isRunning
}

// worker is the main worker loop that processes jobs from the queue
func (w *WorkerService) worker(workerID int) {
	defer w.wg.Done()
	
	log.Printf("Worker %d started", workerID)
	
	for {
		select {
		case <-w.ctx.Done():
			log.Printf("Worker %d stopping due to context cancellation", workerID)
			return
		default:
			// Check context before processing
			select {
			case <-w.ctx.Done():
				log.Printf("Worker %d stopping due to context cancellation (inner check)", workerID)
				return
			default:
			}
			
			// Try to process a job with enhanced error handling
			err := w.processJobWithRetry(workerID)
			if err != nil {
				log.Printf("Worker %d error processing job: %v", workerID, err)
				
				// Check context before sleeping
				select {
				case <-w.ctx.Done():
					log.Printf("Worker %d stopping during error handling", workerID)
					return
				case <-time.After(2 * time.Second):
					// Continue after timeout
				}
				continue
			}

			// Small delay to prevent busy waiting when queue is empty
			select {
			case <-w.ctx.Done():
				log.Printf("Worker %d stopping during idle wait", workerID)
				return
			case <-time.After(1 * time.Second):
				// Continue after timeout
			}
		}
	}
}

// processJobWithRetry processes a job with enhanced retry logic
func (w *WorkerService) processJobWithRetry(workerID int) error {
	// Set worker status
	workerKey := fmt.Sprintf("worker_%d_status", workerID)
	w.schedulerService.redisClient.Set(workerKey, "processing", 30*time.Second)
	defer w.schedulerService.redisClient.Set(workerKey, "idle", 30*time.Second)
	
	// Try to get a job with timeout
	jobData, err := w.schedulerService.redisClient.PopFromQueue(OrderJobQueue)
	if err != nil {
		return fmt.Errorf("failed to pop job from queue: %v", err)
	}

	if jobData == "" {
		return nil // No jobs in queue
	}

	// Parse job
	job, err := models.OrderJobFromJSON(jobData)
	if err != nil {
		log.Printf("Worker %d: Failed to parse job, moving to dead letter queue: %v", workerID, err)
		w.schedulerService.redisClient.PushToQueue("dead_letter_jobs", jobData)
		return nil
	}

	// Check if job has exceeded max retries
	if job.RetryCount >= job.MaxRetries {
		log.Printf("Worker %d: Job %s exceeded max retries (%d), moving to failed jobs", workerID, job.ID, job.MaxRetries)
		w.schedulerService.redisClient.PushToQueue("failed_jobs", jobData)
		return nil
	}

	// Process the job
	err = w.schedulerService.executeJob(job)
	if err != nil {
		// Increment retry count and requeue if retries available
		job.RetryCount++
		if job.RetryCount < job.MaxRetries {
			log.Printf("Worker %d: Job %s failed, requeuing (retry %d/%d): %v", 
				workerID, job.ID, job.RetryCount, job.MaxRetries, err)
			
			retryJobData, _ := job.ToJSON()
			w.schedulerService.redisClient.PushToQueue(OrderJobQueue, retryJobData)
		} else {
			log.Printf("Worker %d: Job %s failed permanently after %d retries", workerID, job.ID, job.RetryCount)
			w.schedulerService.redisClient.PushToQueue("failed_jobs", jobData)
		}
		return err
	}

	// Increment processed counter
	w.schedulerService.redisClient.Incr("order_jobs_processed")
	
	log.Printf("Worker %d: Successfully processed job %s", workerID, job.ID)
	return nil
}

// SetWorkerCount dynamically adjusts the number of workers
func (w *WorkerService) SetWorkerCount(newCount int) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	if newCount < 1 {
		return fmt.Errorf("worker count must be at least 1")
	}

	if newCount > 20 {
		return fmt.Errorf("worker count cannot exceed 20")
	}

	oldCount := w.workerCount
	w.workerCount = newCount

	log.Printf("Worker count changed from %d to %d", oldCount, newCount)

	// If service is running, restart with new count
	if w.isRunning {
		log.Println("Restarting worker service with new count...")
		
		// Stop current workers
		w.cancel()
		w.wg.Wait()
		
		// Create new context and restart
		w.ctx, w.cancel = context.WithCancel(context.Background())
		
		// Start new workers
		for i := 0; i < w.workerCount; i++ {
			w.wg.Add(1)
			go w.worker(i + 1)
		}
		
		log.Printf("Worker service restarted with %d workers", w.workerCount)
	}

	return nil
}

// GetStatus returns the current status of the worker service
func (w *WorkerService) GetStatus() map[string]interface{} {
	w.mu.RLock()
	defer w.mu.RUnlock()

	queueLength, _ := w.schedulerService.GetQueueStatus()
	schedulerState, _ := w.schedulerService.GetSchedulerState()

	// Show active worker count (0 if stopped, configured count if running)
	activeWorkerCount := 0
	if w.isRunning {
		activeWorkerCount = w.workerCount
	}

	return map[string]interface{}{
		"is_running":           w.isRunning,
		"worker_count":         activeWorkerCount,
		"configured_workers":   w.workerCount,
		"queue_length":         queueLength,
		"last_execution_time":  schedulerState.LastExecutionTime,
		"next_schedule_time":   schedulerState.NextScheduleTime,
		"scheduler_is_running": schedulerState.IsRunning,
	}
}
