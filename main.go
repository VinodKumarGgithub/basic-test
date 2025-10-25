package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() // Logger + Recovery middleware

	// Middleware example: simple custom logging
	// r.Use(func(c *gin.Context) {
	//     fmt.Println("Request Path:", c.Request.URL.Path)
	//     c.Next()
	// })

	// Routes
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, Go API!"})
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// Start server
	fmt.Println("🚀 Server running at http://localhost:8080")
	r.Run(":8080")
}
