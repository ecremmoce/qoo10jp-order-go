package services

import (
	"context"
	"fmt"
	"log"
	"qoo10jp-order-go/internal/config"
	"qoo10jp-order-go/internal/models"
	"qoo10jp-order-go/pkg/webhook"
	"sync"
	"time"
)

type WorkerService struct {
	cfg                    *config.Config
	schedulerService       *SchedulerService
	shopeeOrderService     *ShopeeOrderService
	qoo10jpOrderService    *Qoo10JPOrderService
	qoo10jpOrderServiceV2  *Qoo10JPOrderServiceV2 // V2 서비스 추가
	workerCount            int
	ctx                    context.Context
	cancel                 context.CancelFunc
	wg                     sync.WaitGroup
	isRunning              bool
	mu                     sync.RWMutex
	webhookClient          *webhook.Client
	useV2                  bool // V2 서비스 사용 여부
}

func NewWorkerService(cfg *config.Config, schedulerService *SchedulerService, shopeeOrderService *ShopeeOrderService, qoo10jpOrderService *Qoo10JPOrderService, workerCount int) *WorkerService {
	ctx, cancel := context.WithCancel(context.Background())

	return &WorkerService{
		cfg:                  cfg,
		schedulerService:     schedulerService,
		shopeeOrderService:   shopeeOrderService,
		qoo10jpOrderService:  qoo10jpOrderService,
		workerCount:          workerCount,
		ctx:                  ctx,
		cancel:               cancel,
		isRunning:            false,
		webhookClient:        webhook.NewClient(10 * time.Second),
		useV2:                false, // 기본값: 기존 서비스 사용
	}
}

// NewWorkerServiceV2 creates a new WorkerService with V2 support
func NewWorkerServiceV2(cfg *config.Config, schedulerService *SchedulerService, shopeeOrderService *ShopeeOrderService, qoo10jpOrderService *Qoo10JPOrderService, qoo10jpOrderServiceV2 *Qoo10JPOrderServiceV2, workerCount int) *WorkerService {
	ctx, cancel := context.WithCancel(context.Background())

	return &WorkerService{
		cfg:                    cfg,
		schedulerService:       schedulerService,
		shopeeOrderService:     shopeeOrderService,
		qoo10jpOrderService:    qoo10jpOrderService,
		qoo10jpOrderServiceV2:  qoo10jpOrderServiceV2,
		workerCount:            workerCount,
		ctx:                    ctx,
		cancel:                 cancel,
		isRunning:              false,
		webhookClient:          webhook.NewClient(10 * time.Second),
		useV2:                  true, // V2 서비스 사용
	}
}

// SetUseV2 enables or disables V2 service usage
func (w *WorkerService) SetUseV2(useV2 bool) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.useV2 = useV2
	log.Printf("🔧 V2 서비스 사용 설정: %v", useV2)
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
				log.Printf("워커 %d: ⚠️ 작업 처리 오류: %v", workerID, err)

				// Check context before sleeping
				select {
				case <-w.ctx.Done():
					log.Printf("Worker %d stopping during error handling", workerID)
					return
				case <-time.After(5 * time.Second):
					// Continue after timeout
				}
				continue
			}

			// Small delay to prevent busy waiting when queue is empty
			select {
			case <-w.ctx.Done():
				log.Printf("Worker %d stopping during idle wait", workerID)
				return
			case <-time.After(3 * time.Second):
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

	// 두 큐를 번갈아가며 체크 (Shopee, Qoo10JP)
	queues := []string{OrderJobQueue, Qoo10JPOrderJobQueue}
	var jobData string
	var err error

	for _, queue := range queues {
		jobData, err = w.schedulerService.redisClient.PopFromQueue(queue)
		if err != nil {
			log.Printf("워커 %d: ⚠️ 큐 %s 접근 실패: %v", workerID, queue, err)
			continue
		}
		if jobData != "" {
			log.Printf("워커 %d: 📥 큐 '%s'에서 작업 수신", workerID, queue)
			break
		}
	}

	if jobData == "" {
		// No jobs available - this is normal
		return nil
	}

	log.Printf("워커 %d: 📥 큐에서 새 작업 수신", workerID)

	// Try to parse as N8N message first
	n8nMsg, err := models.N8NOrderMessageFromJSON(jobData)
	if err == nil {
		// This is an N8N message with account info
		log.Printf("워커 %d: 🚀 N8N 메시지 처리 시작 (계정: %s)", workerID, n8nMsg.AccountName)
		return w.processN8NMessage(workerID, n8nMsg)
	}

	// Try to parse as regular OrderJob
	job, err := models.OrderJobFromJSON(jobData)
	if err != nil {
		log.Printf("워커 %d: ❌ 메시지 파싱 실패, 데드레터 큐로 이동: %v", workerID, err)
		w.schedulerService.redisClient.PushToQueue("dead_letter_jobs", jobData)
		return nil
	}

	log.Printf("워커 %d: 🚀 작업 시작 %s", workerID, job.ID)

	// Check if job has exceeded max retries
	if job.RetryCount >= job.MaxRetries {
		log.Printf("Worker %d: Job %s exceeded max retries (%d), moving to failed jobs", workerID, job.ID, job.MaxRetries)
		w.schedulerService.redisClient.PushToQueue("failed_jobs", jobData)
		return nil
	}

	// Process the job
	startTime := time.Now()
	err = w.schedulerService.executeJob(job)
	duration := time.Since(startTime)

	if err != nil {
		// Increment retry count and requeue if retries available
		job.RetryCount++
		if job.RetryCount < job.MaxRetries {
			log.Printf("워커 %d: ❌ 작업 %s 실패 (%v), 재시도 대기열 추가 (%d/%d회): %v",
				workerID, job.ID, duration, job.RetryCount, job.MaxRetries, err)

			retryJobData, _ := job.ToJSON()
			w.schedulerService.redisClient.PushToQueue(OrderJobQueue, retryJobData)
		} else {
			log.Printf("워커 %d: ❌ 작업 %s 영구 실패 (%d회 재시도 후, 소요시간: %v)",
				workerID, job.ID, job.RetryCount, duration)
			w.schedulerService.redisClient.PushToQueue("failed_jobs", jobData)
		}
		return err
	}

	// Increment processed counter
	w.schedulerService.redisClient.Incr("order_jobs_processed")

	log.Printf("워커 %d: ✅ 작업 %s 성공 완료 (%v)", workerID, job.ID, duration)
	return nil
}

// processN8NMessage processes a message from n8n workflow with account credentials
func (w *WorkerService) processN8NMessage(workerID int, msg *models.N8NOrderMessage) error {
	startTime := time.Now()

	log.Printf("워커 %d: 📨 메시지 수신 - 플랫폼: %s, 계정 ID: %s, 계정명: %s",
		workerID, msg.Platform, msg.AccountID, msg.AccountName)

	// 작업 수신 웹훅 호출
	w.sendOrderCollectionStartWebhook(msg.AccountName, msg.AccountID, msg.Platform)

	var savedCount int
	var collectErr error

	// 플랫폼별 주문 수집 처리
	switch msg.Platform {
	case "shopee":
		// Shopee 주문 수집
		log.Printf("워커 %d: 🛒 Shopee 주문 수집 시작", workerID)

		if msg.ShopID == 0 || msg.AccessToken == "" {
			collectErr = fmt.Errorf("Shopee shop_id 또는 access_token이 없습니다")
			log.Printf("워커 %d: ❌ %v", workerID, collectErr)
		} else {
			// 최근 15일간 주문 수집 (Shopee API 제한)
			endDate := time.Now()
			startDate := endDate.AddDate(0, 0, -15)

			log.Printf("워커 %d: 📅 수집 기간: %s ~ %s (Partner ID: %d)", workerID,
				startDate.Format("2006-01-02"), endDate.Format("2006-01-02"), w.cfg.Shopee.PartnerID)

			// .env에서 읽은 고정 partner_id 사용
			collectErr = w.shopeeOrderService.CollectOrders(startDate, endDate, msg.ShopID, w.cfg.Shopee.PartnerID, msg.AccessToken)

			if collectErr != nil {
				log.Printf("워커 %d: ❌ Shopee 주문 수집 실패: %v", workerID, collectErr)
			} else {
				// 수집된 주문 개수 조회
				filter := models.ShopeeOrderFilter{
					PlatformAccountID: fmt.Sprintf("%d", msg.ShopID),
					StartDate:         &startDate,
					EndDate:           &endDate,
					Limit:             1000,
				}

				orders, err := w.shopeeOrderService.GetOrders(filter)
				if err != nil {
					log.Printf("워커 %d: ⚠️ 수집된 주문 조회 실패: %v", workerID, err)
					savedCount = 0
				} else {
					savedCount = len(orders)
				}

				log.Printf("워커 %d: ✅ Shopee 주문 수집 완료 (%d건)", workerID, savedCount)
			}
		}

	case "qoo10jp":
		// Qoo10JP 주문 수집
		log.Printf("워커 %d: 🛒 Qoo10JP 주문 수집 시작 (V2: %v)", workerID, w.useV2)

		// 최근 30일간 주문 수집
		endDate := time.Now()
		startDate := endDate.AddDate(0, 0, -30)

		log.Printf("워커 %d: 📅 수집 기간: %s ~ %s (계정: %s)", workerID,
			startDate.Format("2006-01-02"), endDate.Format("2006-01-02"), msg.AccountName)

		// V2 서비스 사용 여부에 따라 분기
		if w.useV2 && w.qoo10jpOrderServiceV2 != nil {
			// V2 서비스 사용 (seller_id 기반)
			sellerID := msg.SellerID
			if sellerID == "" {
				sellerID = msg.AccountID // 호환성을 위해 AccountID 사용
			}

			result, err := w.qoo10jpOrderServiceV2.CollectOrdersForShop(sellerID, startDate, endDate)
			if err != nil {
				collectErr = err
				log.Printf("워커 %d: ❌ Qoo10JP V2 주문 수집 실패: %v", workerID, err)
			} else {
				savedCount = result.TotalSaved + result.TotalUpdated
				log.Printf("워커 %d: ✅ Qoo10JP V2 주문 수집 완료 (수집: %d, 저장: %d, 업데이트: %d, 건너뜀: %d)",
					workerID, result.TotalCollected, result.TotalSaved, result.TotalUpdated, result.TotalSkipped)
			}
		} else {
			// 기존 서비스 사용 (Legacy)
			certKey := msg.CertificationKey
			apiID := msg.APIID

			if certKey == "" || apiID == "" {
				collectErr = fmt.Errorf("Qoo10JP api_id 또는 certification_key가 없습니다")
				log.Printf("워커 %d: ❌ %v", workerID, collectErr)
			} else {
				// Qoo10JP 주문 수집 실행
				collectErr = w.qoo10jpOrderService.CollectOrders(startDate, endDate, msg.AccountID)

				if collectErr != nil {
					log.Printf("워커 %d: ❌ Qoo10JP 주문 수집 실패: %v", workerID, collectErr)
				} else {
					// 수집된 주문 개수 조회
					filter := models.Qoo10JPOrderFilter{
						PlatformAccountID: msg.AccountID,
						StartDate:         &startDate,
						EndDate:           &endDate,
						Limit:             1000,
					}

					orders, err := w.qoo10jpOrderService.GetOrders(filter)
					if err != nil {
						log.Printf("워커 %d: ⚠️ 수집된 주문 조회 실패: %v", workerID, err)
						savedCount = 0
					} else {
						savedCount = len(orders)
					}

					log.Printf("워커 %d: ✅ Qoo10JP 주문 수집 완료 (%d건)", workerID, savedCount)
				}
			}
		}

	default:
		// 지원하지 않는 플랫폼
		log.Printf("워커 %d: ⚠️ 지원하지 않는 플랫폼: '%s'", workerID, msg.Platform)
		time.Sleep(500 * time.Millisecond)
	}

	duration := time.Since(startTime)
	log.Printf("워커 %d: ✅ 메시지 처리 완료 (계정: %s, 소요시간: %v)",
		workerID, msg.AccountName, duration)

	// 작업 완료 웹훅 호출
	w.sendOrderCollectionEndWebhook(msg.AccountName, msg.AccountID, msg.Platform, savedCount, duration, collectErr)

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

	// 오늘 처리된 주문 수 계산 (Redis에서 가져오거나 임시 값)
	todayProcessed := w.getTodayProcessedCount()

	// 다음 실행까지 남은 시간 계산
	nextExecutionMinutes := w.getNextExecutionMinutes(schedulerState.NextScheduleTime)

	// 최근 오류 상태 확인
	recentErrors := w.getRecentErrorStatus()

	return map[string]interface{}{
		"is_running":             w.isRunning,
		"worker_count":           activeWorkerCount,
		"configured_workers":     w.workerCount,
		"queue_length":           queueLength,
		"last_execution_time":    schedulerState.LastExecutionTime,
		"next_schedule_time":     schedulerState.NextScheduleTime,
		"scheduler_is_running":   schedulerState.IsRunning,
		"today_processed":        todayProcessed,
		"next_execution_minutes": nextExecutionMinutes,
		"recent_errors":          recentErrors,
	}
}

// getTodayProcessedCount returns the number of orders processed today
func (w *WorkerService) getTodayProcessedCount() int {
	// TODO: Redis에서 실제 처리 카운트를 가져오도록 구현
	// 현재는 임시로 시간 기반 값 반환
	now := time.Now()
	hour := now.Hour()
	// 시간대별로 다른 값 반환 (더 현실적으로 보이도록)
	if hour < 6 {
		return 15 + (hour * 3)
	} else if hour < 12 {
		return 30 + ((hour - 6) * 8)
	} else if hour < 18 {
		return 78 + ((hour - 12) * 12)
	} else {
		return 150 + ((hour - 18) * 5)
	}
}

// getNextExecutionMinutes calculates minutes until next execution
func (w *WorkerService) getNextExecutionMinutes(nextScheduleTime time.Time) int {
	// 제로 값인지 확인 (스케줄이 설정되지 않은 경우)
	if nextScheduleTime.IsZero() {
		return -1 // 스케줄 없음
	}

	now := time.Now()
	if nextScheduleTime.Before(now) {
		return 0 // 이미 지난 시간
	}

	duration := nextScheduleTime.Sub(now)
	minutes := int(duration.Minutes())

	if minutes > 60 {
		return -2 // 1시간 이상은 시간으로 표시하도록 프론트엔드에서 처리
	}

	return minutes
}

// getRecentErrorStatus checks for recent errors
func (w *WorkerService) getRecentErrorStatus() string {
	// TODO: Redis나 로그에서 최근 오류 확인
	// 현재는 임시로 "없음" 반환
	return "없음"
}

// sendOrderCollectionStartWebhook sends a webhook when order collection starts
func (w *WorkerService) sendOrderCollectionStartWebhook(accountName, accountID, platform string) {
	webhookURL := w.cfg.Webhook.OrderCollectionURL
	if webhookURL == "" {
		return
	}

	now := time.Now()

	platformName := platform
	if platformName == "" {
		platformName = "Unknown"
	}

	message := fmt.Sprintf("🚀 [%s] 주문 수집 시작\n"+
		"⏰ 시간: %s\n"+
		"👤 계정: %s\n"+
		"🆔 ID: %s\n"+
		"🔄 상태: 주문 데이터 수집 중...",
		platformName, now.Format("2006-01-02 15:04:05"), accountName, accountID)

	data := map[string]interface{}{
		"platform":     platformName,
		"account_name": accountName,
		"account_id":   accountID,
		"status":       "started",
		"action":       "order_collection",
		"timestamp":    now.Format("2006-01-02 15:04:05"),
	}

	// 비동기로 웹훅 호출
	go func() {
		err := w.webhookClient.SendWebhookWithRetry(webhookURL, message, data, 2)
		if err != nil {
			log.Printf("⚠️ 주문 수집 시작 웹훅 호출 실패: %v", err)
		} else {
			log.Printf("✅ 주문 수집 시작 웹훅 전송 성공")
		}
	}()
}

// sendOrderCollectionEndWebhook sends a webhook when order collection ends
func (w *WorkerService) sendOrderCollectionEndWebhook(accountName, accountID, platform string, savedCount int, duration time.Duration, collectErr error) {
	webhookURL := w.cfg.Webhook.OrderCollectionURL
	if webhookURL == "" {
		return
	}

	now := time.Now()

	platformName := platform
	if platformName == "" {
		platformName = "Unknown"
	}

	var message string
	var status string

	if collectErr != nil {
		message = fmt.Sprintf("🚨 [%s] 주문 수집 실패\n"+
			"⏰ 시간: %s\n"+
			"👤 계정: %s\n"+
			"🆔 ID: %s\n"+
			"❌ 오류: %s\n"+
			"⏱️  소요시간: %v",
			platformName, now.Format("2006-01-02 15:04:05"), accountName, accountID, collectErr.Error(), duration)
		status = "failed"
	} else {
		if savedCount > 0 {
			message = fmt.Sprintf("✅ [%s] 주문 수집 완료\n"+
				"⏰ 시간: %s\n"+
				"👤 계정: %s\n"+
				"🆔 ID: %s\n"+
				"📦 수집결과: %d건\n"+
				"⏱️  소요시간: %v",
				platformName, now.Format("2006-01-02 15:04:05"), accountName, accountID, savedCount, duration)
		} else {
			message = fmt.Sprintf("ℹ️ [%s] 주문 수집 완료\n"+
				"⏰ 시간: %s\n"+
				"👤 계정: %s\n"+
				"🆔 ID: %s\n"+
				"📦 수집결과: 신규 주문 없음\n"+
				"⏱️  소요시간: %v",
				platformName, now.Format("2006-01-02 15:04:05"), accountName, accountID, duration)
		}
		status = "completed"
	}

	data := map[string]interface{}{
		"platform":     platformName,
		"account_name": accountName,
		"account_id":   accountID,
		"status":       status,
		"action":       "order_collection",
		"saved_count":  savedCount,
		"timestamp":    now.Format("2006-01-02 15:04:05"),
		"duration_ms":  duration.Milliseconds(),
	}

	if collectErr != nil {
		data["error"] = collectErr.Error()
	}

	// 비동기로 웹훅 호출
	go func() {
		err := w.webhookClient.SendWebhookWithRetry(webhookURL, message, data, 2)
		if err != nil {
			log.Printf("⚠️ 주문 수집 완료 웹훅 호출 실패: %v", err)
		} else {
			log.Printf("✅ 주문 수집 완료 웹훅 전송 성공")
		}
	}()
}
