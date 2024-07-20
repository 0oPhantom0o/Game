package domain

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
