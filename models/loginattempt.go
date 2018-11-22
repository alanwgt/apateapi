package models

import (
	"time"
)

type LoginAttempt struct {
	ID        int64      `gorm:"AUTO_INCREMENT;PRIMARY_KEY"`
	UserID    int64      `gorm:"REFERENCES apate.user(id)"`
	CreatedAt *time.Time `gorm:"type:timestamp"`

	User *User `gorm:"foreignkey:UserID;association_foreignkey:ID"`
}
