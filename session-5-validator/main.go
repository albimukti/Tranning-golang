package main

import (
	"log"

	"github.com/albimukti/Tranning-golang/session-5-validator/entity"
	"github.com/albimukti/Tranning-golang/session-5-validator/handler"
	"github.com/albimukti/Tranning-golang/session-5-validator/repository/slice"
	"github.com/albimukti/Tranning-golang/session-5-validator/router"
	"github.com/albimukti/Tranning-golang/session-5-validator/service"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// setup service
	var mockUserDBInSlice []entity.User
	userRepo := slice.NewUserRepository(mockUserDBInSlice)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Routes
	router.SetupRouter(r, userHandler)

	// Run the server
	log.Println("Running server on port 8080")
	r.Run(":8080")
}
