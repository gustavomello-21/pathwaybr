package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gustavomello-21/pathwaybr-backend/internal/entities"
)

var jwtSecret = []byte("secret")

func GenerateToken(user entities.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"email":    user.Email,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return err
	}

	if !token.Valid {
		return err
	}

	return nil
}
