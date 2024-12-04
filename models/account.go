package models

import "time"

// Account todo
type Account struct {
	ID          int       `gorm:"primaryKey;type:bigint;autoIncrement;" json:"id"`
	MAC         string    `gorm:"type:varchar;" json:"mac"`
	IsBanned    bool      `gorm:"type:bool;" json:"is_banned"`
	ExpiredAt   time.Time `gorm:"type:time;" json:"expired_at"`
	CreatedAt   time.Time `gorm:"autoCreateTime;" json:"created_at"`
	LastLoginAt time.Time `gorm:"type:time;" json:"last_login_at"`
	Secret      string    `gorm:"type:varchar;" json:"secret"`
	LastLoginIP string    `gorm:"type:varchar;" json:"last_login_ip"`
}
