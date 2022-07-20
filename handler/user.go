package handler

import (
	"crowdfunding/helper"
	"crowdfunding/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	service user.Service
}

func NewUserHandler(service user.Service) *userHandler {
	return &userHandler{service: service}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	input := user.RegisterUserInput{}
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errMessage := gin.H{"errors": errors}
		response := helper.ApiResponse("Register account failed", http.StatusUnprocessableEntity, "error", errMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	newUser, err := h.service.RegisterUser(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}
	formatter := user.FormatUser(newUser, "tokeninisangatrahasia")
	response := helper.ApiResponse("Your account has been created", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {
	input := user.LoginInput{}
	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errMessage := gin.H{"errors": errors}
		response := helper.ApiResponse("Login failed", http.StatusUnprocessableEntity, "error", errMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	loggedinUser, err := h.service.Login(input)
	if err != nil {
		errMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("Email or password is not valid", http.StatusUnprocessableEntity, "error", errMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(loggedinUser, "tokeninisangatrahasia")
	successResponse := helper.ApiResponse("Login successfuly", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, successResponse)
}

func (h *userHandler) CheckEmailAvailability(c *gin.Context) {
	var input user.CheckEmailInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("Email checking failed", http.StatusUnprocessableEntity, "error", errMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isEmailAvailable, err := h.service.IsEmailAvailable(input)
	if err != nil {
		errMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("Email checking failed", http.StatusUnprocessableEntity, "error", errMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	metaMessage := "Email has been registered"
	if isEmailAvailable {
		metaMessage = "Email is available"
	}
	data := gin.H{"is_available": isEmailAvailable}
	response := helper.ApiResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}
