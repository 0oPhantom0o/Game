package domain

import (
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID         primitive.ObjectID `json:"ID" bson:"_id omitempty"`
	Phone      string             `json:"phone" bson:"phone"`
	RandomCode string             `json:"randomCode" bson:"randomCode"`
	NickName   string             `json:"nickName" bson:"nickName"`
	Result     string             `json:"result" bson:"result"`
}

type UserClaim struct {
	jwt.RegisteredClaims
	ID string
}
