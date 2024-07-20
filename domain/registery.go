package domain

import (
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Phone      string `json:"phone"`
	RandomCode string `json:"random_code"`
}

type RequestPhone struct {
	Phone string `json:"phone"`
}

type RequestAnswer struct {
	Answer string `json:"answer"`
}

type RequestNickName struct {
	NickName string `json:"nick_name"`
}

type UserId struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
}
type UserClaim struct {
	jwt.RegisteredClaims
	ID string
}
