package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	handler "github.com/url-shortner/handlers"
	"github.com/url-shortner/store"
	"github.com/url-shortner/utils"
)

func main() {
	// Store initialization happens here
	store.InitializeStore()

	// Gin HTTP server intialization
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the URL Shortener API",
		})
	})

	// Endpoint for creating short URL
	r.POST("/create-short-url", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})

	// Endpoint for accessing short URLs
	r.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})

	err := r.Run(utils.HTTPPort)
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}

}
