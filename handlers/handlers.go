package handler

import (
	"github.com/gin-gonic/gin"
)

type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
}

func CreateShortUrl(c *gin.Context) {
}

func HandleShortUrlRedirect(c *gin.Context) {
}
