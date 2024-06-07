package router

import (
	"github.com/albimukti/Tranning-golang/session-1/handler"
	"github.com/albimukti/Tranning-golang/session-1/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	r.Use(middleware.AuthMiddleware())
	r.GET("/", handler.RootHandler)
	r.GET("/superman", handler.SupermanHandler)
	r.POST("/post", handler.PostHandler)
	// Tambahkan middleware AuthMiddleware ke rute yang memerlukan autentikasi
	privateEndpoint := r.Group("/private")
	privateEndpoint.Use(middleware.AuthMiddleware())
	{
		privateEndpoint.POST("/post", handler.PostHandler)
	}

}
