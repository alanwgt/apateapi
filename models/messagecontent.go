package models

type MessageContent struct {
	MessageID int64  `gorm:"PRIMARY_KEY"`
	Body      string `gorm:"type:text;not null"`
}
