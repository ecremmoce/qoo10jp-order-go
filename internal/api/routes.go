package api

import (
	"net/http"
	"shopee-order-go/internal/models"
	"shopee-order-go/internal/services"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, orderService *services.OrderService) {
	api := router.Group("/api/v1")
	{
		api.GET("/health", healthCheck)
		api.POST("/orders/collect", collectOrders(orderService))
		api.GET("/orders", getOrders(orderService))
	}
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":    "ok",
		"timestamp": time.Now().Unix(),
		"service":   "qoo10jp-order-collector",
	})
}

func collectOrders(orderService *services.OrderService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			StartDate string `json:"start_date" binding:"required"`
			EndDate   string `json:"end_date" binding:"required"`
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

		if err := orderService.CollectOrders(startDate, endDate); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message":    "Orders collection completed",
			"start_date": req.StartDate,
			"end_date":   req.EndDate,
		})
	}
}

func getOrders(orderService *services.OrderService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter struct {
			StartDate string `form:"start_date"`
			EndDate   string `form:"end_date"`
			Status    string `form:"status"`
			Limit     string `form:"limit"`
			Offset    string `form:"offset"`
		}

		if err := c.ShouldBindQuery(&filter); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		orderFilter := models.OrderFilter{}

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

		if filter.Status != "" {
			orderFilter.Status = filter.Status
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
		})
	}
}
