package models

type MessageContent struct {
	MessageID int64  `gorm:"PRIMARY_KEY;REFERENCES apate.user(id)"`
	Body      string `gorm:"type:text;not null"`
	Nonce     string `gorm:"type:text;not null"`
	Type      int32  `gorm:"type:text;not null"`

	Message *Message `gorm:"foreignkey:MessageID;association_foreignkey:ID"`
}
