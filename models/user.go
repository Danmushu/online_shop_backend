package models

import "time"

// Account todo
//type Account struct {
//	ID          int       `gorm:"primaryKey;type:bigint;autoIncrement;" json:"id"`
//	MAC         string    `gorm:"type:varchar;" json:"mac"`
//	IsBanned    bool      `gorm:"type:bool;" json:"is_banned"`
//	ExpiredAt   time.Time `gorm:"type:time;" json:"expired_at"`
//	CreatedAt   time.Time `gorm:"autoCreateTime;" json:"created_at"`
//	LastLoginAt time.Time `gorm:"type:time;" json:"last_login_at"`
//	Secret      string    `gorm:"type:varchar;" json:"secret"`
//	LastLoginIP string    `gorm:"type:varchar;" json:"last_login_ip"`
//}

type User struct {
	ID        int       `gorm:"primaryKey;type:bigint;autoIncrement;" json:"id"`
	Name      string    `gorm:"size:255;not null;" json:"name"`
	Pwd       string    `gorm:"size:255;not null;" json:"pwd"`
	Email     string    `gorm:"size:255;not null;unique" json:"email"`
	Token     string    `gorm:"" json:"token"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdateAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updatedAt"`
}
