package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"dot-test-vandy/config"

	"dot-test-vandy/internal/setup"
	"dot-test-vandy/lib/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.NewConfig()
	log.Println("Starting application")
	srv := setup.NewServices()
	log.Println("Server has been setup")

	// Initialize Gin router
	route := gin.Default()
	route.Use(middleware.RecoveryMiddleware())

	// Set up routes
	api := route.Group("/api")
	api.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	api.GET("/category", srv.CategoryHandler.GetAll)
	api.GET("/category/:id", srv.CategoryHandler.GetByID)
	api.POST("/category", srv.CategoryHandler.Create)
	api.PUT("/category/:id", srv.CategoryHandler.Update)
	api.DELETE("/category/:id", srv.CategoryHandler.Delete)

	api.GET("/book", srv.BookHandler.GetAll)
	api.GET("/book/:id", srv.BookHandler.GetByID)
	api.POST("/book", srv.BookHandler.Create)
	api.PUT("/book/:id", srv.BookHandler.Update)
	api.DELETE("/book/:id", srv.BookHandler.Delete)

	// Create a server instance
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.PORT),
		Handler: route,
	}

	// Graceful shutdown logic
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to listen: %v\n", err)
		}
	}()

	log.Printf("Server is running on port %s\n", cfg.PORT)

	// Wait for termination signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")
}
