package domain

type InternalUser struct {
	Phone         string `json:"phone" bson:"phone"`
	NickName      string `json:"nickName" bson:"nickName"`
	NickNameLimit int    `bson:"nickNameLimit"`
	Point         int    `json:"point" bson:"point"`
}

type TopPlayers struct {
	NickName string `json:"nickName" bson:"nickName"`
	Point    int    `json:"point" bson:"point"`
}
type Score struct {
	Players  []TopPlayers
	NextPage bool
}
