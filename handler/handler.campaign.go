package handler

import (
	"bwastartup/campaign"
	"bwastartup/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ICampaignHandler interface {
	GetCampaigns(context *gin.Context)
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
		res := helper.ApiResponse(false, "error", http.StatusBadRequest, getCampaign, err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	resp := helper.ApiResponse(true, "list of campaign", http.StatusOK, campaign.MappingResponseCampaign(getCampaign), "")
	c.JSON(http.StatusOK, resp)

}
