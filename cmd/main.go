package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"qoo10jp-order-go/internal/api"
	"qoo10jp-order-go/internal/config"
	"qoo10jp-order-go/internal/services"
	"qoo10jp-order-go/pkg/redis"
	"qoo10jp-order-go/pkg/supabase"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Load configuration
	cfg := config.Load()

	// Initialize Redis
	redisClient, err := redis.NewClient(cfg.Redis)
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	defer redisClient.Close()

	// Initialize Supabase
	supabaseClient, err := supabase.NewClient(cfg.Supabase)
	if err != nil {
		log.Fatalf("Failed to connect to Supabase: %v", err)
	}

	// Initialize services
	orderService := services.NewOrderService(supabaseClient, redisClient)
	qoo10jpOrderService := services.NewQoo10JPOrderService(supabaseClient, redisClient)
	schedulerService := services.NewSchedulerService(redisClient, supabaseClient, orderService)
	workerService := services.NewWorkerService(schedulerService, cfg.Worker.Count)

	// Auto-start worker service
	log.Println("Auto-starting worker service...")
	workerService.Start()

	// Initialize API routes
	router := gin.Default()
	
	// Serve static files (admin panel)
	router.Static("/web", "./web")
	router.GET("/", func(c *gin.Context) {
		c.File("./web/admin.html")
	})
	
	api.SetupRoutes(router, orderService)
	api.SetupQoo10JPRoutes(router, qoo10jpOrderService)
	api.SetupSchedulerRoutes(router, schedulerService, workerService)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Create HTTP server
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	// Start server in goroutine
	go func() {
		log.Printf("Server starting on port %s", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Stop worker service
	log.Println("Stopping worker service...")
	workerService.Stop()

	// Shutdown server with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited")
}
