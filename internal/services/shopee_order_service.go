package services

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"qoo10jp-order-go/internal/config"
	"qoo10jp-order-go/internal/models"
	"qoo10jp-order-go/pkg/shopee"
	"qoo10jp-order-go/pkg/supabase"

	"github.com/google/uuid"
)

type ShopeeOrderService struct {
	cfg            *config.Config
	supabaseClient *supabase.Client
	shopeeClient   *shopee.Client
}

func NewShopeeOrderService(cfg *config.Config, supabaseClient *supabase.Client) *ShopeeOrderService {
	shopeeClient := shopee.NewClient(
		cfg.Shopee.BaseURL,
		cfg.Shopee.PartnerID,
		cfg.Shopee.PartnerKey,
	)

	return &ShopeeOrderService{
		cfg:            cfg,
		supabaseClient: supabaseClient,
		shopeeClient:   shopeeClient,
	}
}

// CollectOrders collects orders from Shopee API for a specific date range
func (s *ShopeeOrderService) CollectOrders(startDate, endDate time.Time, shopID int64, partnerID int64, accessToken string) error {
	log.Printf("📦 Shopee 주문 수집 시작 (Shop ID: %d, Partner ID: %d, 기간: %s ~ %s)",
		shopID, partnerID, startDate.Format("2006-01-02"), endDate.Format("2006-01-02"))
	
	// Create a temporary client with the specific partner_id for this shop
	tempClient := shopee.NewClient(
		s.cfg.Shopee.BaseURL,
		partnerID, // Use the partner_id from the message
		s.cfg.Shopee.PartnerKey,
	)

	totalCollected := 0
	cursor := ""
	hasMore := true

	for hasMore {
		// Get order list from Shopee API
		req := shopee.OrderListRequest{
			TimeRangeField: "create_time",
			TimeFrom:       startDate.Unix(),
			TimeTo:         endDate.Unix(),
			PageSize:       100, // Max page size
			Cursor:         cursor,
			ShopID:         shopID,
			AccessToken:    accessToken,
		}

		log.Printf("🔍 주문 목록 조회 중... (cursor: %s)", cursor)
		listResp, err := tempClient.GetOrderList(req)
		if err != nil {
			return fmt.Errorf("failed to get order list: %w", err)
		}

		orderCount := len(listResp.Response.OrderList)
		totalCollected += orderCount
		log.Printf("✅ %d개 주문 조회 완료", orderCount)

		if orderCount == 0 {
			log.Println("📭 더 이상 조회할 주문이 없습니다")
			break
		}

		// Extract order SNs
		orderSNs := make([]string, 0, orderCount)
		for _, order := range listResp.Response.OrderList {
			orderSNs = append(orderSNs, order.OrderSN)
		}

		// TODO: 나중에 상세 정보 조회 및 저장 구현
		// Get order details in batches (max 50 per request)
		// batchSize := 50
		// for i := 0; i < len(orderSNs); i += batchSize {
		// 	end := i + batchSize
		// 	if end > len(orderSNs) {
		// 		end = len(orderSNs)
		// 	}
		// 	batch := orderSNs[i:end]

		// 	log.Printf("📄 주문 상세 정보 조회 중... (%d-%d/%d)", i+1, end, len(orderSNs))

		// 	detailReq := shopee.OrderDetailRequest{
		// 		OrderSNList: batch,
		// 		ResponseOptionalFields: []string{
		// 			"buyer_user_id",
		// 			"buyer_username",
		// 			"recipient_address",
		// 			"item_list",
		// 		},
		// 		ShopID:      shopID,
		// 		AccessToken: accessToken,
		// 	}

		// 	detailResp, err := s.shopeeClient.GetOrderDetail(detailReq)
		// 	if err != nil {
		// 		log.Printf("⚠️ 주문 상세 정보 조회 실패: %v", err)
		// 		continue
		// 	}

		// 	// Save orders to database
		// 	for _, orderDetail := range detailResp.Response.OrderList {
		// 		err := s.saveOrder(orderDetail, fmt.Sprintf("%d", shopID))
		// 		if err != nil {
		// 			log.Printf("⚠️ 주문 저장 실패 (OrderSN: %s): %v", orderDetail.OrderSN, err)
		// 			continue
		// 		}
		// 		totalSaved++
		// 	}

		// 	log.Printf("💾 %d개 주문 저장 완료", len(detailResp.Response.OrderList))
		// }

		log.Printf("📋 주문 SN 목록 (%d개): %v", len(orderSNs), orderSNs)

		// Check if there are more pages
		hasMore = listResp.Response.More
		cursor = listResp.Response.NextCursor

		if hasMore {
			log.Printf("➡️  다음 페이지로 이동 (cursor: %s)", cursor)
			// Small delay to avoid rate limiting
			time.Sleep(500 * time.Millisecond)
		}
	}

	log.Printf("✅ Shopee 주문 수집 완료 (조회: %d건)", totalCollected)
	return nil
}

// saveOrder saves or updates an order in the database
func (s *ShopeeOrderService) saveOrder(orderDetail shopee.OrderDetail, platformAccountID string) error {
	// Convert items to JSON
	itemsJSON, err := json.Marshal(orderDetail.ItemList)
	if err != nil {
		return fmt.Errorf("failed to marshal items: %w", err)
	}

	order := models.ShopeeOrder{
		ID:                uuid.New().String(),
		OrderSN:           orderDetail.OrderSN,
		PlatformAccountID: platformAccountID,
		OrderStatus:       orderDetail.OrderStatus,
		CreateTime:        time.Unix(orderDetail.CreateTime, 0),
		UpdateTime:        time.Unix(orderDetail.UpdateTime, 0),
		BuyerUserID:       orderDetail.BuyerUserID,
		BuyerUsername:     orderDetail.BuyerUsername,
		RecipientName:     orderDetail.RecipientAddress.Name,
		RecipientPhone:    orderDetail.RecipientAddress.Phone,
		RecipientAddress:  orderDetail.RecipientAddress.FullAddress,
		RecipientDistrict: orderDetail.RecipientAddress.District,
		RecipientCity:     orderDetail.RecipientAddress.City,
		RecipientState:    orderDetail.RecipientAddress.State,
		RecipientCountry:  orderDetail.RecipientAddress.Country,
		RecipientZipcode:  orderDetail.RecipientAddress.Zipcode,
		TotalAmount:       orderDetail.TotalAmount,
		Currency:          orderDetail.Currency,
		PaymentMethod:     orderDetail.PaymentMethod,
		ShippingCarrier:   orderDetail.ShippingCarrier,
		TrackingNumber:    orderDetail.TrackingNumber,
		ItemsJSON:         string(itemsJSON),
	}

	// Use Supabase REST API Insert (with upsert using BulkUpsert for single item)
	orderData := []map[string]interface{}{
		{
			"id":                  order.ID,
			"order_sn":            order.OrderSN,
			"platform_account_id": order.PlatformAccountID,
			"order_status":        order.OrderStatus,
			"create_time":         order.CreateTime,
			"update_time":         order.UpdateTime,
			"buyer_user_id":       order.BuyerUserID,
			"buyer_username":      order.BuyerUsername,
			"recipient_name":      order.RecipientName,
			"recipient_phone":     order.RecipientPhone,
			"recipient_address":   order.RecipientAddress,
			"recipient_district":  order.RecipientDistrict,
			"recipient_city":      order.RecipientCity,
			"recipient_state":     order.RecipientState,
			"recipient_country":   order.RecipientCountry,
			"recipient_zipcode":   order.RecipientZipcode,
			"total_amount":        order.TotalAmount,
			"currency":            order.Currency,
			"payment_method":      order.PaymentMethod,
			"shipping_carrier":    order.ShippingCarrier,
			"tracking_number":     order.TrackingNumber,
			"items_json":          order.ItemsJSON,
		},
	}

	// Use BulkUpsert with order_sn as conflict column
	err = s.supabaseClient.BulkUpsert("shopee_orders", orderData, "order_sn")
	if err != nil {
		return fmt.Errorf("failed to save order: %w", err)
	}

	return nil
}

// GetOrders retrieves orders from the database with filtering
func (s *ShopeeOrderService) GetOrders(filter models.ShopeeOrderFilter) ([]models.ShopeeOrder, error) {
	queryParts := []string{}

	if filter.PlatformAccountID != "" {
		queryParts = append(queryParts, fmt.Sprintf("platform_account_id=eq.%s", filter.PlatformAccountID))
	}

	if filter.OrderStatus != "" {
		queryParts = append(queryParts, fmt.Sprintf("order_status=eq.%s", filter.OrderStatus))
	}

	if filter.StartDate != nil {
		queryParts = append(queryParts, fmt.Sprintf("create_time=gte.%s", filter.StartDate.Format(time.RFC3339)))
	}

	if filter.EndDate != nil {
		queryParts = append(queryParts, fmt.Sprintf("create_time=lte.%s", filter.EndDate.Format(time.RFC3339)))
	}

	if filter.BuyerUsername != "" {
		queryParts = append(queryParts, fmt.Sprintf("buyer_username=ilike.*%s*", filter.BuyerUsername))
	}

	if filter.OrderSN != "" {
		queryParts = append(queryParts, fmt.Sprintf("order_sn=eq.%s", filter.OrderSN))
	}

	// Order by create_time DESC
	queryParts = append(queryParts, "order=create_time.desc")

	// Limit
	limit := filter.Limit
	if limit == 0 {
		limit = 100 // Default
	}
	queryParts = append(queryParts, fmt.Sprintf("limit=%d", limit))

	// Offset
	if filter.Offset > 0 {
		queryParts = append(queryParts, fmt.Sprintf("offset=%d", filter.Offset))
	}

	query := strings.Join(queryParts, "&")

	var orders []models.ShopeeOrder
	err := s.supabaseClient.Select("shopee_orders", query, &orders)
	if err != nil {
		return nil, fmt.Errorf("failed to query orders: %w", err)
	}

	return orders, nil
}

// GetOrderByOrderSN retrieves a specific order by order SN
func (s *ShopeeOrderService) GetOrderByOrderSN(orderSN string) (*models.ShopeeOrder, error) {
	query := fmt.Sprintf("order_sn=eq.%s", orderSN)

	var orders []models.ShopeeOrder
	err := s.supabaseClient.Select("shopee_orders", query, &orders)
	if err != nil {
		return nil, fmt.Errorf("failed to query order: %w", err)
	}

	if len(orders) == 0 {
		return nil, fmt.Errorf("order not found")
	}

	return &orders[0], nil
}
