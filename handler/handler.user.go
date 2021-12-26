package handler

import (
	"bwastartup/helper"
	"bwastartup/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IUserHandler interface {
	RegisterUser(context *gin.Context)
	LoginUser(context *gin.Context)
	IsDuplicateEmail(context *gin.Context)
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
		res := helper.ApiResponse(false, "error", http.StatusBadRequest, nil, errDTO.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	duplicate, _ := h.userService.IsDuplicateEmail(input.Email)
	if duplicate {
		res := helper.ApiResponse(false, "email has been registered", http.StatusBadRequest, "", "")
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		res := helper.ApiResponse(false, "error", http.StatusBadRequest, nil, err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	mapping := helper.MappingResponseUser(newUser, "test")
	resp := helper.ApiResponse(true,"success",http.StatusOK,mapping,"")
	c.JSON(http.StatusCreated, resp)

}

func (h *userHandler) LoginUser(c *gin.Context) {

	var input user.DtoLoginUserInput
	errDTO := c.ShouldBind(&input)
	if errDTO != nil {
		res := helper.ApiResponse(false, "error", http.StatusBadRequest, nil, errDTO.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	userLogin, err := h.userService.LoginUser(input)

	if err != nil {
		res := helper.ApiResponse(false, "error", http.StatusBadRequest, nil, err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	resp := helper.ApiResponse(true,"success",http.StatusOK,userLogin,"")
	c.JSON(http.StatusCreated, resp)

}

func (h *userHandler) IsDuplicateEmail (c *gin.Context) {
	var input user.DtoEmailChecker
	errDTO := c.ShouldBind(&input)

	if errDTO != nil {
		res := helper.ApiResponse(false, "error", http.StatusBadRequest, nil, errDTO.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	duplicate, _ := h.userService.IsDuplicateEmail(input.Email)

	mapping := gin.H{
		"is_duplicate": duplicate,
	}

	metaMessage := "email has been registered"
	if !duplicate {
		metaMessage = "email not been registered"
	}

	resp := helper.ApiResponse(true,metaMessage,http.StatusOK,mapping,"")
	c.JSON(http.StatusOK, resp)

}