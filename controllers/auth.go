package controllers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"project/clients/postgres"
	"project/middleware"
	"project/models"
	"time"
)

type UserCreate struct {
	Name  string `json:"username" binding:"required"`
	Pwd   string `json:"pwd" binding:"required"`
	Email string `json:"email" binding:"required"`
}

// RegisterUser todo
func RegisterUser(c *gin.Context) {
	var data UserCreate
	// todo 日志记录
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pwd, _ := bcrypt.GenerateFromPassword([]byte(data.Pwd), bcrypt.DefaultCost)
	user := models.User{}
	user.Name = data.Name
	user.Pwd = string(pwd)
	user.Email = data.Email
	user.UpdateAt = time.Now()

	_, err = CreateUser(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    user,
		"message": "Registration success",
	})
}

type LoginData struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginUser todo
func LoginUser(c *gin.Context) {
	var data LoginData
	//todo 日志
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := &models.User{}
	user.Name = data.Username
	user.Pwd = data.Password

	// todo 登录检查
	var token string
	user, token, err = LoginCheck(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "username or password is incorrect!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token, "data": user})
}

func LoginCheck(user *models.User) (*models.User, string, error) {

	pwd := user.Pwd

	// 存在 判断
	if postgres.DB.Model(models.User{}).Where("name = ?", user.Name).Limit(1).Find(user).RowsAffected == 0 ||
		VerifyPassword(pwd, user.Pwd) != nil {
		return nil, "", errors.New("username or password is incorrect")
	}

	token, _ := middleware.GenerateToken(user.ID)

	SaveToken(user.ID, token)
	return user, token, nil
}

func SaveToken(id int, token string) {
	postgres.DB.Model(models.User{}).Where("id = ?", id).Updates(models.User{Token: token, UpdateAt: time.Now()})
}

func VerifyPassword(pwd, hashPwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPwd), []byte(pwd))
}

func LogoutUser(c *gin.Context) {
	var user models.User
	token := middleware.ExtractToken(c)
	if err := postgres.DB.First(&user, "token = ?", token).Error; err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"status":  "Failed",
			"message": "Failed",
			"data":    "Wrong Request",
		})
		return
	}
	fmt.Println(token)
	postgres.DB.Model(models.User{}).Where("token = ?", token).Updates(models.User{Token: "-", UpdateAt: time.Now()})
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "success",
		"data":    "User Logout",
	})
}

func CurrentUser(c *gin.Context) {
	userID, err := middleware.ExtractTokenID(c)

	//todo 日志记录

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	fmt.Println(userID)
	user, err := GetUserByID(userID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data": gin.H{
			"name":  user.Name,
			"Email": user.Email,
		},
	})
}
