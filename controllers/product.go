package controllers

import (
	"fmt"
	"github.com/spf13/cast"
	"net/http"
	"project/clients/postgres"
	"project/middleware"
	"project/models"
	"strings"

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

func ProductGet(c *gin.Context) {

	// todo 日志记录

	categoryQuery := c.Query("category")
	keyQuery := c.Query("key")
	limit := cast.ToInt(c.Query("limit"))
	offset := cast.ToInt(c.Query("offset"))
	orderBy := c.Query("orderBy")
	id := cast.ToInt(c.Param("id"))
	ids := cast.ToIntSlice(strings.Split(c.Query("ids"), ","))
	// 单个查询
	if id != 0 {
		product := models.Product{}
		postgres.DB.Where("id = ?", id).Limit(1).Find(&product)
		c.JSON(http.StatusOK, Response{
			Status:  "success",
			Message: "success",
			Data:    product,
		})
		return
	}

	// 多个ids查询
	if len(ids) != 0 {
		var products []models.Product
		postgres.DB.Where("id in ?", ids).Find(&products)
		c.JSON(http.StatusOK, Response{
			Status:  "success",
			Message: "success",
			Data:    products,
		})
		return
	}

	// 多个检索
	var products []models.Product
	if limit < 0 || limit > 30 {
		limit = 30
	}
	if offset < 0 {
		offset = 0
	}
	query := postgres.DB.
		Model(models.Product{})

	// 如果有类别
	if categoryQuery != "" {
		query.Where("category_code = ?", categoryQuery)
	}

	// 如果有关键字
	if keyQuery != "" {
		query.Where("name like ? OR description like ?", "%"+keyQuery+"%", "%"+keyQuery+"%")
	}

	// 排序方法
	if orderBy == "price" || orderBy == "created_at" {
		query.Order(fmt.Sprintf("%s desc", orderBy))
	} else {
		// 默认 排序
		query.Order("created_at desc")
	}

	query.Limit(limit).Offset(offset).
		Find(&products)

	c.JSON(http.StatusOK, Response{
		Status:  "success",
		Message: "success",
		Data:    products,
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
