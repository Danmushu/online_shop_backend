package controllers

import (
	"fmt"
	"net/http"
	"project/clients/postgres"
	"project/middleware"
	"project/models"

	"github.com/gin-gonic/gin"
)

type ProductInput struct {
	CategoryCode string `json:"categorycode"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Price        int    `json:"price"`
	Stock        int    `json:"stock"`
}

func CreateProduct(c *gin.Context) {
	var input ProductInput
	// todo 日志记录
	userID, err := middleware.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := models.Product{}

	product.Name = input.Name
	product.Description = input.Description
	product.Price = input.Price
	//product.Stock = input.Stock
	product.CategoryCode = input.CategoryCode
	product.UserID = userID

	erro := postgres.DB.Create(&product).Error

	if erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, Response{
		Status:  "Success",
		Message: "Success",
		Data:    product,
	})
}

func GetAllProduct(c *gin.Context) {
	var product []models.Product
	var category models.Category
	// todo 日志记录

	categoryQuery := c.Query("category")
	if categoryQuery != "" {
		fmt.Println(categoryQuery)
		postgres.DB.Where("name = ?", categoryQuery).Or("code = ?", categoryQuery).First(&category)
		postgres.DB.Where("category_code = ?", category.Code).Find(&product)
		fmt.Println(category.Code)
	} else {
		postgres.DB.Find(&product)
	}

	c.JSON(http.StatusOK, Response{
		Status:  "success",
		Message: "success",
		Data:    product,
	})
}

func GetProductById(c *gin.Context) {
	var product models.Product
	// todo 日志记录

	id := c.Param("id")
	if err := postgres.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, Response{
			Status:  "Not Found",
			Message: "product id " + id + " not found",
			Data:    NullResponse{},
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		Status:  "success",
		Message: "success",
		Data:    product,
	})
}

func UpdateProduct(c *gin.Context) {
	var product models.Product
	var input ProductInput
	var updateinput models.Product
	// todo 日志记录

	id := c.Param("id")
	user_id, err := middleware.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := postgres.DB.First(&product, id).Error; err != nil {
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
	updateinput.Description = input.Description
	updateinput.Price = input.Price
	//updateinput.Stock = input.Stock
	updateinput.UserID = user_id
	updateinput.CategoryCode = input.CategoryCode

	postgres.DB.Model(&product).Updates(updateinput)

	c.JSON(http.StatusOK, Response{
		Status:  "Success",
		Message: "Success",
		Data:    product,
	})
}

func DeleteProduct(c *gin.Context) {
	var product models.Product
	// todo 日志记录

	id := c.Param("id")
	if err := postgres.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, Response{
			Status:  "Not Found",
			Message: "user id " + id + " not found",
			Data:    NullResponse{},
		})
		return
	}
	postgres.DB.Delete(&product)

	c.JSON(http.StatusOK, Response{
		Status:  "Success",
		Message: "Success",
		Data:    NullResponse{},
	})
}
