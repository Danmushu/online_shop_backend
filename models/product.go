package models

import "time"

type Product struct {
	Id           uint   `gorm:"primaryKey;" json:"id"`
	UserID       uint   `json:"userID"`
	CategoryCode string `json:"category_code"`
	Name         string `gorm:"size:255;not null;" json:"name"`
	Description  string `gorm:"size:255;not null;" json:"desc"`
	Price        int    `json:"price"`
	//Stock        int       `json:"stock"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdateAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	//User         User      `gorm:"foreignKey:Id" json:"user"`       //use User.Id
	//Category     Category  `gorm:"foreignKey:Code" json:"category"` //Use Category.Code
}
