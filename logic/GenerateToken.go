package logic

import (
	"fmt"
	"game/constants"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var secretKey = []byte(constants.SecretKey)

func GenerateToken(Id string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"ID":  Id,
			"exp": time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	fmt.Println(tokenString)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
