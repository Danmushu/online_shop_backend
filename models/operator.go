package models

import "project/utils/perms"

type Operator struct {
	ID   int        `gorm:"primaryKey;type:int;" json:"id"`
	Key  string     `gorm:"type:varchar;" json:"key"`
	Usr  string     `gorm:"type:varchar;unique;" json:"usr"`
	Pwd  string     `gorm:"type:varchar;" json:"-"`
	Perm perms.Perm `gorm:"type:int;" json:"perm"`
}
