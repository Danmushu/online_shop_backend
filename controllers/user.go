package controllers

import (
	"project/clients/postgres"
	"project/models"
)

type UserResponse struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// CreateUser todo
func CreateUser(user *models.User) (*models.User, error) {
	var err error
	err = postgres.DB.Create(user).Error
	if err != nil {
		return &models.User{}, err
	}
	return user, nil
}
