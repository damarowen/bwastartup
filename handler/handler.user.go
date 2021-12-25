package handler

import (
	"bwastartup/helper"
	"bwastartup/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IUserHandler interface {
	RegisterUser(context *gin.Context)
}

type userHandler struct {
	userService user.IUserService
}

func NewUserHandler(userService user.IUserService) IUserHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.DtoRegisterUserInput
	errDTO := c.ShouldBind(&input)
	if errDTO != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "error")
	}
	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "error")
		return
	}

	mapping := helper.MappingResponseUser(newUser, "test")
	resp := helper.ApiResponse(true,"success",http.StatusOK,mapping,"")
	c.JSON(http.StatusCreated, resp)

}
