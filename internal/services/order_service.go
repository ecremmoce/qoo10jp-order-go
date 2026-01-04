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

type OrderService struct {
	supabaseClient *supabase.Client
	redisClient    *redis.Client
	qoo10jpClient  *qoo10jp.Client
}

func NewOrderService(supabaseClient *supabase.Client, redisClient *redis.Client) *OrderService {
	return &OrderService{
		supabaseClient: supabaseClient,
		redisClient:    redisClient,
	}
}

func (s *OrderService) SetQoo10JPClient(client *qoo10jp.Client) {
	s.qoo10jpClient = client
}

func (s *OrderService) CollectOrders(startDate, endDate time.Time) error {
	if s.qoo10jpClient == nil {
		return fmt.Errorf("qoo10jp client not initialized")
	}

	page := 1
	pageSize := 100

	for {
		// Check if we've already processed this page recently
		cacheKey := fmt.Sprintf("orders:collected:%s:%s:page:%d",
			startDate.Format("2006-01-02"), endDate.Format("2006-01-02"), page)

		exists, err := s.redisClient.Exists(cacheKey)
		if err != nil {
			log.Printf("Redis error: %v", err)
		} else if exists {
			log.Printf("Page %d already processed recently, skipping", page)
			page++
			continue
		}

		// Fetch orders from Qoo10JP API
		orderResp, err := s.qoo10jpClient.GetOrders(startDate, endDate, page, pageSize)
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

		// Convert and save orders
		for _, qooOrder := range orderResp.ResultObject {
			order, err := s.convertQoo10JPOrder(qooOrder)
			if err != nil {
				log.Printf("Failed to convert order %d: %v", qooOrder.OrderNo, err)
				continue
			}

			// Check if order already exists
			var existingOrders []models.Order
			query := fmt.Sprintf("order_no=eq.%s", order.OrderNo)
			err = s.supabaseClient.Select("orders", query, &existingOrders)
			if err != nil {
				log.Printf("Failed to check existing order: %v", err)
				continue
			}

			if len(existingOrders) > 0 {
				log.Printf("Order %s already exists, updating", order.OrderNo)
				updateQuery := fmt.Sprintf("order_no=eq.%s", order.OrderNo)
				err = s.supabaseClient.Update("orders", updateQuery, order)
			} else {
				log.Printf("Inserting new order %s", order.OrderNo)
				err = s.supabaseClient.Insert("orders", order)
			}

			if err != nil {
				log.Printf("Failed to save order %s: %v", order.OrderNo, err)
				continue
			}

			// Save order items
			for _, item := range order.Items {
				item.OrderID = order.ID
				err = s.supabaseClient.Insert("order_items", item)
				if err != nil {
					log.Printf("Failed to save order item: %v", err)
				}
			}
		}

		// Cache this page as processed for 1 hour
		s.redisClient.Set(cacheKey, "processed", time.Hour)

		log.Printf("Processed page %d with %d orders", page, len(orderResp.ResultObject))

		// If we got less than pageSize, we're done
		if len(orderResp.ResultObject) < pageSize {
			break
		}

		page++
	}

	return nil
}

func (s *OrderService) GetOrders(filter models.OrderFilter) ([]models.Order, error) {
	query := s.buildQuery(filter)

	var orders []models.Order
	err := s.supabaseClient.Select("orders", query, &orders)
	if err != nil {
		return nil, err
	}

	// Load order items for each order
	for i := range orders {
		var items []models.OrderItem
		itemQuery := fmt.Sprintf("order_id=eq.%s", orders[i].ID)
		err = s.supabaseClient.Select("order_items", itemQuery, &items)
		if err != nil {
			log.Printf("Failed to load items for order %s: %v", orders[i].ID, err)
		} else {
			orders[i].Items = items
		}
	}

	return orders, nil
}

func (s *OrderService) convertQoo10JPOrder(qooOrder qoo10jp.Order) (*models.Order, error) {
	orderDate, err := time.Parse("2006-01-02 15:04:05", qooOrder.OrderDate)
	if err != nil {
		return nil, err
	}

	order := &models.Order{
		ID:              fmt.Sprintf("qoo10jp_%d", qooOrder.OrderNo),
		OrderNo:         fmt.Sprintf("%d", qooOrder.OrderNo),
		OrderDate:       orderDate,
		CustomerID:      qooOrder.SellerID,
		CustomerName:    qooOrder.Buyer,
		CustomerEmail:   qooOrder.BuyerEmail,
		CustomerPhone:   qooOrder.BuyerMobile,
		TotalAmount:     qooOrder.Total,
		PaymentStatus:   qooOrder.PaymentMethod,
		OrderStatus:     qooOrder.ShippingStatus,
		ShippingAddress: qooOrder.ShippingAddress,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	// Create single order item from the order data
	item := models.OrderItem{
		ID:          fmt.Sprintf("%s_%s", order.ID, qooOrder.ItemNo),
		OrderID:     order.ID,
		ProductID:   qooOrder.ItemNo,
		ProductName: qooOrder.ItemTitle,
		Quantity:    qooOrder.OrderQty,
		Price:       qooOrder.OrderPrice,
		TotalPrice:  qooOrder.Total,
	}
	order.Items = append(order.Items, item)

	return order, nil
}

func (s *OrderService) buildQuery(filter models.OrderFilter) string {
	var conditions []string

	if filter.StartDate != nil {
		conditions = append(conditions, fmt.Sprintf("order_date=gte.%s", filter.StartDate.Format("2006-01-02")))
	}

	if filter.EndDate != nil {
		conditions = append(conditions, fmt.Sprintf("order_date=lte.%s", filter.EndDate.Format("2006-01-02")))
	}

	if filter.Status != "" {
		conditions = append(conditions, fmt.Sprintf("order_status=eq.%s", filter.Status))
	}

	query := ""
	if len(conditions) > 0 {
		query = fmt.Sprintf("%s&", fmt.Sprintf("%s", conditions[0]))
		for i := 1; i < len(conditions); i++ {
			query += fmt.Sprintf("%s&", conditions[i])
		}
	}

	if filter.Limit > 0 {
		query += fmt.Sprintf("limit=%d&", filter.Limit)
	}

	if filter.Offset > 0 {
		query += fmt.Sprintf("offset=%d&", filter.Offset)
	}

	// Remove trailing &
	if len(query) > 0 && query[len(query)-1] == '&' {
		query = query[:len(query)-1]
	}

	return query
}
