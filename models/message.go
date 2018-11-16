package models

import (
	"time"
)

type Message struct {
	ID          int64          `gorm:"AUTO_INCREMENT;PRIMARY_KEY"`
	UserID      int64          `gorm:"not null"`
	RecipientID int64          `gorm:"not null"`
	CreatedAt   time.Time      `gorm:"type:timestamp"`
	OpenedAt    time.Time      `gorm:"type:timestamp;default:null"`
	DeletedAt   time.Time      `gorm:"type:timestamp;default:null"`
	Sender      User           `gorm:"foreignkey:UserID;association_foreignkey:UserID"`
	Receiver    User           `gorm:"foreignkey:UserID;association_foreignkey:RecipientID"`
	Body        MessageContent `gorm:"foreignkey:MessageID;association_foreignkey:ID"`
}
