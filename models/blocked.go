package models

import (
	"time"
)

type Blocked struct {
	ID           int64     `gorm:"AUTO_INCREMENT;PRIMARY_KEY"`
	UserID       int64     `gorm:"not null"`
	BlockedID    int64     `gorm:"not null"`
	CreatedAt    time.Time `gorm:"type:timestamp"`
	DeletedAt    time.Time `gorm:"type:timestamp;default:null"`
	Blocked      User      `gorm:"foreignkey:UserID;association_foreignkey:BlockedID"`
	RequestBlock User      `gorm:"foreignkey:UserID;association_foreignkey:UserID"`
}
