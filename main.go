package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// ✅ Release mode for Gin (no debug logs)
	gin.SetMode(gin.ReleaseMode)

	// Create router without default Logger middleware (optional for max speed)
	r := gin.New()
	r.Use(gin.Recovery()) // keep only Recovery middleware

	// Routes
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, Go API!"})
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// HTTP server with Keep-Alive and timeouts
	srv := &http.Server{
		Addr:           "0.0.0.0:8080", // accessible from any IP
		Handler:        r,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		IdleTimeout:    120 * time.Second, // Keep-Alive connections
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Println("🚀 Server running at http://0.0.0.0:8080")

	// Start server
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Printf("Server error: %v\n", err)
	}
}
