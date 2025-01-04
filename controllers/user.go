package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"project/clients/postgres"
	"project/models"
	"time"
)

type UserResponse struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// CreateUser todo
func CreateUser(user *models.User) (*models.User, error) {
	err := postgres.DB.Create(user).Error
	return user, err
}

func GetUserById(c *gin.Context) {
	var user models.User
	// todo 增加一个日志记录的中间件
	id := c.Param("id")
	err := postgres.DB.First(&user, id).Error
	if err == nil {
		c.JSON(404, gin.H{
			"status":  "Not Found",
			"message": "user id " + id + " not found",
			"data":    "",
		})
		return
	}
	c.JSON(200, gin.H{
		"status":  "OK",
		"message": "success",
		"data":    user,
	})
}

func GetUserByID(id uint) (models.User, error) {
	var user models.User
	if err := postgres.DB.First(&user, id).Error; err != nil {
		return user, errors.New("User not found!")
	}
	return user, nil
}

func UpdateUser(c *gin.Context) {
	var user models.User
	var input UserCreate
	var updateinput models.User

	// todo 加入日志记录

	id := c.Param("id")
	err := postgres.DB.First(&user, id).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "Not Found",
			"message": "user id " + id + " not found",
			"data":    "",
		})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(input.Pwd), bcrypt.DefaultCost)

	updateinput.Name = input.Name
	updateinput.Name = input.Name
	updateinput.Pwd = string(password)
	updateinput.Email = input.Email
	updateinput.UpdateAt = time.Now()

	postgres.DB.Model(&user).Updates(updateinput)

	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    user,
	})
}

func DeleteUser(c *gin.Context) {
	var user models.User

	// todo 加入日志记录

	id := c.Param("id")
	if err := postgres.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "Not Found",
			"message": "user id " + id + " not found",
			"data":    "",
		})
		return
	}
	postgres.DB.Delete(&user)

	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    "",
	})
}
