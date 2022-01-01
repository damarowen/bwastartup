package handler

import (
	"bwastartup/campaign"
	"bwastartup/helper"
	"bwastartup/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ICampaignHandler interface {
	GetCampaigns(context *gin.Context)
	GetCampaign(context *gin.Context)
	CreateCampaign(context *gin.Context)
	UpdateCampaign(context *gin.Context)
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
		res := helper.ApiResponse(false, "error in updated campaign", http.StatusBadRequest, helper.EmptyObj{}, "not authorize, different user")
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	resp := helper.ApiResponse(true, "success", http.StatusOK, helper.MappingResponseCampaign(updatedCampaign), "")
	c.JSON(http.StatusOK, resp)


}