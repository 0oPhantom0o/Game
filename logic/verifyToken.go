package logic

import (
	"fmt"
	"game/domain"
	"github.com/golang-jwt/jwt/v5"
	"strings"
)

func VerifyToken(reqToken string) (string, error) {
	if reqToken == "" {
		return "", fmt.Errorf("empty token")
	}
	splitToken := strings.Split(reqToken, "Bearer")

	tokenString := strings.TrimSpace(splitToken[1])

	var userClaim domain.UserClaim
	token, err := jwt.ParseWithClaims(tokenString, &userClaim, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return "", fmt.Errorf("problem in decoding jwt")
	}
	if !token.Valid {
		return "", fmt.Errorf("invalid token")
	}
	Id := userClaim.ID
	return Id, nil
}
