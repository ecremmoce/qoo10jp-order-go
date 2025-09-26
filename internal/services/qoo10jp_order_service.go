package services

import (
	"fmt"
	"log"
	"qoo10jp-order-go/internal/models"
	"qoo10jp-order-go/pkg/qoo10jp"
	"qoo10jp-order-go/pkg/redis"
	"qoo10jp-order-go/pkg/supabase"
	"time"
)

type Qoo10JPOrderService struct {
	supabaseClient *supabase.Client
	redisClient    *redis.Client
}

func NewQoo10JPOrderService(supabaseClient *supabase.Client, redisClient *redis.Client) *Qoo10JPOrderService {
	return &Qoo10JPOrderService{
		supabaseClient: supabaseClient,
		redisClient:    redisClient,
	}
}

// Supabase에서 계정 정보를 가져와서 Qoo10JP 클라이언트 생성
func (s *Qoo10JPOrderService) createQoo10JPClient(platformAccountID string) (*qoo10jp.Client, error) {
	// Supabase에서 계정 정보 조회
	var accounts []struct {
		ID              string `json:"id"`
		SellerID        string `json:"seller_id"`
		APIId           string `json:"api_id"`
		CertificationKey string `json:"certification_key"`
		IsActive        bool   `json:"is_active"`
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
	
	// Qoo10JP 클라이언트 생성
	baseURL := "https://api.qoo10.com" // 올바른 URL
	return qoo10jp.NewClient(account.APIId, account.CertificationKey, baseURL), nil
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

	for {
		// 캐시 키 생성
		cacheKey := fmt.Sprintf("qoo10jp:orders:collected:%s:%s:account:%s:page:%d",
			startDate.Format("2006-01-02"), endDate.Format("2006-01-02"), platformAccountID, page)

		// 이미 처리된 페이지인지 확인
		exists, err := s.redisClient.Exists(cacheKey)
		if err != nil {
			log.Printf("Redis error: %v", err)
		} else if exists {
			log.Printf("Page %d already processed recently, skipping", page)
			page++
			continue
		}

		// Qoo10JP API에서 주문 데이터 가져오기
		orderResp, err := qoo10jpClient.GetOrders(startDate, endDate, page, pageSize)
		if err != nil {
			return fmt.Errorf("failed to fetch orders: %v", err)
		}

		if orderResp.ResultCode != "0" {
			return fmt.Errorf("API error: %s", orderResp.ResultMessage)
		}

		if len(orderResp.ResultObject) == 0 {
			log.Println("No more orders to process")
			break
		}

		// 주문 데이터 변환 및 저장
		for _, qooOrder := range orderResp.ResultObject {
			order, err := s.convertQoo10JPOrder(qooOrder, platformAccountID)
			if err != nil {
				log.Printf("Failed to convert order %s: %v", qooOrder.OrderNo, err)
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
				// 기존 주문 업데이트
				log.Printf("Order %s already exists, updating", order.OrderNo)
				updateQuery := fmt.Sprintf("order_no=eq.%s", order.OrderNo)
				err = s.supabaseClient.Update("orders_qoo10jp", updateQuery, order)
			} else {
				// 새 주문 삽입
				log.Printf("Inserting new order %s", order.OrderNo)
				err = s.supabaseClient.Insert("orders_qoo10jp", order)
			}

			if err != nil {
				log.Printf("Failed to save order %s: %v", order.OrderNo, err)
				continue
			}

			// 주문 상품 저장
			for _, item := range order.Items {
				item.OrderID = order.ID
				
				// 기존 주문 상품 삭제 (업데이트의 경우)
				if len(existingOrders) > 0 {
					deleteQuery := fmt.Sprintf("order_id=eq.%s", order.ID)
					s.supabaseClient.Delete("order_items_qoo10jp", deleteQuery)
				}
				
				err = s.supabaseClient.Insert("order_items_qoo10jp", item)
				if err != nil {
					log.Printf("Failed to save order item: %v", err)
				}
			}
		}

		// 처리 완료된 페이지를 캐시에 저장 (1시간)
		s.redisClient.Set(cacheKey, "processed", time.Hour)

		log.Printf("Processed page %d with %d orders", page, len(orderResp.ResultObject))

		// 페이지 크기보다 적으면 마지막 페이지
		if len(orderResp.ResultObject) < pageSize {
			break
		}

		page++
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
		ID:                fmt.Sprintf("qoo10jp_%s", qooOrder.OrderNo),
		OrderNo:           qooOrder.OrderNo,
		OrderDate:         orderDate,
		OrderStatus:       qooOrder.OrderStatus,
		PaymentStatus:     qooOrder.PaymentStatus,
		BuyerID:           qooOrder.BuyerID,
		BuyerName:         qooOrder.BuyerName,
		BuyerEmail:        qooOrder.BuyerEmail,
		BuyerPhone:        qooOrder.BuyerPhone,
		TotalAmount:       qooOrder.TotalAmount,
		PaymentAmount:     qooOrder.TotalAmount, // API에서 별도 제공되지 않으면 동일하게 설정
		Currency:          "JPY",
		ShippingAddress:   qooOrder.ShipAddress,
		PlatformAccountID: platformAccountID,
		RawData:           qooOrder,
		SyncedAt:          time.Now(),
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}

	// 주문 상품 변환
	for _, qooItem := range qooOrder.OrderItems {
		item := models.Qoo10JPOrderItem{
			ID:         fmt.Sprintf("%s_%s", order.ID, qooItem.ItemCode),
			OrderID:    order.ID,
			ItemCode:   qooItem.ItemCode,
			ItemName:   qooItem.ItemName,
			Quantity:   qooItem.Quantity,
			UnitPrice:  qooItem.ItemPrice,
			TotalPrice: qooItem.TotalPrice,
			RawData:    qooItem,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}
		order.Items = append(order.Items, item)
	}

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



