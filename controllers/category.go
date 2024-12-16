package controllers

import (
	"net/http"
	"project/clients/postgres"
	"project/models"

	"github.com/gin-gonic/gin"
)

type CategoryInput struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

func CreateCategory(c *gin.Context) {
	var input CategoryInput
	//todo 日志记录
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	category := models.Category{}
	category.Name = input.Name
	category.Code = input.Code

	err := postgres.DB.Create(&category).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, Response{
		Status:  "success",
		Message: "success",
		Data:    category,
	})
}

func GetAllCategory(c *gin.Context) {
	var category []models.Category
	//todo 日志记录
	postgres.DB.Find(&category)
	c.JSON(http.StatusOK, Response{
		Status:  "success",
		Message: "success",
		Data:    category,
	})
}

func GetCategoryById(c *gin.Context) {
	var category models.Category
	//todo 日志记录
	id := c.Param("id")
	if err := postgres.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, Response{
			Status:  "Not Found",
			Message: "category id " + id + " not found",
			Data:    NullResponse{},
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		Status:  "success",
		Message: "success",
		Data:    category,
	})
}

func UpdateCategory(c *gin.Context) {
	var category models.Category
	var input CategoryInput
	var updateinput models.Category
	//todo 日志记录
	id := c.Param("id")

	if err := postgres.DB.First(&category, id).Error; err != nil {
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

	postgres.DB.Model(&category).Updates(updateinput)

	c.JSON(http.StatusOK, Response{
		Status:  "Success",
		Message: "Success",
		Data:    category,
	})
}

func DeleteCategory(c *gin.Context) {
	var category models.Category
	//todo 日志记录
	id := c.Param("id")
	if err := postgres.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, Response{
			Status:  "Not Found",
			Message: "user id " + id + " not found",
			Data:    NullResponse{},
		})
		return
	}
	postgres.DB.Delete(&category)

	c.JSON(http.StatusOK, Response{
		Status:  "Success",
		Message: "Success",
		Data:    NullResponse{},
	})
}
