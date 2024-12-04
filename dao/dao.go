package dao

import (
	"gorm.io/gorm"
	"project/clients/postgres"
)

var db *gorm.DB

func init() {
	db = postgres.DB
}
