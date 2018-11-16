package models

import (
	"time"
)

type FriendRequest struct {
	ID          int64     `gorm:"AUTO_INCREMENT;PRIMARY_KEY"`
	UserID      int64     `gorm:"not null"`
	RequestTo   int64     `gorm:"not null"`
	CreatedAt   time.Time `gorm:"type:timestamp"`
	AcceptedAt  time.Time `gorm:"type:timestamp;default:null`
	DeletedAt   time.Time `gorm:"type:timestamp;default:null"`
	Requester   User      `gorm:"foreignkey:UserID;association_foreignkey:UserID"`
	RequestedTo User      `gorm:"foreignkey:UserID;association_foreignkey:RequestTo"`
}
