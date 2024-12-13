package controllers

import (
	"github.com/gin-gonic/gin"
)

type UserCreate struct {
	Name  string `json:"username" binding:"required"`
	Pwd   string `json:"password" binding:"required"`
	Email string `json:"email" binding:"required"`
}

// RegisterUser todo
func RegisterUser(c *gin.Context) {

}

func LoginUser(c *gin.Context) {}

func LogoutUser(c *gin.Context) {}

func CurrentUser(c *gin.Context) {}
