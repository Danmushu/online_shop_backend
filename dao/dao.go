package dao

import "gorm.io/gorm"

var db *gorm.DB

func init() {
	db = postgres.DB
}
