package domain

import (
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID         primitive.ObjectID `json:"ID" bson:"_id,omitempty"`
	Phone      string             `json:"phone" bson:"phone"`
	RandomCode string             `json:"randomCode" bson:"randomCode"`
}

type RequestPhone struct {
	Phone string `json:"phone" bson:"phone"`
}
type UserId struct {
	ID primitive.ObjectID `json:"ID" bson:"_id,omitempty"`
}
type UserAnswer struct {
	Answer string `json:"answer" bson:"answer"`
}

type RequestNickName struct {
	NickName string `json:"nickName" bson:"nickName"`
}

type UserClaim struct {
	jwt.RegisteredClaims
	ID string
}
