package controllers

import (
	"net/http"
	"project/clients/postgres"
	"project/models"

	"github.com/gin-gonic/gin"
)

type TransactionInput struct {
	CartID            uint   `json:"cart_id"`
	PaymentMethodCode string `json:"code"`
}

func CreateTransaction(c *gin.Context) {
	var input TransactionInput
	//todo 日志记录
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	transaction := models.Transaction{}
	transaction.CartID = input.CartID
	transaction.PaymentMethodCode = input.PaymentMethodCode

	err := postgres.DB.Create(&transaction).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, Response{
		Status:  "success",
		Message: "success",
		Data:    transaction,
	})
}

func GetAllTransaction(c *gin.Context) {
	var transaction []models.Transaction
	//todo 日志记录

	postgres.DB.Find(&transaction)
	c.JSON(http.StatusOK, Response{
		Status:  "success",
		Message: "success",
		Data:    transaction,
	})
}

func GetTransactionById(c *gin.Context) {
	var transaction models.Transaction
	//todo 日志记录

	id := c.Param("id")
	if err := postgres.DB.First(&transaction, id).Error; err != nil {
		c.JSON(http.StatusNotFound, Response{
			Status:  "Not Found",
			Message: "transaction id " + id + " not found",
			Data:    NullResponse{},
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		Status:  "success",
		Message: "success",
		Data:    transaction,
	})
}

func UpdateTransaction(c *gin.Context) {
	var transaction models.Transaction
	var input TransactionInput
	var updateinput models.Transaction
	//todo 日志记录

	id := c.Param("id")

	if err := postgres.DB.First(&transaction, id).Error; err != nil {
		c.JSON(http.StatusNotFound, Response{
			Status:  "Not Found",
			Message: "user id " + id + " not found",
			Data:    NullResponse{},
		})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateinput.CartID = input.CartID
	updateinput.PaymentMethodCode = input.PaymentMethodCode

	postgres.DB.Model(&transaction).Updates(updateinput)

	c.JSON(http.StatusOK, Response{
		Status:  "Success",
		Message: "Success",
		Data:    transaction,
	})
}

func DeleteTransaction(c *gin.Context) {
	var transaction models.Transaction
	//todo 日志记录

	id := c.Param("id")
	if err := postgres.DB.First(&transaction, id).Error; err != nil {
		c.JSON(http.StatusNotFound, Response{
			Status:  "Not Found",
			Message: "user id " + id + " not found",
			Data:    NullResponse{},
		})
		return
	}
	postgres.DB.Delete(&transaction)

	c.JSON(http.StatusOK, Response{
		Status:  "Success",
		Message: "Success",
		Data:    NullResponse{},
	})
}
