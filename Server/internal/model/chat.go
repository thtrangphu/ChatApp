package model

import "gorm.io/gorm"

type ChatType uint

const (
	ChatText uint = iota
	ChatImage
	ChatFile
)

type Chat struct {
	gorm.Model
	RoomID  uint
	UserID  uint
	From    User
	Content string
	Type    ChatType
}
