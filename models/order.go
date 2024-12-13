package models

import "time"

type Order struct {
	ID     int    `gorm:"type:bigint;primaryKey;autoIncrement;" json:"id"`
	No     string `gorm:"type:varchar" json:"no"`
	Amount string `gorm:"type:varchar;" json:"amount"`
	Type   int    `gorm:"type:int;" json:"type"`

	Status int    `gorm:"type:int;"        json:"status"` // 0 未成交 1 成交 2 结束 (已发货)
	Usr    string `gorm:"type:varchar;"    json:"usr"`
	Email  string `gorm:"type:varchar;" json:"email"`
	//Secret   string `gorm:"type:varchar" json:"secret"`
	//Platform int    `gorm:"type:int" json:"platform"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;" json:"updated_at"`
}

const h = time.Duration(3600 * 1e9)

func (order *Order) GetDuration() time.Duration {
	return []time.Duration{6 * h, 24 * h, 168 * h, 720 * h, 2160 * h, 8760 * h}[order.Type]
}
