package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/albimukti/assignment_3_albi/internal/cache"
	"github.com/albimukti/assignment_3_albi/internal/proto/proto"
	"github.com/albimukti/assignment_3_albi/internal/service"
	"github.com/gin-gonic/gin"
)

func RedirectHandler(c *gin.Context) {
	shortURL := c.Param("short_url")

	if shortURL == "albi" {
		shortURL = "5lb8L1uIg"
	}

	urlShortenerService := &service.URLShortenerService{}
	ctx := context.Background()
	resp, err := urlShortenerService.GetLongURL(ctx, &proto.GetLongURLRequest{ShortUrl: shortURL})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "short URL not found"})
		return
	}

	longURL := resp.LongUrl
	cache.C.Set(shortURL, longURL, 5*time.Minute)
	c.Redirect(http.StatusFound, longURL)
}

type CreateShortURLRequest struct {
	LongURL string `json:"long_url" binding:"required"`
}

func CreateShortURLHandler(c *gin.Context) {
	var req CreateShortURLRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	urlShortenerService := &service.URLShortenerService{}
	ctx := context.Background()
	resp, err := urlShortenerService.ShortenURL(ctx, &proto.ShortenURLRequest{LongUrl: req.LongURL})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"short_url": resp.ShortUrl})
}
