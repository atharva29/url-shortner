package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	handler "github.com/url-shortner/handlers"
	"github.com/url-shortner/store"
	"github.com/url-shortner/utils"
)

func main() {
	port := utils.HTTPPort
	if getEnvVariable("p") != "" {
		port = getEnvVariable("p")
	}
	flag.Parse()
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
		handler.CreateShortUrl(c, port)
	})

	// Endpoint for accessing short URLs
	r.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})

	err := r.Run(":" + port)
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}

func getEnvVariable(key string) string {
	return os.Getenv(key)
}
