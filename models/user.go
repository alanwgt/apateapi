package models

import (
	"time"
)

type User struct {
	ID         int64     `gorm:"AUTO_INCREMENT;PRIMARY_KEY"`
	Username   string    `gorm:"type:varchar(20);unique;not null"`
	PubKey     string    `gorm:"type:text;not null"`
	RecoverKey string    `gorm:"type:text"`
	FcmToken   string    `gorm:"type:varchar(255)"`
	CreatedAt  time.Time `gorm:"type:timestamp"`
	UpdatedAt  time.Time `gorm:"type:timestamp"`
	DeletedAt  time.Time `gorm:"type:timestamp;default:null"`
}
