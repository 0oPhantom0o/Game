package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Internaluser struct {
	ID         primitive.ObjectID `json:"ID" bson:"_id"`
	Id         string             `json:"Id" bson:"Id"`
	Phone      string             `json:"phone" bson:"phone"`
	RandomCode string             `json:"randomCode" bson:"randomCode"`
	NickName   string             `json:"nickName" bson:"nickName"`
	Result     string             `json:"result" bson:"result"`
	Point      int                `json:"point" bson:"point"`
}
