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

	userID, _ := strconv.Atoi(c.Query("userId"))

	campaign, err := h.campaignService.GetCampaigns(userID)

	if err != nil {
		res := helper.ApiResponse(false, "error", http.StatusBadRequest, nil, err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	resp := helper.ApiResponse(true, "list of campaign", http.StatusOK, campaign, "")
	c.JSON(http.StatusOK, resp)

}
