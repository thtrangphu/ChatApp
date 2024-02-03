package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mekanican/chat-backend/internal/config"
)

// Generate AuthN & AuthZ token
func GenerateRoomJWT(userID, channel string, expireHour time.Duration) (string, error) {
	var claims jwt.MapClaims
	if channel != "" {
		claims = jwt.MapClaims{
			"sub":     userID,
			"channel": channel,
			"exp":     time.Now().Add(expireHour * time.Hour).Unix(),
			"iat":     time.Now().Unix(),
		}
	} else {
		claims = jwt.MapClaims{
			"sub": userID,
			"exp": time.Now().Add(expireHour * time.Hour).Unix(),
			"iat": time.Now().Unix(),
		}
	}

	// uuid, err := uuid.Parse(config.GetString("TOKEN"))
	// if err != nil {
	// 	return "", err
	// }

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// tokenString, err := token.SignedString(uuid[:])
	tokenString, err := token.SignedString([]byte(config.GetString("TOKEN")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
