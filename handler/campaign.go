package handler

import (
	"crowdfunding/campaign"
	"crowdfunding/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service: service}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.service.GetCampaigns(userID)
	if err != nil {
		apiResponse := helper.ApiResponse("Failed get campaigns", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, apiResponse)
		return
	}

	apiResponse := helper.ApiResponse("Get list of campaigns success", http.StatusOK, "success", campaign.CampaignsFormatter(campaigns))
	c.JSON(http.StatusOK, apiResponse)
}

func (h *campaignHandler) GetCampaignByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	campaignDetails, err := h.service.GetCampaignByID(id)
	if err != nil {
		apiResponse := helper.ApiResponse("Failed get campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, apiResponse)
		return
	}

	apiResponse := helper.ApiResponse("Get list of campaigns success", http.StatusOK, "success", campaign.CampaignDetailsFormatter(campaignDetails))
	c.JSON(http.StatusOK, apiResponse)
}
