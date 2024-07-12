package domain

type UserNickName struct {
	Phone    string `json:"phone" bson:"phone"`
	NickName string `json:"nickName" bson:"nickName"`
}
