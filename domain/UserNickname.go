package domain

type UserNickName struct {
	Phone    string `json:"phone" bson:"phone"`
	NickName string `json:"nickName" bson:"nickName"`
}
type UserScoreBoard struct {
	NickName string `json:"nickName" bson:"nickName"`
	Point    int    `json:"point" bson:"point"`
}
