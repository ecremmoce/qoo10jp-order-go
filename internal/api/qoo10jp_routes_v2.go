package api

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"qoo10jp-order-go/internal/models"
	"qoo10jp-order-go/internal/services"

	"github.com/gin-gonic/gin"
)

// SetupQoo10JPRoutesV2 sets up V2 Qoo10JP API routes
func SetupQoo10JPRoutesV2(router *gin.Engine, orderService *services.Qoo10JPOrderServiceV2) {
	v2 := router.Group("/api/v2/qoo10jp")
	{
		// Health check
		v2.GET("/health", qoo10jpV2HealthCheck)

		// Shop management
		v2.GET("/shops", getActiveShopsV2(orderService))
		v2.GET("/shops/:seller_id", getShopBySellerIDV2(orderService))

		// Order collection
		v2.POST("/orders/collect", collectQoo10JPOrdersV2(orderService))
		v2.POST("/orders/collect/:seller_id", collectOrdersForShopV2(orderService))

		// Order queries
		v2.GET("/orders", getQoo10JPOrdersV2(orderService))
		v2.GET("/orders/stats", getQoo10JPOrderStatsV2(orderService))
		v2.GET("/orders/:order_no", getQoo10JPOrderByOrderNoV2(orderService))
	}
}

// qoo10jpV2HealthCheck returns health status
func qoo10jpV2HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"version": "v2",
		"service": "qoo10jp-order-service-v2",
		"time":    time.Now().Format(time.RFC3339),
	})
}

// getActiveShopsV2 returns all active shops
func getActiveShopsV2(orderService *services.Qoo10JPOrderServiceV2) gin.HandlerFunc {
	return func(c *gin.Context) {
		shops, err := orderService.GetActiveShops()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data": gin.H{
				"shops": shops,
				"count": len(shops),
			},
		})
	}
}

// getShopBySellerIDV2 returns a shop by seller ID
func getShopBySellerIDV2(orderService *services.Qoo10JPOrderServiceV2) gin.HandlerFunc {
	return func(c *gin.Context) {
		sellerID := c.Param("seller_id")
		if sellerID == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error":   "seller_id is required",
			})
			return
		}

		shop, err := orderService.GetShopBySellerID(sellerID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    shop,
		})
	}
}

// collectQoo10JPOrdersV2 collects orders for all active shops
func collectQoo10JPOrdersV2(orderService *services.Qoo10JPOrderServiceV2) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse request body
		var req struct {
			StartDate string `json:"start_date"` // YYYY-MM-DD format
			EndDate   string `json:"end_date"`   // YYYY-MM-DD format
			Days      int    `json:"days"`       // Alternative: collect last N days
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			// Allow empty body, use defaults
		}

		// Calculate date range
		var startDate, endDate time.Time
		if req.StartDate != "" && req.EndDate != "" {
			var err error
			startDate, err = time.Parse("2006-01-02", req.StartDate)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"success": false,
					"error":   "Invalid start_date format. Use YYYY-MM-DD",
				})
				return
			}
			endDate, err = time.Parse("2006-01-02", req.EndDate)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"success": false,
					"error":   "Invalid end_date format. Use YYYY-MM-DD",
				})
				return
			}
		} else {
			days := req.Days
			if days <= 0 {
				days = 30 // Default: 30 days
			}
			endDate = time.Now()
			startDate = endDate.AddDate(0, 0, -days)
		}

		// Get all active shops
		shops, err := orderService.GetActiveShops()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"error":   fmt.Sprintf("Failed to get active shops: %v", err),
			})
			return
		}

		if len(shops) == 0 {
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"message": "No active shops found",
				"data": gin.H{
					"results":       []interface{}{},
					"total_shops":   0,
					"start_date":    startDate.Format("2006-01-02"),
					"end_date":      endDate.Format("2006-01-02"),
				},
			})
			return
		}

		// Collect orders for each shop
		var results []*models.Qoo10JPCollectionResult
		totalCollected := 0
		totalSaved := 0
		totalUpdated := 0
		successCount := 0
		failedCount := 0

		for _, shop := range shops {
			result, err := orderService.CollectOrdersForShop(shop.SellerID, startDate, endDate)
			if err != nil {
				failedCount++
				results = append(results, &models.Qoo10JPCollectionResult{
					SellerID: shop.SellerID,
					Errors:   []string{err.Error()},
				})
				continue
			}

			successCount++
			totalCollected += result.TotalCollected
			totalSaved += result.TotalSaved
			totalUpdated += result.TotalUpdated
			results = append(results, result)
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data": gin.H{
				"results":         results,
				"total_shops":     len(shops),
				"successful":      successCount,
				"failed":          failedCount,
				"total_collected": totalCollected,
				"total_saved":     totalSaved,
				"total_updated":   totalUpdated,
				"start_date":      startDate.Format("2006-01-02"),
				"end_date":        endDate.Format("2006-01-02"),
			},
		})
	}
}

// collectOrdersForShopV2 collects orders for a specific shop
func collectOrdersForShopV2(orderService *services.Qoo10JPOrderServiceV2) gin.HandlerFunc {
	return func(c *gin.Context) {
		sellerID := c.Param("seller_id")
		if sellerID == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error":   "seller_id is required",
			})
			return
		}

		// Parse request body
		var req struct {
			StartDate string `json:"start_date"`
			EndDate   string `json:"end_date"`
			Days      int    `json:"days"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			// Allow empty body
		}

		// Calculate date range
		var startDate, endDate time.Time
		if req.StartDate != "" && req.EndDate != "" {
			var err error
			startDate, err = time.Parse("2006-01-02", req.StartDate)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"success": false,
					"error":   "Invalid start_date format",
				})
				return
			}
			endDate, err = time.Parse("2006-01-02", req.EndDate)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"success": false,
					"error":   "Invalid end_date format",
				})
				return
			}
		} else {
			days := req.Days
			if days <= 0 {
				days = 30
			}
			endDate = time.Now()
			startDate = endDate.AddDate(0, 0, -days)
		}

		// Collect orders
		result, err := orderService.CollectOrdersForShop(sellerID, startDate, endDate)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    result,
		})
	}
}

// getQoo10JPOrdersV2 returns orders with filter
func getQoo10JPOrdersV2(orderService *services.Qoo10JPOrderServiceV2) gin.HandlerFunc {
	return func(c *gin.Context) {
		filter := models.Qoo10JPOrderFilter{
			SellerID: c.Query("seller_id"),
			Status:   c.Query("status"),
		}

		// Parse dates
		if startDateStr := c.Query("start_date"); startDateStr != "" {
			if t, err := time.Parse("2006-01-02", startDateStr); err == nil {
				filter.StartDate = &t
			}
		}
		if endDateStr := c.Query("end_date"); endDateStr != "" {
			if t, err := time.Parse("2006-01-02", endDateStr); err == nil {
				filter.EndDate = &t
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

		orders, err := orderService.GetOrdersV2(filter)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data": gin.H{
				"orders": orders,
				"count":  len(orders),
			},
		})
	}
}

// getQoo10JPOrderStatsV2 returns order statistics
func getQoo10JPOrderStatsV2(orderService *services.Qoo10JPOrderServiceV2) gin.HandlerFunc {
	return func(c *gin.Context) {
		sellerID := c.Query("seller_id")
		if sellerID == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error":   "seller_id is required",
			})
			return
		}

		days := 30
		if daysStr := c.Query("days"); daysStr != "" {
			if d, err := strconv.Atoi(daysStr); err == nil && d > 0 {
				days = d
			}
		}

		stats, err := orderService.GetOrderStatsV2(sellerID, days)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    stats,
		})
	}
}

// getQoo10JPOrderByOrderNoV2 returns a specific order by order number
func getQoo10JPOrderByOrderNoV2(orderService *services.Qoo10JPOrderServiceV2) gin.HandlerFunc {
	return func(c *gin.Context) {
		orderNo := c.Param("order_no")
		if orderNo == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error":   "order_no is required",
			})
			return
		}

		// Use filter with order_no (this would need to be added to GetOrdersV2)
		filter := models.Qoo10JPOrderFilter{}
		orders, err := orderService.GetOrdersV2(filter)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"error":   err.Error(),
			})
			return
		}

		// Find the specific order
		for _, order := range orders {
			if order.OrderNo == orderNo {
				c.JSON(http.StatusOK, gin.H{
					"success": true,
					"data":    order,
				})
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Order not found",
		})
	}
}
