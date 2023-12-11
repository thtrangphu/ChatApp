package model

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	Chats []Chat
}
