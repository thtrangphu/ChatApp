package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mekanican/chat-backend/internal/config"
)

var secret string = config.GetString("TOKEN")

func GenerateRoomJWT(userID, channel string, expireHour time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"sub":     userID,
		"channel": channel,
		"exp":     time.Now().Add(expireHour * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
