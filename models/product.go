package models

import "time"

type Product struct {
	ID     uint `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID uint `json:"userID"`

	CategoryCode string    `json:"categoryCode"`
	Name         string    `gorm:"size:255;not null;" json:"name"`
	Description  string    `gorm:"size:255;not null;" json:"desc"`
	Price        int       `json:"price"`
	ImageSrc     string    `gorm:"size:255;" json:"imageSrc"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdateAt     time.Time `gorm:"autoUpdateTime" json:"updateAt"`
}
