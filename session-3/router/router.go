package router

import (
	"github.com/albimukti/Tranning-golang/session-3/handler"
	"github.com/albimukti/Tranning-golang/session-3/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	r.GET("/", handler.RootHandler)

	// Tambahkan middleware AuthMiddleware ke rute yang memerlukan autentikasi
	privateEndpoint := r.Group("/private")
	privateEndpoint.Use(middleware.AuthMiddleware())
	{
		privateEndpoint.POST("/post", handler.PostHandler)
	}
}
