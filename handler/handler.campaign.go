package handler

import (
	"bwastartup/campaign"
	"bwastartup/helper"
	"bwastartup/user"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type ICampaignHandler interface {
	GetCampaigns(context *gin.Context)
	GetCampaign(context *gin.Context)
	CreateCampaign(context *gin.Context)
	UpdateCampaign(context *gin.Context)
	UploadImage(context *gin.Context)

}

type campaignHandler struct {
	campaignService campaign.ICampaignService
}

func NewCampaignHandler(campaign campaign.ICampaignService) ICampaignHandler {
	return &campaignHandler{campaign}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {

	userID, _ := strconv.Atoi(c.Query("user_id"))

	getCampaign, err := h.campaignService.GetCampaigns(userID)

	if err != nil {
		res := helper.ApiResponse(false, "error", http.StatusBadRequest, helper.EmptyObj{}, err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	resp := helper.ApiResponse(true, "list of campaign", http.StatusOK, helper.MappingResponseCampaigns(getCampaign), "")
	c.JSON(http.StatusOK, resp)

}

func (h *campaignHandler) GetCampaign(c *gin.Context) {

	var param campaign.DtoCampaignDetailById
	err := c.ShouldBindUri(&param)

	if err != nil {
		res := helper.ApiResponse(false, "error in binding uri", http.StatusBadRequest, helper.EmptyObj{}, err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
     _c, errs := h.campaignService.GetCampaignById(param)
	if errs != nil {
		res := helper.ApiResponse(false, "failed to get detail campaign", http.StatusBadRequest, helper.EmptyObj{}, errs.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	resp := helper.ApiResponse(true, "success", http.StatusOK, helper.MappingResponseDetailCampaign(_c), "")
	c.JSON(http.StatusOK, resp)

}

func (h *campaignHandler) CreateCampaign(c *gin.Context) {

	var input campaign.DtoCreateCampaign
	err := c.ShouldBindJSON(&input)

	if err != nil {
		res := helper.ApiResponse(false, "error in binding input", http.StatusBadRequest, helper.EmptyObj{}, err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	currentUser := c.MustGet("CurrentUser")
	input.User = currentUser.(user.User)
	data, err := h.campaignService.CreateCampaign(input)
	if err != nil {
		res := helper.ApiResponse(false, "error in create campaign", http.StatusBadRequest, helper.EmptyObj{}, err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	resp := helper.ApiResponse(true, "success", http.StatusOK, helper.MappingResponseCampaign(data), "")
	c.JSON(http.StatusOK, resp)

}

func (h *campaignHandler) UpdateCampaign(c *gin.Context) {

	var idCampaign campaign.DtoCampaignDetailById
	err := c.ShouldBindUri(&idCampaign)
	if err != nil {
		res := helper.ApiResponse(false, "error in binding uri", http.StatusBadRequest, helper.EmptyObj{}, err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	var input campaign.DtoUpdateCampaign
	err = c.ShouldBindJSON(&input)

	if err != nil {
		res := helper.ApiResponse(false, "error in binding input", http.StatusBadRequest, helper.EmptyObj{}, err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	currentUser := c.MustGet("CurrentUser")
	input.User = currentUser.(user.User)
	updatedCampaign, errs := h.campaignService.UpdateCampaign(idCampaign,input)
	if errs != nil {
		res := helper.ApiResponse(false, "error in updated campaign", http.StatusBadRequest, helper.EmptyObj{}, errs.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	resp := helper.ApiResponse(true, "success", http.StatusOK, helper.MappingResponseCampaign(updatedCampaign), "")
	c.JSON(http.StatusOK, resp)


}

func (h *campaignHandler) UploadImage(c *gin.Context) {

	var input campaign.DtoCreateCampaignImage
	//bind input form
	err := c.ShouldBind(&input)
	if err != nil {
		res := helper.ApiResponse(false, "error in binding input", http.StatusBadRequest, nil,err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}



	file, err := c.FormFile("file")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		res := helper.ApiResponse(false, "error in form file", http.StatusBadRequest, data,"file maximum 1 mb, your file is more than 1 mb")
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	date := time.Now().Unix()
	currentUser := c.MustGet("CurrentUser")
	input.User = currentUser.(user.User)
	userId := currentUser.(user.User).ID
	path := fmt.Sprintf("images/%d-%d-%s", userId, date, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		res := helper.ApiResponse(false, "error in save upload file", http.StatusBadRequest, data, err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}



	_, err = h.campaignService.SaveCampaignImage(input, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		res := helper.ApiResponse(false, "error in save campaign image", http.StatusBadRequest, data, err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	data := gin.H{"is_uploaded": true}
	res := helper.ApiResponse(true, "campaign image succesfuly uploaded", http.StatusOK, data, "")
	c.JSON(http.StatusOK, res)

}