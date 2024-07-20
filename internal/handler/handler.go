package handler

import (
	"net/http"
	"time"

	"github.com/albimukti/assignment_3_albi/internal/cache"
	"github.com/albimukti/assignment_3_albi/internal/database"
	"github.com/gin-gonic/gin"
)

func RedirectHandler(c *gin.Context) {
	shortURL := c.Param("short_url")

	if longURL, found := cache.C.Get(shortURL); found {
		c.Redirect(http.StatusFound, longURL.(string))
		return
	}

	var longURL string
	err := database.DB.QueryRow(c, "SELECT long_url FROM url_mappings WHERE short_url = $1", shortURL).Scan(&longURL)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "short URL not found"})
		return
	}

	cache.C.Set(shortURL, longURL, 5*time.Minute)
	c.Redirect(http.StatusFound, longURL)
}
