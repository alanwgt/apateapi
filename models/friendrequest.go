package models

import (
	"time"
)

type FriendRequest struct {
	ID         int64      `gorm:"AUTO_INCREMENT;PRIMARY_KEY"`
	UserID     int64      `gorm:"REFERENCES apate.user(id)"`
	RequestTo  int64      `gorm:"REFERENCES apate.user(id)"`
	CreatedAt  *time.Time `gorm:"type:timestamp"`
	AcceptedAt *time.Time `gorm:"type:timestamp;default:null"`
	DeletedAt  *time.Time `gorm:"type:timestamp;default:null"`

	Requester   *User `gorm:"foreignkey:UserID;association_foreignkey:ID"`
	RequestedTo *User `gorm:"foreignkey:RequestTo;association_foreignkey:ID"`
}
