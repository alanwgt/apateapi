package models

import (
	"time"
)

type Blocked struct {
	ID        int64      `gorm:"AUTO_INCREMENT;PRIMARY_KEY"`
	UserID    int64      `gorm:"REFERENCES apate.user(id)"`
	BlockedID int64      `gorm:"REFERENCES apate.user(id)"`
	CreatedAt *time.Time `gorm:"type:timestamp"`
	DeletedAt *time.Time `gorm:"type:timestamp;default:null"`

	Blocked      *User `gorm:"foreignkey:BlockedID;association_foreignkey:ID"`
	RequestBlock *User `gorm:"foreignkey:UserID;association_foreignkey:UD"`
}
