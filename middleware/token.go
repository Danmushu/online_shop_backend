package middleware

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"os"
	"project/clients/postgres"
	"project/global"
	"project/models"
	"strconv"
	"strings"
	"time"
)

// todo

func GenerateToken(userID int) (string, error) {

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(global.JWTDuration).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	fmt.Println(token)
	fmt.Println(123123123)
	return token.SignedString([]byte(global.JWTSecret))
}

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func ExtractTokenID(c *gin.Context) (uint, error) {
	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)
		if err != nil {
			return 0, err
		}
		return uint(uid), nil
	}
	return 0, nil
}

func GetUserByToken(token string) error {
	var user models.User
	if err := postgres.DB.First(&user, "token = ?", token).Error; err != nil {
		return errors.New("token Expired")
	}
	return nil
}

func TokenValid(c *gin.Context) error {
	tokenString := ExtractToken(c)
	erro := GetUserByToken(tokenString)
	if erro != nil {
		return erro
	}
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return err
	}
	return nil
}
