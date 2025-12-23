package api

import (
	"net/http"
	"shopee-order-go/internal/models"
	"shopee-order-go/internal/services"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupQoo10JPRoutes(router *gin.Engine, orderService *services.Qoo10JPOrderService) {
	qoo10jp := router.Group("/api/v1/qoo10jp")
	{
		qoo10jp.GET("/health", qoo10jpHealthCheck)
		qoo10jp.POST("/orders/collect", collectQoo10JPOrders(orderService))
		qoo10jp.GET("/orders", getQoo10JPOrders(orderService))
		qoo10jp.GET("/orders/stats", getQoo10JPOrderStats(orderService))
	}
}

func qoo10jpHealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":    "ok",
		"timestamp": time.Now().Unix(),
		"service":   "qoo10jp-order-collector",
	})
}

func collectQoo10JPOrders(orderService *services.Qoo10JPOrderService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			StartDate         string `json:"start_date" binding:"required"`
			EndDate           string `json:"end_date" binding:"required"`
			PlatformAccountID string `json:"platform_account_id" binding:"required"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		startDate, err := time.Parse("2006-01-02", req.StartDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start_date format. Use YYYY-MM-DD"})
			return
		}

		endDate, err := time.Parse("2006-01-02", req.EndDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end_date format. Use YYYY-MM-DD"})
			return
		}

		if err := orderService.CollectOrders(startDate, endDate, req.PlatformAccountID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message":             "Qoo10JP orders collection completed",
			"start_date":          req.StartDate,
			"end_date":            req.EndDate,
			"platform_account_id": req.PlatformAccountID,
		})
	}
}

func getQoo10JPOrders(orderService *services.Qoo10JPOrderService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter struct {
			StartDate         string `form:"start_date"`
			EndDate           string `form:"end_date"`
			OrderStatus       string `form:"order_status"`
			PaymentStatus     string `form:"payment_status"`
			PlatformAccountID string `form:"platform_account_id"`
			BuyerID           string `form:"buyer_id"`
			Limit             string `form:"limit"`
			Offset            string `form:"offset"`
		}

		if err := c.ShouldBindQuery(&filter); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		orderFilter := models.Qoo10JPOrderFilter{}

		if filter.StartDate != "" {
			startDate, err := time.Parse("2006-01-02", filter.StartDate)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start_date format"})
				return
			}
			orderFilter.StartDate = &startDate
		}

		if filter.EndDate != "" {
			endDate, err := time.Parse("2006-01-02", filter.EndDate)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end_date format"})
				return
			}
			orderFilter.EndDate = &endDate
		}

		if filter.OrderStatus != "" {
			orderFilter.OrderStatus = filter.OrderStatus
		}

		if filter.PaymentStatus != "" {
			orderFilter.PaymentStatus = filter.PaymentStatus
		}

		if filter.PlatformAccountID != "" {
			orderFilter.PlatformAccountID = filter.PlatformAccountID
		}

		if filter.BuyerID != "" {
			orderFilter.BuyerID = filter.BuyerID
		}

		if filter.Limit != "" {
			limit, err := strconv.Atoi(filter.Limit)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
				return
			}
			orderFilter.Limit = limit
		}

		if filter.Offset != "" {
			offset, err := strconv.Atoi(filter.Offset)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset"})
				return
			}
			orderFilter.Offset = offset
		}

		orders, err := orderService.GetOrders(orderFilter)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"orders": orders,
			"count":  len(orders),
			"filter": orderFilter,
		})
	}
}

func getQoo10JPOrderStats(orderService *services.Qoo10JPOrderService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var query struct {
			PlatformAccountID string `form:"platform_account_id"`
			StartDate         string `form:"start_date"`
			EndDate           string `form:"end_date"`
		}

		if err := c.ShouldBindQuery(&query); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var startDate, endDate *time.Time

		if query.StartDate != "" {
			sd, err := time.Parse("2006-01-02", query.StartDate)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start_date format"})
				return
			}
			startDate = &sd
		}

		if query.EndDate != "" {
			ed, err := time.Parse("2006-01-02", query.EndDate)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end_date format"})
				return
			}
			endDate = &ed
		}

		stats, err := orderService.GetOrderStats(query.PlatformAccountID, startDate, endDate)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"stats": stats,
			"query": query,
		})
	}
}
