package handler

import (
	"net/http"

	"github.com/albimukti/Tranning-golang/session-4/langkah-4/service"
	"github.com/gin-gonic/gin"
)

// IUserHandler mendefinisikan interface untuk handler user
type IUserHandler interface {
	GetAllUsers(c *gin.Context)
}

type UserHandler struct {
	userService service.IUserService
}

// NewUserHandler membuat instance baru dari UserHandler
func NewUserHandler(userService service.IUserService) IUserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// GetAllUsers menghandle permintaan untuk mendapatkan semua user
func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users := h.userService.GetAllUsers()
	c.JSON(http.StatusOK, users)
}
