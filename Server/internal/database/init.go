package database

import (
	"github.com/mekanican/chat-backend/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB = nil

func InitializeDatabase() {
	dsn := "host=localhost user=postgres password=postgres dbname=db port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to load sqlite database")
	}
	db.AutoMigrate(&model.User{}, &model.Chat{}, &model.Room{}) // Loading the model
	DB = db
}
