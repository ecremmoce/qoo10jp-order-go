package api

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"qoo10jp-order-go/internal/models"
	"qoo10jp-order-go/internal/services"

	"github.com/gin-gonic/gin"
)

// SetupShopeeRoutes sets up Shopee-related API routes
func SetupShopeeRoutes(r *gin.Engine, shopeeOrderService *services.ShopeeOrderService) {
	shopeeGroup := r.Group("/api/v1/shopee")
	{
		// Order collection endpoint
		shopeeGroup.POST("/orders/collect", func(c *gin.Context) {
			collectShopeeOrders(c, shopeeOrderService)
		})

		// Get orders with filtering
		shopeeGroup.GET("/orders", func(c *gin.Context) {
			getShopeeOrders(c, shopeeOrderService)
		})

		// Get specific order by order SN
		shopeeGroup.GET("/orders/:order_sn", func(c *gin.Context) {
			getShopeeOrderByOrderSN(c, shopeeOrderService)
		})
	}
}

// collectShopeeOrders collects orders from Shopee API
func collectShopeeOrders(c *gin.Context, service *services.ShopeeOrderService) {
	var req struct {
		ShopID      int64  `json:"shop_id" binding:"required"`
		PartnerID   int64  `json:"partner_id" binding:"required"`
		AccessToken string `json:"access_token" binding:"required"`
		StartDate   string `json:"start_date" binding:"required"` // YYYY-MM-DD
		EndDate     string `json:"end_date" binding:"required"`   // YYYY-MM-DD
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("❌ Invalid request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "invalid_request",
			"message": err.Error(),
		})
		return
	}

	// Parse dates
	startDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "invalid_start_date",
			"message": "start_date must be in YYYY-MM-DD format",
		})
		return
	}

	endDate, err := time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "invalid_end_date",
			"message": "end_date must be in YYYY-MM-DD format",
		})
		return
	}

	// Set time to end of day for endDate
	endDate = endDate.Add(23*time.Hour + 59*time.Minute + 59*time.Second)

	log.Printf("📦 Shopee 주문 수집 요청 (Shop ID: %d, 기간: %s ~ %s)",
		req.ShopID, startDate.Format("2006-01-02"), endDate.Format("2006-01-02"))

	// Collect orders
	err = service.CollectOrders(startDate, endDate, req.ShopID, req.PartnerID, req.AccessToken)
	if err != nil {
		log.Printf("❌ 주문 수집 실패: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "collection_failed",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "주문 수집 완료",
		"shop_id":    req.ShopID,
		"start_date": req.StartDate,
		"end_date":   req.EndDate,
	})
}

// getShopeeOrders retrieves orders with filtering
func getShopeeOrders(c *gin.Context, service *services.ShopeeOrderService) {
	// Build filter from query parameters
	filter := models.ShopeeOrderFilter{
		PlatformAccountID: c.Query("platform_account_id"),
		OrderStatus:       c.Query("order_status"),
		BuyerUsername:     c.Query("buyer_username"),
		OrderSN:           c.Query("order_sn"),
	}

	// Parse dates
	if startDateStr := c.Query("start_date"); startDateStr != "" {
		startDate, err := time.Parse("2006-01-02", startDateStr)
		if err == nil {
			filter.StartDate = &startDate
		}
	}

	if endDateStr := c.Query("end_date"); endDateStr != "" {
		endDate, err := time.Parse("2006-01-02", endDateStr)
		if err == nil {
			// Set to end of day
			endDate = endDate.Add(23*time.Hour + 59*time.Minute + 59*time.Second)
			filter.EndDate = &endDate
		}
	}

	// Parse pagination
	if limitStr := c.Query("limit"); limitStr != "" {
		if limit, err := strconv.Atoi(limitStr); err == nil {
			filter.Limit = limit
		}
	}

	if offsetStr := c.Query("offset"); offsetStr != "" {
		if offset, err := strconv.Atoi(offsetStr); err == nil {
			filter.Offset = offset
		}
	}

	// Default limit
	if filter.Limit == 0 {
		filter.Limit = 50
	}

	log.Printf("🔍 Shopee 주문 조회 (필터: %+v)", filter)

	// Get orders
	orders, err := service.GetOrders(filter)
	if err != nil {
		log.Printf("❌ 주문 조회 실패: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "query_failed",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"count":  len(orders),
		"orders": orders,
		"filter": filter,
	})
}

// getShopeeOrderByOrderSN retrieves a specific order by order SN
func getShopeeOrderByOrderSN(c *gin.Context, service *services.ShopeeOrderService) {
	orderSN := c.Param("order_sn")
	if orderSN == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "invalid_order_sn",
			"message": "order_sn is required",
		})
		return
	}

	log.Printf("🔍 Shopee 주문 조회 (OrderSN: %s)", orderSN)

	order, err := service.GetOrderByOrderSN(orderSN)
	if err != nil {
		log.Printf("❌ 주문 조회 실패: %v", err)

		if err.Error() == "order not found" {
			c.JSON(http.StatusNotFound, gin.H{
				"error":   "order_not_found",
				"message": fmt.Sprintf("Order with SN %s not found", orderSN),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "query_failed",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, order)
}
