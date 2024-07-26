package domain

type RequestRandomCode struct {
	Phone      string `json:"phone"`
	RandomCode string `json:"random_code"`
}

type RequestPhone struct {
	Phone string `json:"phone"`
}

type RequestAnswer struct {
	Answer string `json:"answer"`
}

type RequestNickName struct {
	NickName string `json:"nick_name"`
}
