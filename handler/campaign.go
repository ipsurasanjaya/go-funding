package handler

import (
	"crowdfunding/campaign"
	"crowdfunding/helper"
	"crowdfunding/user"
	"fmt"
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

func (h *campaignHandler) CreateCampaign(c *gin.Context) {
	var input campaign.CampaignInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errMessage := gin.H{"errors": errors}
		response := helper.ApiResponse("Create campaign failed", http.StatusUnprocessableEntity, "error", errMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	user := c.MustGet("currentUser").(user.User)
	input.User = user

	newCampaign, err := h.service.CreateCampaign(input)
	if err != nil {
		errMessage := gin.H{"errors": err}
		response := helper.ApiResponse("Create campaign failed", http.StatusUnprocessableEntity, "error", errMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	campaignFormatter := campaign.CampaignFormatter(newCampaign)
	campaignResponse := helper.ApiResponse("Campaign successfuly created", http.StatusCreated, "success", campaignFormatter)
	c.JSON(http.StatusCreated, campaignResponse)
}

func (h *campaignHandler) UpdateCampaign(c *gin.Context) {
	var (
		campaignInput campaign.CampaignInput
	)

	campaignID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errMessage := gin.H{"errors": err}
		response := helper.ApiResponse("Update campaign failed", http.StatusUnprocessableEntity, "error", errMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err = c.ShouldBindJSON(&campaignInput)
	if err != nil {
		if err != nil {
			errors := helper.FormatValidationError(err)
			errMessage := gin.H{"errors": errors}
			response := helper.ApiResponse("Update campaign failed", http.StatusUnprocessableEntity, "error", errMessage)
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}
	}

	currentUser := c.MustGet("currentUser").(user.User)
	campaignInput.User = currentUser

	updatedCampaign, err := h.service.UpdateCampaign(campaignInput, campaignID)
	if err != nil {
		errMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("Update campaign failed", http.StatusUnprocessableEntity, "error", errMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	campaignFormatter := campaign.CampaignFormatter(updatedCampaign)
	campaignResponse := helper.ApiResponse("Campaign successfuly created", http.StatusCreated, "success", campaignFormatter)
	c.JSON(http.StatusCreated, campaignResponse)
}

func (h *campaignHandler) CreateCampaignImage(c *gin.Context) {
	var input campaign.CampaignImageInput

	err := c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errMessage := gin.H{"errors": errors}
		response := helper.ApiResponse("Failed to upload campaign image binding", http.StatusUnprocessableEntity, "error", errMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser
	userId := currentUser.Id

	file, err := c.FormFile("file")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.ApiResponse("Failed to upload campaign image formfile", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	path := fmt.Sprintf("images/%d-%s", userId, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.ApiResponse("Failed to upload campaign image save uploaded file", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.service.CreateCampaignImage(input, path)
	if err != nil {
		data := gin.H{"is_uploaded": false, "errors": err.Error()}
		response := helper.ApiResponse("Failed to upload campaign image service", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.ApiResponse("Campaign image successfuly uploaded", http.StatusOK, "succes", data)
	c.JSON(http.StatusOK, response)
}
