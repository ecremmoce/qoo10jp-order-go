package api

import (
	"net/http"
	"qoo10jp-order-go/internal/services"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupSchedulerRoutes(router *gin.Engine, schedulerService *services.SchedulerService, workerService *services.WorkerService) {
	api := router.Group("/api/v1/scheduler")
	{
		api.POST("/job", createJob(schedulerService))
		api.GET("/status", getSchedulerStatus(schedulerService, workerService))
		api.POST("/worker/start", startWorker(workerService))
		api.POST("/worker/stop", stopWorker(workerService))
		api.POST("/worker/count", setWorkerCount(workerService))
		api.POST("/schedule-next", scheduleNext(schedulerService))
	}

	// Test endpoint for Redis queue
	router.POST("/api/test-redis-push", testRedisPush(schedulerService))
}

func createJob(schedulerService *services.SchedulerService) gin.HandlerFunc {
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

		if err := schedulerService.CreateOrderJob(startDate, endDate); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message":    "Job created successfully",
			"start_date": req.StartDate,
			"end_date":   req.EndDate,
		})
	}
}

func getSchedulerStatus(schedulerService *services.SchedulerService, workerService *services.WorkerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		status := workerService.GetStatus()
		c.JSON(http.StatusOK, status)
	}
}

func startWorker(workerService *services.WorkerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		if workerService.IsRunning() {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Worker service is already running"})
			return
		}

		workerService.Start()
		c.JSON(http.StatusOK, gin.H{"message": "Worker service started"})
	}
}

func stopWorker(workerService *services.WorkerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !workerService.IsRunning() {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Worker service is not running"})
			return
		}

		workerService.Stop()
		c.JSON(http.StatusOK, gin.H{"message": "Worker service stopped"})
	}
}

func setWorkerCount(workerService *services.WorkerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Count int `json:"count" binding:"required,min=1,max=20"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := workerService.SetWorkerCount(req.Count); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Worker count updated successfully",
			"count":   req.Count,
		})
	}
}

func scheduleNext(schedulerService *services.SchedulerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := schedulerService.ScheduleNextJob(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Next job scheduled successfully"})
	}
}

func testRedisPush(schedulerService *services.SchedulerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Message string `json:"message" binding:"required"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Push message to the worker queue
		queueName := "shopee_order_queue"
		if err := schedulerService.GetRedisClient().PushToQueue(queueName, req.Message); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Test message pushed to Redis queue successfully",
			"queue":   queueName,
		})
	}
}
