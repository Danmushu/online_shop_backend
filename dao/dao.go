package dao

import (
	"gorm.io/gorm"
	"project/clients/postgres"
)

var db *gorm.DB

// todo
func init() {
	db = postgres.DB
}
