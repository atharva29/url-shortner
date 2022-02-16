package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	handler "github.com/url-shortner/handlers"
	"github.com/url-shortner/store"
)

func main() {
	// Note store initialization happens here
	store.InitializeStore()

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the URL Shortener API",
		})
	})
 
	r.POST("/create-short-url", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})

	r.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})

	err := r.Run(":8100")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}

}
