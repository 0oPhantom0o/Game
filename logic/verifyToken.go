package logic

import (
	"fmt"
	"game/domain"
	"github.com/golang-jwt/jwt/v5"
	"strings"
)

func VerifyToken(reqToken string) (string, error) {
	splitToken := strings.Split(reqToken, "Bearer")
	if len(splitToken) != 2 {
		return "", fmt.Errorf("error")
	}

	tokenString := strings.TrimSpace(splitToken[1])

	var userClaim domain.UserClaim
	token, err := jwt.ParseWithClaims(tokenString, &userClaim, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", fmt.Errorf("invalid token")
	}
	fmt.Printf("Parsed User Claim: %s ", userClaim.ID)
	Id := userClaim.ID
	return Id, nil
}
