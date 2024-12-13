package models

import "time"

type Cart struct {
	Id        uint      `gorm:"primaryKey;" json:"id"`
	UserID    int       `json:"userID"`
	ProductID int       `json:"productID"`
	Quantity  int       `json:"quantity"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdateAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	//User      User      `gorm:"foreignKey:UserID"`    //use User.Id
	//Product   Product   `gorm:"foreignKey:ProductID"` //use Product.Id
}
