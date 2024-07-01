package domain

type InternalUser struct {
	//ID       primitive.ObjectID `json:"id" bson:"_id omitempty"`
	Phone    string `json:"phone" bson:"phone"`
	NickName string `json:"nickName" bson:"nickName"`
	Point    int    `json:"point" bson:"point"`
}
