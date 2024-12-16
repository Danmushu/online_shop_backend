package controllers

import (
	"net/http"
	"project/clients/postgres"
	"project/models"

	"github.com/gin-gonic/gin"
)

type PaymentInput struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

func CreatePayment(c *gin.Context) {
	var input PaymentInput
	//todo 日志记录
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	payment := models.Category{}
	payment.Name = input.Name
	payment.Code = input.Code

	err := postgres.DB.Create(&payment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, Response{
		Status:  "Success",
		Message: "Success",
		Data:    payment,
	})
}

func GetAllPayment(c *gin.Context) {
	var payment []models.PaymentMethod

	//todo 日志记录

	postgres.DB.Find(&payment)
	c.JSON(http.StatusOK, Response{
		Status:  "success",
		Message: "success",
		Data:    payment,
	})
}

func GetPaymentById(c *gin.Context) {
	var payment models.PaymentMethod

	//todo 日志记录

	id := c.Param("id")
	if err := postgres.DB.First(&payment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, Response{
			Status:  "Not Found",
			Message: "payment id " + id + " not found",
			Data:    NullResponse{},
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		Status:  "success",
		Message: "success",
		Data:    payment,
	})
}

func UpdatePayment(c *gin.Context) {
	var payment models.PaymentMethod
	var input PaymentInput
	var updateinput models.PaymentMethod

	//todo 日志记录

	id := c.Param("id")

	if err := postgres.DB.First(&payment, id).Error; err != nil {
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

	updateinput.Name = input.Name
	updateinput.Code = input.Code

	postgres.DB.Model(&payment).Updates(updateinput)

	c.JSON(http.StatusOK, Response{
		Status:  "Success",
		Message: "Success",
		Data:    payment,
	})
}

func DeletePayment(c *gin.Context) {
	var payment models.PaymentMethod

	//todo 日志记录

	id := c.Param("id")
	if err := postgres.DB.First(&payment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, Response{
			Status:  "Not Found",
			Message: "user id " + id + " not found",
			Data:    NullResponse{},
		})
		return
	}
	postgres.DB.Delete(&payment)

	c.JSON(http.StatusOK, Response{
		Status:  "Success",
		Message: "Success",
		Data:    NullResponse{},
	})
}
