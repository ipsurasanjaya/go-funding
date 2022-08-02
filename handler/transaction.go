package handler

import (
	"crowdfunding/helper"
	"crowdfunding/transaction"
	"crowdfunding/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service transaction.Service
}

func NewTransactionHandler(service transaction.Service) *handler {
	return &handler{service: service}
}

func (h *handler) GetTransactionByCampaignID(c *gin.Context) {
	var input transaction.GetCampaignTransactionInput
	err := c.ShouldBindUri(&input)

	if err != nil {
		errorMessage := gin.H{"errors": err}
		response := helper.ApiResponse("Parameter is not valid", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	curretUser := c.MustGet("currentUser").(user.User)
	input.User = curretUser

	transactions, err := h.service.GetTransactionByCampaignID(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("Failed to get transactions", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatTransaction := transaction.TransactionsFormatter(transactions)
	response := helper.ApiResponse("Retrieve transactions success", http.StatusOK, "error", formatTransaction)
	c.JSON(http.StatusOK, response)
}

func (h *handler) GetTransactionByUserID(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.Id

	transactions, err := h.service.GetTransactionByUserID(userID)
	if err != nil {
		response := helper.ApiResponse("Failed to get user's transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	transactionsFormatter := transaction.UserTransactionsFormatter(transactions)
	response := helper.ApiResponse("Success get user's transactions", http.StatusOK, "success", transactionsFormatter)
	c.JSON(http.StatusOK, response)
}
