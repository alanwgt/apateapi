package models

import (
	"time"
)

type User struct {
	ID         int64      `gorm:"AUTO_INCREMENT;PRIMARY_KEY" json:",omitempty"`
	Username   string     `gorm:"type:varchar(20);unique;not null" json:",omitempty"`
	PubKey     string     `gorm:"type:text;not null" json:",omitempty"`
	RecoverKey string     `gorm:"type:text" json:"-"`
	FcmToken   string     `gorm:"type:varchar(255)" json:"-"`
	CreatedAt  *time.Time `gorm:"type:timestamp" json:",omitempty"`
	UpdatedAt  *time.Time `gorm:"type:timestamp" json:"-"`
	DeletedAt  *time.Time `gorm:"type:timestamp;default:null" json:"-"`

	FriendRequests   []*FriendRequest `gorm:"foreignkey:ID;association_foreignkey:RequestTo"`
	RequestedFriends []*FriendRequest `gorm:"foreignkey:ID;association_foreignkey:UserID"`

	Blocked   []*Blocked `gorm:"foreignkey:ID;association_foreignkey:BlockedID"`
	BlockedBy []*Blocked `gorm:"foreignkey:ID;association_foreignkey:UserID"`

	SentMessages     []*Message `gorm:"foreignkey:ID;association_foreignkey:UserID"`
	ReceivedMessages []*Message `gorm:"foreignkey:ID;association_foreignkey:RecipientID"`
}
