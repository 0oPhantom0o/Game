package logic

import (
	"game/constants"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var secretKey = []byte(constants.JwtSecretKey)

func GenerateToken(id string) (string, error) {

	tokenGenerator := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"ID":       id,
			"IssuedAt": time.Now(),
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	token, err := tokenGenerator.SignedString(secretKey)

	if err != nil {
		return "", err

	}

	return token, nil
}
