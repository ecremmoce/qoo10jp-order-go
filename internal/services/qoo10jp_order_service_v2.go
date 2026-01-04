package services

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"qoo10jp-order-go/internal/config"
	"qoo10jp-order-go/internal/models"
	"qoo10jp-order-go/pkg/crypto"
	"qoo10jp-order-go/pkg/qoo10jp"
	"qoo10jp-order-go/pkg/redis"
	"qoo10jp-order-go/pkg/supabase"

	"github.com/google/uuid"
)

// Qoo10JPOrderServiceV2 handles Qoo10JP order operations with v2 tables
type Qoo10JPOrderServiceV2 struct {
	cfg            *config.Config
	supabaseClient *supabase.Client
	redisClient    *redis.Client
	syncJobService *SyncJobService
	aesCrypto      *crypto.AESCrypto
}

// NewQoo10JPOrderServiceV2 creates a new Qoo10JPOrderServiceV2 instance
func NewQoo10JPOrderServiceV2(cfg *config.Config, supabaseClient *supabase.Client, redisClient *redis.Client) *Qoo10JPOrderServiceV2 {
	svc := &Qoo10JPOrderServiceV2{
		cfg:            cfg,
		supabaseClient: supabaseClient,
		redisClient:    redisClient,
		syncJobService: NewSyncJobService(supabaseClient),
	}

	// AES 암호화 초기화 (선택사항)
	if cfg.Encryption.Key != "" {
		aesCrypto, err := crypto.NewAESCrypto(cfg.Encryption.Key)
		if err != nil {
			log.Printf("⚠️ AES 초기화 실패 (평문 키 사용): %v", err)
		} else {
			svc.aesCrypto = aesCrypto
			log.Printf("🔐 AES 암호화 모듈 초기화 완료")
		}
	}

	return svc
}

// ========================================
// Shop Management
// ========================================

// GetActiveShops retrieves all active Qoo10JP shops
func (s *Qoo10JPOrderServiceV2) GetActiveShops() ([]models.Qoo10JPShopV2, error) {
	var shops []models.Qoo10JPShopV2
	query := "is_active=eq.true&order=created_at.asc"

	err := s.supabaseClient.Select("qoo10jp_shops_v2", query, &shops)
	if err != nil {
		return nil, fmt.Errorf("failed to get active shops: %w", err)
	}

	log.Printf("📦 활성 Qoo10JP 상점 %d개 조회", len(shops))
	return shops, nil
}

// GetShopBySellerID retrieves a shop by seller ID
func (s *Qoo10JPOrderServiceV2) GetShopBySellerID(sellerID string) (*models.Qoo10JPShopV2, error) {
	var shops []models.Qoo10JPShopV2
	query := fmt.Sprintf("seller_id=eq.%s", sellerID)

	err := s.supabaseClient.Select("qoo10jp_shops_v2", query, &shops)
	if err != nil {
		return nil, fmt.Errorf("failed to get shop: %w", err)
	}

	if len(shops) == 0 {
		return nil, fmt.Errorf("shop not found: %s", sellerID)
	}

	return &shops[0], nil
}

// DecryptCertificationKey decrypts the certification key if encrypted
func (s *Qoo10JPOrderServiceV2) DecryptCertificationKey(encryptedKey string) (string, error) {
	if s.aesCrypto == nil {
		// 암호화되지 않은 키로 가정
		return encryptedKey, nil
	}

	decrypted, err := s.aesCrypto.DecryptToQoo10Token(encryptedKey)
	if err != nil {
		// 복호화 실패 시 원본 반환 (평문일 수 있음)
		log.Printf("⚠️ 인증키 복호화 실패, 평문으로 사용: %v", err)
		return encryptedKey, nil
	}

	return decrypted, nil
}

// ========================================
// Order Collection (V2)
// ========================================

// CollectOrdersForShop collects orders for a specific shop (v2 method)
func (s *Qoo10JPOrderServiceV2) CollectOrdersForShop(sellerID string, startDate, endDate time.Time) (*models.Qoo10JPCollectionResult, error) {
	startTime := time.Now()

	result := &models.Qoo10JPCollectionResult{
		SellerID:  sellerID,
		StartDate: startDate,
		EndDate:   endDate,
	}

	log.Printf("🚀 [V2] Qoo10JP 주문 수집 시작: seller=%s, 기간=%s~%s",
		sellerID, startDate.Format("2006-01-02"), endDate.Format("2006-01-02"))

	// 중복 수집 방지 확인
	canStart, activeJob, err := s.syncJobService.CanStartCollection(sellerID, "order_collection")
	if err != nil {
		return result, fmt.Errorf("failed to check collection status: %w", err)
	}

	if !canStart {
		log.Printf("⚠️ 이미 수집 중: seller=%s, job=%s", sellerID, activeJob.ID)
		return result, fmt.Errorf("collection already in progress: %s", activeJob.ID)
	}

	// SyncJob 생성
	job, err := s.syncJobService.CreateJob(sellerID, startDate, endDate, "order_collection")
	if err != nil {
		return result, fmt.Errorf("failed to create sync job: %w", err)
	}

	// 작업 시작
	if err := s.syncJobService.StartJob(job.ID); err != nil {
		return result, fmt.Errorf("failed to start sync job: %w", err)
	}

	// 상점 정보 조회
	shop, err := s.GetShopBySellerID(sellerID)
	if err != nil {
		s.syncJobService.MarkJobFailed(job.ID, fmt.Sprintf("Shop not found: %s", sellerID))
		return result, fmt.Errorf("shop not found: %s", sellerID)
	}

	// 인증키 복호화
	certKey, err := s.DecryptCertificationKey(shop.CertificationKey)
	if err != nil {
		s.syncJobService.MarkJobFailed(job.ID, fmt.Sprintf("Failed to decrypt key: %v", err))
		return result, fmt.Errorf("failed to decrypt certification key: %w", err)
	}

	// Qoo10JP API 클라이언트 생성
	qoo10Client := qoo10jp.NewClient(shop.APIID, certKey, s.cfg.Qoo10JP.BaseURL)

	// 주문 수집 실행
	collectErr := s.collectOrdersInternal(qoo10Client, shop, startDate, endDate, job.ID, result)

	// 결과 처리
	result.DurationMs = time.Since(startTime).Milliseconds()

	if collectErr != nil {
		s.syncJobService.MarkJobFailed(job.ID, collectErr.Error())
		result.Errors = append(result.Errors, collectErr.Error())
		return result, collectErr
	}

	// 작업 완료
	if err := s.syncJobService.CompleteJob(job.ID, result); err != nil {
		log.Printf("⚠️ SyncJob 완료 처리 실패: %v", err)
	}

	// 상점 동기화 시간 업데이트
	s.updateShopSyncTime(sellerID)

	log.Printf("✅ [V2] 주문 수집 완료: seller=%s, collected=%d, saved=%d, updated=%d, skipped=%d, duration=%dms",
		sellerID, result.TotalCollected, result.TotalSaved, result.TotalUpdated, result.TotalSkipped, result.DurationMs)

	return result, nil
}

// collectOrdersInternal performs the actual order collection
func (s *Qoo10JPOrderServiceV2) collectOrdersInternal(
	client *qoo10jp.Client,
	shop *models.Qoo10JPShopV2,
	startDate, endDate time.Time,
	jobID string,
	result *models.Qoo10JPCollectionResult,
) error {
	page := 1
	pageSize := 100
	hasMore := true

	for hasMore {
		log.Printf("📄 주문 조회 페이지 %d (seller=%s)", page, shop.SellerID)

		// API 호출 (전체 상태 조회)
		orderResp, err := client.GetOrders(startDate, endDate, page, pageSize)

		if err != nil {
			log.Printf("❌ API 호출 실패: %v", err)
			return fmt.Errorf("API call failed: %w", err)
		}

		if orderResp.ResultCode != 0 {
			log.Printf("❌ API 오류: code=%d, msg=%s", orderResp.ResultCode, orderResp.ResultMsg)
			return fmt.Errorf("API error: %s (code: %d)", orderResp.ResultMsg, orderResp.ResultCode)
		}

		orders := orderResp.ResultObject
		if len(orders) == 0 {
			log.Printf("📭 더 이상 주문 없음 (page=%d)", page)
			hasMore = false
			continue
		}

		result.TotalCollected += len(orders)
		log.Printf("📦 %d개 주문 수신 (page=%d, total=%d)", len(orders), page, result.TotalCollected)

		// 주문 처리
		for _, order := range orders {
			if err := s.processOrderV2(order, shop, nil, result); err != nil {
				log.Printf("⚠️ 주문 처리 실패: order_no=%d, err=%v", order.OrderNo, err)
				result.TotalFailed++
				result.Errors = append(result.Errors, fmt.Sprintf("OrderNo %d: %v", order.OrderNo, err))
			}
		}

		// 진행률 업데이트
		s.syncJobService.UpdateJobProgress(jobID,
			result.TotalCollected,
			result.TotalSaved,
			result.TotalUpdated,
			result.TotalSkipped,
			result.TotalFailed,
			float64(page*pageSize)/float64(result.TotalCollected+pageSize)*100,
		)

		// 다음 페이지 확인
		if len(orders) < pageSize {
			hasMore = false
		} else {
			page++
			time.Sleep(100 * time.Millisecond) // Rate limiting
		}
	}

	return nil
}

// processOrderV2 processes a single order for v2 tables
func (s *Qoo10JPOrderServiceV2) processOrderV2(
	apiOrder qoo10jp.Order,
	shop *models.Qoo10JPShopV2,
	rawJSON map[string]interface{},
	result *models.Qoo10JPCollectionResult,
) error {
	orderNo := strconv.FormatInt(apiOrder.OrderNo, 10)
	packNo := strconv.FormatInt(apiOrder.PackNo, 10)

	// 기존 주문 확인
	existingOrder, err := s.getExistingOrderV2(orderNo, packNo)
	if err != nil && !strings.Contains(err.Error(), "not found") {
		return fmt.Errorf("failed to check existing order: %w", err)
	}

	// 동기화 건너뛰기 여부 확인
	if existingOrder != nil && existingOrder.ShouldSkipSync() {
		log.Printf("⏭️ 동기화 건너뜀: order_no=%s, status=%s", orderNo, existingOrder.OmsStatus)
		result.TotalSkipped++
		return nil
	}

	// 주문 변환
	orderV2, err := s.convertToOrderV2(apiOrder, shop, rawJSON)
	if err != nil {
		return fmt.Errorf("failed to convert order: %w", err)
	}

	// 기존 주문 보존 필드 처리
	if existingOrder != nil {
		// 송장번호와 OMS 상태는 기존 값 유지 (외부에서 업데이트됨)
		if existingOrder.TrackingNumber != "" && orderV2.TrackingNumber == "" {
			orderV2.TrackingNumber = existingOrder.TrackingNumber
		}
		if existingOrder.OmsStatus != "" && existingOrder.OmsStatus != models.OmsStatusNew {
			orderV2.OmsStatus = existingOrder.OmsStatus
		}
		orderV2.ID = existingOrder.ID
		orderV2.CreatedAt = existingOrder.CreatedAt
	}

	// 저장 또는 업데이트
	if existingOrder == nil {
		// 새 주문 저장
		if err := s.saveOrderV2(orderV2); err != nil {
			return fmt.Errorf("failed to save order: %w", err)
		}
		result.TotalSaved++
		log.Printf("💾 새 주문 저장: order_no=%s, pack_no=%s", orderNo, packNo)
	} else {
		// 기존 주문 업데이트
		if s.shouldUpdateOrderV2(existingOrder, orderV2) {
			if err := s.updateOrderV2(orderV2); err != nil {
				return fmt.Errorf("failed to update order: %w", err)
			}
			result.TotalUpdated++
			log.Printf("🔄 주문 업데이트: order_no=%s", orderNo)
		} else {
			result.TotalSkipped++
		}
	}

	// 주문 아이템 저장
	if err := s.saveOrderItemV2(apiOrder, orderV2.OrderNo, orderV2.PackNo, shop.SellerID); err != nil {
		log.Printf("⚠️ 주문 아이템 저장 실패: %v", err)
	}

	return nil
}

// getExistingOrderV2 checks if an order already exists
func (s *Qoo10JPOrderServiceV2) getExistingOrderV2(orderNo, packNo string) (*models.Qoo10JPOrderV2, error) {
	var orders []models.Qoo10JPOrderV2
	query := fmt.Sprintf("order_no=eq.%s&pack_no=eq.%s", orderNo, packNo)

	err := s.supabaseClient.Select("orders_qoo10jp_v2", query, &orders)
	if err != nil {
		return nil, err
	}

	if len(orders) == 0 {
		return nil, fmt.Errorf("order not found: %s/%s", orderNo, packNo)
	}

	return &orders[0], nil
}

// convertToOrderV2 converts API order to v2 model
func (s *Qoo10JPOrderServiceV2) convertToOrderV2(apiOrder qoo10jp.Order, shop *models.Qoo10JPShopV2, rawJSON map[string]interface{}) (*models.Qoo10JPOrderV2, error) {
	now := time.Now()

	orderV2 := &models.Qoo10JPOrderV2{
		ID:       uuid.New().String(),
		OrderNo:  strconv.FormatInt(apiOrder.OrderNo, 10),
		PackNo:   strconv.FormatInt(apiOrder.PackNo, 10),
		SellerID: shop.SellerID,
		Region:   shop.Region,

		// 상태
		Qoo10Status: apiOrder.ShippingStatus,
		OmsStatus:   models.GetOmsStatusFromQoo10Status(apiOrder.ShippingStatus),

		// 금액
		OrderPrice:     apiOrder.OrderPrice,
		TotalAmount:    apiOrder.Total,
		Discount:       apiOrder.Discount,
		SellerDiscount: apiOrder.SellerDiscount,
		ShippingRate:   apiOrder.ShippingRate,
		SettlePrice:    apiOrder.SettlePrice,
		Currency:       apiOrder.Currency,

		// 구매자
		BuyerName:   apiOrder.Buyer,
		BuyerKana:   apiOrder.BuyerKana,
		BuyerPhone:  apiOrder.BuyerTel,
		BuyerMobile: apiOrder.BuyerMobile,
		BuyerEmail:  apiOrder.BuyerEmail,

		// 수령인
		RecipientName:     apiOrder.Receiver,
		RecipientKana:     apiOrder.ReceiverKana,
		RecipientPhone:    apiOrder.ReceiverTel,
		RecipientMobile:   apiOrder.ReceiverMobile,
		RecipientZipcode:  apiOrder.ZipCode,
		RecipientAddress:  apiOrder.ShippingAddress,
		RecipientAddress1: apiOrder.Address1,
		RecipientAddress2: apiOrder.Address2,

		// 배송
		DeliveryCompany:  apiOrder.DeliveryCompany,
		TrackingNumber:   apiOrder.TrackingNo,
		PackingNo:        apiOrder.PackingNo,
		SellerDeliveryNo: apiOrder.SellerDeliveryNo,

		// 결제
		PaymentMethod: apiOrder.PaymentMethod,

		// 상품 정보
		ItemNo:         apiOrder.ItemNo,
		ItemTitle:      apiOrder.ItemTitle,
		SellerItemCode: apiOrder.SellerItemCode,
		OptionName:     apiOrder.Option,
		OptionCode:     apiOrder.OptionCode,
		OrderQty:       apiOrder.OrderQty,

		// 시스템
		SyncedAt:  &now,
		CreatedAt: now,
		UpdatedAt: now,
	}

	// 날짜 파싱
	orderV2.OrderDate = s.parseQoo10Date(apiOrder.OrderDate)
	orderV2.PaymentDate = s.parseQoo10Date(apiOrder.PaymentDate)
	orderV2.ShippingDate = s.parseQoo10Date(apiOrder.ShippingDate)
	orderV2.DeliveredDate = s.parseQoo10Date(apiOrder.DeliveredDate)

	// Raw data 저장
	if rawJSON != nil {
		if rawBytes, err := json.Marshal(rawJSON); err == nil {
			orderV2.RawData = rawBytes
		}
	}

	return orderV2, nil
}

// parseQoo10Date parses Qoo10 date string
func (s *Qoo10JPOrderServiceV2) parseQoo10Date(dateStr string) *time.Time {
	if dateStr == "" {
		return nil
	}

	// 다양한 날짜 형식 시도
	formats := []string{
		"2006-01-02 15:04:05",
		"2006/01/02 15:04:05",
		"20060102150405",
		"2006-01-02",
		"2006/01/02",
		"20060102",
	}

	for _, format := range formats {
		if t, err := time.Parse(format, dateStr); err == nil {
			return &t
		}
	}

	log.Printf("⚠️ 날짜 파싱 실패: %s", dateStr)
	return nil
}

// shouldUpdateOrderV2 checks if order should be updated
func (s *Qoo10JPOrderServiceV2) shouldUpdateOrderV2(existing, new *models.Qoo10JPOrderV2) bool {
	// Qoo10 상태가 변경된 경우
	if existing.Qoo10Status != new.Qoo10Status {
		return true
	}

	// 배송사/송장번호가 새로 입력된 경우
	if existing.TrackingNumber == "" && new.TrackingNumber != "" {
		return true
	}
	if existing.DeliveryCompany == "" && new.DeliveryCompany != "" {
		return true
	}

	// 금액이 변경된 경우
	if existing.TotalAmount != new.TotalAmount {
		return true
	}

	return false
}

// saveOrderV2 saves a new order to the v2 table
func (s *Qoo10JPOrderServiceV2) saveOrderV2(order *models.Qoo10JPOrderV2) error {
	return s.supabaseClient.Insert("orders_qoo10jp_v2", order)
}

// updateOrderV2 updates an existing order in the v2 table
func (s *Qoo10JPOrderServiceV2) updateOrderV2(order *models.Qoo10JPOrderV2) error {
	order.UpdatedAt = time.Now()
	query := fmt.Sprintf("id=eq.%s", order.ID)
	return s.supabaseClient.Update("orders_qoo10jp_v2", query, order)
}

// saveOrderItemV2 saves order item to v2 table
func (s *Qoo10JPOrderServiceV2) saveOrderItemV2(apiOrder qoo10jp.Order, orderNo, packNo, sellerID string) error {
	item := &models.Qoo10JPOrderItemV2{
		ID:         uuid.New().String(),
		OrderNo:    orderNo,
		PackNo:     packNo,
		SellerID:   sellerID,
		ItemNo:     apiOrder.ItemNo,
		ItemCode:   apiOrder.SellerItemCode,
		ItemName:   apiOrder.ItemTitle,
		ItemTitle:  apiOrder.ItemTitle,
		OptionName: apiOrder.Option,
		OptionCode: apiOrder.OptionCode,
		Quantity:   apiOrder.OrderQty,
		UnitPrice:  apiOrder.OrderPrice,
		TotalPrice: apiOrder.Total,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	// Upsert (기존 아이템 있으면 업데이트)
	return s.supabaseClient.BulkUpsert("order_items_qoo10jp_v2", []interface{}{item}, "order_no", "pack_no", "item_no")
}

// updateShopSyncTime updates the last sync time for a shop
func (s *Qoo10JPOrderServiceV2) updateShopSyncTime(sellerID string) {
	now := time.Now()
	update := map[string]interface{}{
		"last_sync_at": now,
		"updated_at":   now,
	}
	s.supabaseClient.Update("qoo10jp_shops_v2", fmt.Sprintf("seller_id=eq.%s", sellerID), update)
}

// ========================================
// Order Queries
// ========================================

// GetOrdersV2 retrieves orders with filter
func (s *Qoo10JPOrderServiceV2) GetOrdersV2(filter models.Qoo10JPOrderFilter) ([]models.Qoo10JPOrderV2, error) {
	var orders []models.Qoo10JPOrderV2

	queryParts := []string{}

	if filter.SellerID != "" {
		queryParts = append(queryParts, fmt.Sprintf("seller_id=eq.%s", filter.SellerID))
	}
	if filter.Status != "" {
		queryParts = append(queryParts, fmt.Sprintf("oms_status=eq.%s", filter.Status))
	}
	if filter.StartDate != nil {
		queryParts = append(queryParts, fmt.Sprintf("order_date=gte.%s", filter.StartDate.Format("2006-01-02")))
	}
	if filter.EndDate != nil {
		queryParts = append(queryParts, fmt.Sprintf("order_date=lte.%s", filter.EndDate.Format("2006-01-02")))
	}

	queryParts = append(queryParts, "order=order_date.desc")

	if filter.Limit > 0 {
		queryParts = append(queryParts, fmt.Sprintf("limit=%d", filter.Limit))
	}
	if filter.Offset > 0 {
		queryParts = append(queryParts, fmt.Sprintf("offset=%d", filter.Offset))
	}

	query := strings.Join(queryParts, "&")

	err := s.supabaseClient.Select("orders_qoo10jp_v2", query, &orders)
	if err != nil {
		return nil, fmt.Errorf("failed to get orders: %w", err)
	}

	return orders, nil
}

// GetOrderStatsV2 retrieves order statistics
func (s *Qoo10JPOrderServiceV2) GetOrderStatsV2(sellerID string, days int) (*models.Qoo10JPOrderStats, error) {
	startDate := time.Now().AddDate(0, 0, -days)

	var orders []models.Qoo10JPOrderV2
	query := fmt.Sprintf("seller_id=eq.%s&order_date=gte.%s", sellerID, startDate.Format("2006-01-02"))

	err := s.supabaseClient.Select("orders_qoo10jp_v2", query, &orders)
	if err != nil {
		return nil, fmt.Errorf("failed to get orders for stats: %w", err)
	}

	stats := &models.Qoo10JPOrderStats{
		TotalOrders: len(orders),
	}

	for _, order := range orders {
		stats.TotalAmount += order.TotalAmount

		switch order.OmsStatus {
		case models.OmsStatusNew:
			stats.PendingOrders++
		case models.OmsStatusShipped, models.OmsStatusDelivered:
			stats.ShippedOrders++
		case models.OmsStatusCompleted:
			stats.CompletedOrders++
		case models.OmsStatusCancelled:
			stats.CancelledOrders++
		}
	}

	return stats, nil
}
