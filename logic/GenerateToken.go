package logic

import (
	"game/constants"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var secretKey = []byte(constants.SecretKey)

func GenerateToken(Id string) (string, error) {
	tokenGenerator := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"ID":  Id,
			"exp": time.Now().Add(time.Hour * 24).Unix(),
		})

	token, err := tokenGenerator.SignedString(secretKey)

	if err != nil {
		return "", err
	}

	return token, nil
}
