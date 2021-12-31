package handler

import (
	"bwastartup/auth"
	"bwastartup/helper"
	"bwastartup/user"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type IUserHandler interface {
	RegisterUser(context *gin.Context)
	LoginUser(context *gin.Context)
	IsDuplicateEmail(context *gin.Context)
	UploadAvatar(context *gin.Context)
}

type userHandler struct {
	userService user.IUserService
	authService auth.IJwtService
}

func NewUserHandler(userService user.IUserService, authService auth.IJwtService) IUserHandler {
	return &userHandler{userService, authService}
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

	token := h.authService.GenerateToken(newUser.ID, newUser.Name, newUser.Email)
	mapping := helper.MappingResponseUser(newUser, token)
	resp := helper.ApiResponse(true, "success", http.StatusOK, mapping, "")
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
	token := h.authService.GenerateToken(userLogin.ID, userLogin.Name, userLogin.Email)
	mapping := helper.MappingResponseUser(userLogin, token)
	resp := helper.ApiResponse(true, "success", http.StatusOK, mapping, "")
	c.JSON(http.StatusCreated, resp)

}

func (h *userHandler) IsDuplicateEmail(c *gin.Context) {
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

	resp := helper.ApiResponse(true, metaMessage, http.StatusOK, mapping, "")
	c.JSON(http.StatusOK, resp)

}

func (h *userHandler) UploadAvatar(c *gin.Context) {

	file, err := c.FormFile("avatar")
	if err != nil {
		res := helper.ApiResponse(false, "error in form file", http.StatusBadRequest, nil,"file maximum 1 mb, your file is more than 1 mb")
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	date := time.Now().Unix()
	userId := c.MustGet("CurrentUser").(jwt.MapClaims)["user_id"]
	path := fmt.Sprintf("images/%d-%d-%s", int(userId.(float64)), date, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		res := helper.ApiResponse(false, "error in save upload file", http.StatusBadRequest, data, err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	_, err = h.userService.SaveAvatarUser(int(userId.(float64)), path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		res := helper.ApiResponse(false, "error in save avatar user", http.StatusBadRequest, data, err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	data := gin.H{"is_uploaded": true}
	res := helper.ApiResponse(true, "avatar succesfuly uploaded", http.StatusOK, data, "")
	c.JSON(http.StatusOK, res)

}
