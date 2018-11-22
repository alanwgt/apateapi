package models

import (
	"time"
)

type Message struct {
	ID          int64      `gorm:"AUTO_INCREMENT;PRIMARY_KEY"`
	UserID      int64      `gorm:"REFERENCES apate.user(id)"`
	RecipientID int64      `gorm:"REFERENCES apate.user(id)"`
	CreatedAt   *time.Time `gorm:"type:timestamp"`
	OpenedAt    *time.Time `gorm:"type:timestamp;default:null"`
	DeletedAt   *time.Time `gorm:"type:timestamp;default:null"`

	Sender   *User           `gorm:"foreignkey:UserID;association_foreignkey:ID"`
	Receiver *User           `gorm:"foreignkey:RecipientID;association_foreignkey:ID"`
	Body     *MessageContent `gorm:"foreignkey:ID;association_foreignkey:MessageID"`
}
