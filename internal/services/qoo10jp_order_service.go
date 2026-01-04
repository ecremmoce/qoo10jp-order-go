package services

import (
	"fmt"
	"log"
	"qoo10jp-order-go/internal/config"
	"qoo10jp-order-go/internal/models"
	"qoo10jp-order-go/pkg/qoo10jp"
	"qoo10jp-order-go/pkg/redis"
	"qoo10jp-order-go/pkg/supabase"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Qoo10JPOrderService struct {
	cfg            *config.Config
	supabaseClient *supabase.Client
	redisClient    *redis.Client
}

func NewQoo10JPOrderService(cfg *config.Config, supabaseClient *supabase.Client, redisClient *redis.Client) *Qoo10JPOrderService {
	return &Qoo10JPOrderService{
		cfg:            cfg,
		supabaseClient: supabaseClient,
		redisClient:    redisClient,
	}
}

// Supabase에서 계정 정보를 가져와서 Qoo10JP 클라이언트 생성
func (s *Qoo10JPOrderService) createQoo10JPClient(platformAccountID string) (*qoo10jp.Client, error) {
	// Supabase에서 계정 정보 조회
	var accounts []struct {
		ID               string `json:"id"`
		SellerID         string `json:"seller_id"`
		APIId            string `json:"api_id"`
		CertificationKey string `json:"certification_key"`
		IsActive         bool   `json:"is_active"`
	}

	query := fmt.Sprintf("id=eq.%s&is_active=eq.true", platformAccountID)
	err := s.supabaseClient.Select("sales_platform_accounts_qoo10jp", query, &accounts)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch account info: %v", err)
	}

	if len(accounts) == 0 {
		return nil, fmt.Errorf("no active account found with ID: %s", platformAccountID)
	}

	account := accounts[0]
	if account.APIId == "" || account.CertificationKey == "" {
		return nil, fmt.Errorf("account %s missing API credentials", platformAccountID)
	}

	// Qoo10JP 클라이언트 생성 (환경 설정에서 BaseURL 사용)
	return qoo10jp.NewClient(account.APIId, account.CertificationKey, s.cfg.Qoo10JP.BaseURL), nil
}

// 주문 수집
func (s *Qoo10JPOrderService) CollectOrders(startDate, endDate time.Time, platformAccountID string) error {
	// 계정 정보로 클라이언트 생성
	qoo10jpClient, err := s.createQoo10JPClient(platformAccountID)
	if err != nil {
		return fmt.Errorf("failed to create qoo10jp client: %v", err)
	}

	page := 1
	pageSize := 100
	maxPages := 1000            // 최대 페이지 수 제한 (100,000개 주문)
	consecutiveCachedPages := 0 // 연속으로 캐시된 페이지 수

	for page <= maxPages {
		// 캐시 키 생성
		cacheKey := fmt.Sprintf("qoo10jp:orders:collected:%s:%s:account:%s:page:%d",
			startDate.Format("2006-01-02"), endDate.Format("2006-01-02"), platformAccountID, page)

		// 이미 처리된 페이지인지 확인
		exists, err := s.redisClient.Exists(cacheKey)
		if err != nil {
			log.Printf("Redis error: %v", err)
		} else if exists {
			log.Printf("Page %d already processed recently, skipping", page)
			consecutiveCachedPages++

			// 연속으로 10개 페이지가 캐시되어 있으면 중단 (무한루프 방지)
			if consecutiveCachedPages >= 10 {
				log.Printf("Found %d consecutive cached pages, assuming no more new data", consecutiveCachedPages)
				break
			}

			page++
			continue
		}

		// 캐시되지 않은 페이지를 만나면 카운터 리셋
		consecutiveCachedPages = 0

		// Qoo10JP API에서 주문 데이터 가져오기
		orderResp, err := qoo10jpClient.GetOrders(startDate, endDate, page, pageSize)
		if err != nil {
			return fmt.Errorf("failed to fetch orders: %v", err)
		}

		if orderResp.ResultCode != 0 {
			return fmt.Errorf("API error: %s", orderResp.ResultMsg)
		}

		if len(orderResp.ResultObject) == 0 {
			log.Println("No more orders to process")
			break
		}

		// 배치 처리를 위한 주문 데이터 수집 및 분류
		err = s.processBatchOrders(orderResp.ResultObject, platformAccountID)
		if err != nil {
			log.Printf("Failed to process batch orders: %v", err)
			// 실패 시 개별 처리로 폴백
			s.processFallbackOrders(orderResp.ResultObject, platformAccountID)
		}

		// 처리 완료된 페이지를 캐시에 저장 (1시간)
		s.redisClient.Set(cacheKey, "processed", time.Hour)

		log.Printf("Processed page %d with %d orders", page, len(orderResp.ResultObject))

		// 페이지 크기보다 적으면 마지막 페이지
		if len(orderResp.ResultObject) < pageSize {
			log.Printf("Reached last page (page %d), stopping", page)
			break
		}

		page++

		// 안전장치: 너무 많은 페이지 처리 방지
		if page > maxPages {
			log.Printf("Reached maximum page limit (%d), stopping", maxPages)
			break
		}
	}

	return nil
}

// 주문 조회
func (s *Qoo10JPOrderService) GetOrders(filter models.Qoo10JPOrderFilter) ([]models.Qoo10JPOrder, error) {
	query := s.buildQuery(filter)

	var orders []models.Qoo10JPOrder
	err := s.supabaseClient.Select("orders_qoo10jp", query, &orders)
	if err != nil {
		return nil, err
	}

	// 각 주문의 상품 정보 로드
	for i := range orders {
		var items []models.Qoo10JPOrderItem
		itemQuery := fmt.Sprintf("order_id=eq.%s", orders[i].ID)
		err = s.supabaseClient.Select("order_items_qoo10jp", itemQuery, &items)
		if err != nil {
			log.Printf("Failed to load items for order %s: %v", orders[i].ID, err)
		} else {
			orders[i].Items = items
		}
	}

	return orders, nil
}

// 주문 통계
func (s *Qoo10JPOrderService) GetOrderStats(platformAccountID string, startDate, endDate *time.Time) (*models.Qoo10JPOrderStats, error) {
	var conditions []string

	if platformAccountID != "" {
		conditions = append(conditions, fmt.Sprintf("platform_account_id=eq.%s", platformAccountID))
	}

	if startDate != nil {
		conditions = append(conditions, fmt.Sprintf("order_date=gte.%s", startDate.Format("2006-01-02")))
	}

	if endDate != nil {
		conditions = append(conditions, fmt.Sprintf("order_date=lte.%s", endDate.Format("2006-01-02")))
	}

	query := ""
	if len(conditions) > 0 {
		query = fmt.Sprintf("select=count(*),total_amount.sum(),order_status&%s",
			fmt.Sprintf("%s", conditions[0]))
		for i := 1; i < len(conditions); i++ {
			query += fmt.Sprintf("&%s", conditions[i])
		}
	}

	// 통계 쿼리 실행 (실제 구현은 더 복잡할 수 있음)
	var orders []models.Qoo10JPOrder
	err := s.supabaseClient.Select("orders_qoo10jp", query, &orders)
	if err != nil {
		return nil, err
	}

	// 통계 계산
	stats := &models.Qoo10JPOrderStats{}
	stats.TotalOrders = len(orders)

	for _, order := range orders {
		stats.TotalAmount += order.TotalAmount

		switch order.OrderStatus {
		case "pending", "processing":
			stats.PendingOrders++
		case "completed", "delivered":
			stats.CompletedOrders++
		case "cancelled", "refunded":
			stats.CancelledOrders++
		}
	}

	return stats, nil
}

// Qoo10JP 주문을 내부 모델로 변환
func (s *Qoo10JPOrderService) convertQoo10JPOrder(qooOrder qoo10jp.Order, platformAccountID string) (*models.Qoo10JPOrder, error) {
	orderDate, err := time.Parse("2006-01-02 15:04:05", qooOrder.OrderDate)
	if err != nil {
		return nil, err
	}

	order := &models.Qoo10JPOrder{
		ID:                uuid.New().String(),
		OrderNo:           fmt.Sprintf("%d", qooOrder.OrderNo),
		OrderDate:         orderDate,
		OrderStatus:       qooOrder.ShippingStatus,
		PaymentStatus:     qooOrder.PaymentMethod,
		BuyerID:           qooOrder.SellerID,
		BuyerName:         qooOrder.Buyer,
		BuyerEmail:        qooOrder.BuyerEmail,
		BuyerPhone:        qooOrder.BuyerMobile,
		TotalAmount:       qooOrder.Total,
		PaymentAmount:     qooOrder.Total,
		Currency:          qooOrder.Currency,
		ShippingAddress:   qooOrder.ShippingAddress,
		PlatformAccountID: platformAccountID,
		RawData:           qooOrder,
		SyncedAt:          time.Now(),
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}

	// 주문 상품 변환 (API 응답에서 각 주문이 개별 상품)
	item := models.Qoo10JPOrderItem{
		ID:         uuid.New().String(),
		OrderID:    order.ID,
		ItemCode:   qooOrder.ItemNo,
		ItemName:   qooOrder.ItemTitle,
		Quantity:   qooOrder.OrderQty,
		UnitPrice:  qooOrder.OrderPrice,
		TotalPrice: qooOrder.Total,
		RawData:    qooOrder,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	order.Items = append(order.Items, item)

	return order, nil
}

// 쿼리 빌더
func (s *Qoo10JPOrderService) buildQuery(filter models.Qoo10JPOrderFilter) string {
	var conditions []string

	if filter.StartDate != nil {
		conditions = append(conditions, fmt.Sprintf("order_date=gte.%s", filter.StartDate.Format("2006-01-02")))
	}

	if filter.EndDate != nil {
		conditions = append(conditions, fmt.Sprintf("order_date=lte.%s", filter.EndDate.Format("2006-01-02")))
	}

	if filter.OrderStatus != "" {
		conditions = append(conditions, fmt.Sprintf("order_status=eq.%s", filter.OrderStatus))
	}

	if filter.PaymentStatus != "" {
		conditions = append(conditions, fmt.Sprintf("payment_status=eq.%s", filter.PaymentStatus))
	}

	if filter.PlatformAccountID != "" {
		conditions = append(conditions, fmt.Sprintf("platform_account_id=eq.%s", filter.PlatformAccountID))
	}

	if filter.BuyerID != "" {
		conditions = append(conditions, fmt.Sprintf("buyer_id=eq.%s", filter.BuyerID))
	}

	query := ""
	if len(conditions) > 0 {
		query = fmt.Sprintf("%s", conditions[0])
		for i := 1; i < len(conditions); i++ {
			query += fmt.Sprintf("&%s", conditions[i])
		}
	}

	if filter.Limit > 0 {
		if query != "" {
			query += "&"
		}
		query += fmt.Sprintf("limit=%d", filter.Limit)
	}

	if filter.Offset > 0 {
		if query != "" {
			query += "&"
		}
		query += fmt.Sprintf("offset=%d", filter.Offset)
	}

	// 최신 주문부터 정렬
	if query != "" {
		query += "&"
	}
	query += "order=order_date.desc"

	return query
}

// 배치 처리로 주문 저장 (성능 최적화)
func (s *Qoo10JPOrderService) processBatchOrders(qooOrders []qoo10jp.Order, platformAccountID string) error {
	if len(qooOrders) == 0 {
		return nil
	}

	// 1. 모든 주문 번호 수집
	orderNos := make([]string, 0, len(qooOrders))
	for _, qooOrder := range qooOrders {
		orderNos = append(orderNos, fmt.Sprintf("%d", qooOrder.OrderNo))
	}

	// 2. 기존 주문 일괄 조회
	var existingOrders []models.Qoo10JPOrder
	err := s.supabaseClient.SelectIn("orders_qoo10jp", "order_no", orderNos, "", &existingOrders)
	if err != nil {
		return fmt.Errorf("failed to fetch existing orders: %v", err)
	}

	// 3. 기존 주문을 맵으로 변환 (빠른 조회)
	existingOrderMap := make(map[string]models.Qoo10JPOrder)
	for _, order := range existingOrders {
		existingOrderMap[order.OrderNo] = order
	}

	// 4. 신규/업데이트 분류
	var newOrders []models.Qoo10JPOrder
	var updateOrders []models.Qoo10JPOrder
	var allNewOrderItems []models.Qoo10JPOrderItem
	var updateOrderIDs []string

	for _, qooOrder := range qooOrders {
		order, err := s.convertQoo10JPOrder(qooOrder, platformAccountID)
		if err != nil {
			log.Printf("Failed to convert order %d: %v", qooOrder.OrderNo, err)
			continue
		}

		if existingOrder, exists := existingOrderMap[order.OrderNo]; exists {
			// 업데이트 필요 여부 확인
			if s.shouldUpdateOrder(existingOrder, *order) {
				order.ID = existingOrder.ID               // 기존 ID 유지
				order.CreatedAt = existingOrder.CreatedAt // 생성일 유지
				updateOrders = append(updateOrders, *order)
				updateOrderIDs = append(updateOrderIDs, order.ID)

				// 주문 상품도 업데이트 대상에 추가
				for i := range order.Items {
					order.Items[i].OrderID = order.ID
					allNewOrderItems = append(allNewOrderItems, order.Items[i])
				}
			}
		} else {
			// 신규 주문
			newOrders = append(newOrders, *order)

			// 주문 상품도 신규 대상에 추가
			for i := range order.Items {
				order.Items[i].OrderID = order.ID
				allNewOrderItems = append(allNewOrderItems, order.Items[i])
			}
		}
	}

	// 5. 벌크 처리 실행 (에러 처리 개선)
	var bulkErrors []string

	// 신규 주문 벌크 삽입
	if len(newOrders) > 0 {
		log.Printf("Bulk inserting %d new orders", len(newOrders))
		err = s.bulkInsertWithRetry("orders_qoo10jp", newOrders)
		if err != nil {
			bulkErrors = append(bulkErrors, fmt.Sprintf("new orders: %v", err))
			// 실패 시 개별 처리
			s.insertOrdersIndividually(newOrders)
		}
	}

	// 기존 주문 벌크 업데이트
	if len(updateOrders) > 0 {
		log.Printf("Bulk updating %d existing orders", len(updateOrders))
		err = s.bulkUpsertWithRetry("orders_qoo10jp", updateOrders, "order_no")
		if err != nil {
			bulkErrors = append(bulkErrors, fmt.Sprintf("update orders: %v", err))
			// 실패 시 개별 처리
			s.updateOrdersIndividually(updateOrders)
		} else {
			// 업데이트 성공 시에만 기존 상품 삭제
			if len(updateOrderIDs) > 0 {
				err = s.supabaseClient.BulkDelete("order_items_qoo10jp", "order_id", updateOrderIDs)
				if err != nil {
					log.Printf("Failed to delete existing order items: %v", err)
				}
			}
		}
	}

	// 주문 상품 벌크 삽입
	if len(allNewOrderItems) > 0 {
		log.Printf("Bulk inserting %d order items", len(allNewOrderItems))
		err = s.bulkInsertWithRetry("order_items_qoo10jp", allNewOrderItems)
		if err != nil {
			bulkErrors = append(bulkErrors, fmt.Sprintf("order items: %v", err))
			// 실패 시 개별 처리
			s.insertOrderItemsIndividually(allNewOrderItems)
		}
	}

	// 부분 실패가 있었다면 경고 로그
	if len(bulkErrors) > 0 {
		log.Printf("Bulk processing had partial failures: %s", strings.Join(bulkErrors, "; "))
	}

	log.Printf("Successfully processed %d new orders and %d updated orders", len(newOrders), len(updateOrders))
	return nil
}

// 개별 처리 폴백 (배치 처리 실패 시)
func (s *Qoo10JPOrderService) processFallbackOrders(qooOrders []qoo10jp.Order, platformAccountID string) {
	log.Printf("Using fallback individual processing for %d orders", len(qooOrders))

	for _, qooOrder := range qooOrders {
		order, err := s.convertQoo10JPOrder(qooOrder, platformAccountID)
		if err != nil {
			log.Printf("Failed to convert order %d: %v", qooOrder.OrderNo, err)
			continue
		}

		// 기존 주문 확인
		var existingOrders []models.Qoo10JPOrder
		query := fmt.Sprintf("order_no=eq.%s", order.OrderNo)
		err = s.supabaseClient.Select("orders_qoo10jp", query, &existingOrders)
		if err != nil {
			log.Printf("Failed to check existing order: %v", err)
			continue
		}

		if len(existingOrders) > 0 {
			// 업데이트 필요 여부 확인
			if s.shouldUpdateOrder(existingOrders[0], *order) {
				log.Printf("Updating order %s", order.OrderNo)
				order.ID = existingOrders[0].ID
				order.CreatedAt = existingOrders[0].CreatedAt
				updateQuery := fmt.Sprintf("order_no=eq.%s", order.OrderNo)
				err = s.supabaseClient.Update("orders_qoo10jp", updateQuery, order)

				if err == nil {
					// 기존 주문 상품 삭제 후 재삽입
					deleteQuery := fmt.Sprintf("order_id=eq.%s", order.ID)
					s.supabaseClient.Delete("order_items_qoo10jp", deleteQuery)

					for i := range order.Items {
						order.Items[i].OrderID = order.ID
					}
					s.supabaseClient.Insert("order_items_qoo10jp", order.Items)
				}
			} else {
				log.Printf("Order %s unchanged, skipping", order.OrderNo)
			}
		} else {
			// 새 주문 삽입
			log.Printf("Inserting new order %s", order.OrderNo)
			err = s.supabaseClient.Insert("orders_qoo10jp", order)

			if err == nil && len(order.Items) > 0 {
				for i := range order.Items {
					order.Items[i].OrderID = order.ID
				}
				s.supabaseClient.Insert("order_items_qoo10jp", order.Items)
			}
		}

		if err != nil {
			log.Printf("Failed to save order %s: %v", order.OrderNo, err)
		}
	}
}

// 주문 업데이트 필요 여부 판단 (시간 기반 스킵 제거, 내용 비교)
func (s *Qoo10JPOrderService) shouldUpdateOrder(existing, new models.Qoo10JPOrder) bool {
	// 주요 필드들의 변경 여부 확인
	return existing.OrderStatus != new.OrderStatus ||
		existing.PaymentStatus != new.PaymentStatus ||
		existing.BuyerName != new.BuyerName ||
		existing.BuyerEmail != new.BuyerEmail ||
		existing.BuyerPhone != new.BuyerPhone ||
		existing.ShippingAddress != new.ShippingAddress ||
		existing.ShippingMethod != new.ShippingMethod ||
		existing.ShippingStatus != new.ShippingStatus ||
		existing.TrackingNumber != new.TrackingNumber ||
		existing.TotalAmount != new.TotalAmount ||
		existing.PaymentAmount != new.PaymentAmount ||
		existing.DiscountAmount != new.DiscountAmount ||
		existing.ShippingFee != new.ShippingFee
}

// 재시도 로직이 포함된 벌크 삽입
func (s *Qoo10JPOrderService) bulkInsertWithRetry(table string, data interface{}) error {
	const maxRetries = 3
	var lastErr error

	for i := 0; i < maxRetries; i++ {
		err := s.supabaseClient.BulkInsert(table, data)
		if err == nil {
			return nil
		}

		lastErr = err
		log.Printf("Bulk insert attempt %d/%d failed for table %s: %v", i+1, maxRetries, table, err)

		// 재시도 전 잠시 대기
		if i < maxRetries-1 {
			time.Sleep(time.Duration(i+1) * time.Second)
		}
	}

	return fmt.Errorf("bulk insert failed after %d retries: %v", maxRetries, lastErr)
}

// 재시도 로직이 포함된 벌크 업서트
func (s *Qoo10JPOrderService) bulkUpsertWithRetry(table string, data interface{}, conflictColumns ...string) error {
	const maxRetries = 3
	var lastErr error

	for i := 0; i < maxRetries; i++ {
		err := s.supabaseClient.BulkUpsert(table, data, conflictColumns...)
		if err == nil {
			return nil
		}

		lastErr = err
		log.Printf("Bulk upsert attempt %d/%d failed for table %s: %v", i+1, maxRetries, table, err)

		// 재시도 전 잠시 대기
		if i < maxRetries-1 {
			time.Sleep(time.Duration(i+1) * time.Second)
		}
	}

	return fmt.Errorf("bulk upsert failed after %d retries: %v", maxRetries, lastErr)
}

// 개별 주문 삽입 (벌크 실패 시 폴백)
func (s *Qoo10JPOrderService) insertOrdersIndividually(orders []models.Qoo10JPOrder) {
	log.Printf("Falling back to individual insert for %d orders", len(orders))

	for _, order := range orders {
		err := s.supabaseClient.Insert("orders_qoo10jp", order)
		if err != nil {
			log.Printf("Failed to insert order %s individually: %v", order.OrderNo, err)
		}
	}
}

// 개별 주문 업데이트 (벌크 실패 시 폴백)
func (s *Qoo10JPOrderService) updateOrdersIndividually(orders []models.Qoo10JPOrder) {
	log.Printf("Falling back to individual update for %d orders", len(orders))

	for _, order := range orders {
		updateQuery := fmt.Sprintf("order_no=eq.%s", order.OrderNo)
		err := s.supabaseClient.Update("orders_qoo10jp", updateQuery, order)
		if err != nil {
			log.Printf("Failed to update order %s individually: %v", order.OrderNo, err)
		}
	}
}

// 개별 주문 상품 삽입 (벌크 실패 시 폴백)
func (s *Qoo10JPOrderService) insertOrderItemsIndividually(items []models.Qoo10JPOrderItem) {
	log.Printf("Falling back to individual insert for %d order items", len(items))

	for _, item := range items {
		err := s.supabaseClient.Insert("order_items_qoo10jp", item)
		if err != nil {
			log.Printf("Failed to insert order item %s individually: %v", item.ID, err)
		}
	}
}
