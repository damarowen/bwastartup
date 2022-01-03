package handler

import (
	"bwastartup/helper"
	"bwastartup/transaction"
	"bwastartup/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ITransactionHandler interface {
	GetTransactionsByCampaignId(context *gin.Context)
	GetUserTransactionsByUserId(context *gin.Context)
	CreateTransaction(context *gin.Context)

}

type TransactionHandler struct {
	TransactionService transaction.ITransactionService
}

func NewTransactionHandler(Transaction transaction.ITransactionService) ITransactionHandler {
	return &TransactionHandler{Transaction}
}


func (h *TransactionHandler) GetTransactionsByCampaignId(c *gin.Context) {

	var input transaction.DtoTransactionByCampaignId
	errUri := c.ShouldBindUri(&input)
	if errUri != nil {
		res := helper.ApiResponse(false, "error in binding uri", http.StatusBadRequest, helper.EmptyObj{}, errUri.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	currentUser := c.MustGet("CurrentUser")
	input.User = currentUser.(user.User)
	t, err := h.TransactionService.GetTransactionByCampaignId(input)
	if err != nil {
		res := helper.ApiResponse(false, "failed to get Transaction", http.StatusBadRequest, helper.EmptyObj{}, err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	resp := helper.ApiResponse(true, "success", http.StatusOK, helper.MappingResponseCampaignTransactions(t), "")
	c.JSON(http.StatusOK, resp)

}


func (h *TransactionHandler) GetUserTransactionsByUserId(c *gin.Context) {
	currentUser := c.MustGet("CurrentUser")
	userID := currentUser.(user.User).ID

	t, err := h.TransactionService.GetTransactionsByUserID(userID)
	if err != nil {
		res := helper.ApiResponse(false, "failed to get user Transaction", http.StatusBadRequest, helper.EmptyObj{}, err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	resp := helper.ApiResponse(true, "list of User Transaction", http.StatusOK, helper.MappingResponseUserTransactions(t), "")
	c.JSON(http.StatusOK, resp)
}

func (h *TransactionHandler) CreateTransaction(c *gin.Context) {
	var input transaction.DtoCreateTransaction

	err := c.ShouldBindJSON(&input)

	if err != nil {
		res := helper.ApiResponse(false, "error in binding input", http.StatusBadRequest, helper.EmptyObj{}, err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	currentUser := c.MustGet("CurrentUser")
	input.User = currentUser.(user.User)

	newTransaction, err := h.TransactionService.CreateTransaction(input)

	if err != nil {
		res := helper.ApiResponse(false, "error in create transaction", http.StatusBadRequest, helper.EmptyObj{}, err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}


	resp := helper.ApiResponse(true, "list of User Transaction", http.StatusOK, helper.MappingFormatTransaction(newTransaction), "")
	c.JSON(http.StatusOK, resp)
}

//func (h *transactionHandler) GetNotification(c *gin.Context) {
//	var input transaction.TransactionNotificationInput
//
//	err := c.ShouldBindJSON(&input)
//	if err != nil {
//		response := helper.APIResponse("Failed to process notification", http.StatusBadRequest, "error", nil)
//		c.JSON(http.StatusBadRequest, response)
//
//		return
//	}
//
//	err = h.service.ProcessPayment(input)
//	if err != nil {
//		response := helper.APIResponse("Failed to process notification", http.StatusBadRequest, "error", nil)
//		c.JSON(http.StatusBadRequest, response)
//
//		return
//	}
//
//	c.JSON(http.StatusOK, input)
//}