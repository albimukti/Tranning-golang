package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ibrahimker/golang-praisindo-advanced/session-1-introduction-gin/handler"
	"github.com/ibrahimker/golang-praisindo-advanced/session-1-introduction-gin/middleware"
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
