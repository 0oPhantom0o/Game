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

	// Split the token string by "Bearer " (note the space after "Bearer")
	splitToken := strings.SplitN(reqToken, "Bearer ", 2)

	// Check if we have the expected number of parts
	if len(splitToken) != 2 {
		return "", fmt.Errorf("invalid token format")
	}

	// The actual token string will be the second part
	tokenString := strings.TrimSpace(splitToken[1])

	var userClaim domain.UserClaim
	token, err := jwt.ParseWithClaims(tokenString, &userClaim, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return "", fmt.Errorf("problem in decoding jwt: %v", err)
	}
	if !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	Id := userClaim.ID
	return Id, nil
}
