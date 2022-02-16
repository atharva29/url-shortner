package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	shortener "github.com/url-shortner/shortner"
	"github.com/url-shortner/store"
)

type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
}

func appendHTTPSIfNotExists(longURL string) string {
	if !strings.Contains(longURL, "https://") {
		return "https://" + longURL
	}
	return longURL
}

func CreateShortUrl(c *gin.Context) {
	var creationRequest UrlCreationRequest
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	creationRequest.LongUrl = appendHTTPSIfNotExists(creationRequest.LongUrl)

	shortUrl := shortener.GenerateShortLink(creationRequest.LongUrl)
	store.SaveUrlMapping(shortUrl, creationRequest.LongUrl)

	host := "http://localhost:8100/"
	c.JSON(200, gin.H{
		"message":   "short url created successfully",
		"short_url": host + shortUrl,
	})
}

func HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	initialUrl, err := store.RetrieveInitialUrl(shortUrl)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Redirect(302, initialUrl)
}
