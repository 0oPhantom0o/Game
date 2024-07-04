package domain

type InternalUser struct {
	Phone    string `json:"phone" bson:"phone"`
	NickName string `json:"nickName" bson:"nickName"`
	Point    int    `json:"point" bson:"point"`
}
