package domain

import (
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InternalUser struct {
	Phone         string `json:"phone" bson:"phone"`
	NickName      string `json:"nick_name" bson:"nick_name"`
	NickNameLimit int    `bson:"nick_name_limit"`
	Point         int    `json:"point" bson:"point"`
}

type TopPlayers struct {
	NickName string `json:"nick_name" bson:"nick_name"`
	Point    int    `json:"point" bson:"point"`
}
type Score struct {
	NextPage bool         `json:"next_page"`
	Players  []TopPlayers `json:"players"`
}
type UserClaim struct {
	jwt.RegisteredClaims
	ID string
}
type ResponseToken struct {
	Token string `json:"jwt_token"`
}

type UserId struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
}
