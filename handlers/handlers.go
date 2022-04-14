package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	shortener "github.com/url-shortner/shortner"
	"github.com/url-shortner/store"
)

// UrlCreationRequest request object for url-shortner
type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
}

// appendHTTPSIfNotExists add https header if not present
func appendHTTPSIfNotExists(longURL string) string {
	if !strings.Contains(longURL, "https://") || !strings.Contains(longURL, "http://") {
		return "https://" + longURL
	}
	return longURL
}

// CreateShortUrl creates short URL from long URL coming from request and stores in text file
func CreateShortUrl(c *gin.Context, port string) {
	var creationRequest UrlCreationRequest
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	creationRequest.LongUrl = appendHTTPSIfNotExists(creationRequest.LongUrl)

	shortUrl := shortener.GenerateShortLink(creationRequest.LongUrl)
	store.SaveUrlMapping(shortUrl, creationRequest.LongUrl)

	c.JSON(200, gin.H{
		"message":   "short url created successfully",
		"short_url": "http://localhost:" + port + "/" + shortUrl,
	})
}

// HandleShortUrlRedirect retrieves ShortURL from store and redirects to long URL
func HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	initialUrl, err := store.RetrieveInitialUrl(shortUrl)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Redirect(302, initialUrl)
}
