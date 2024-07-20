package domain

type ErrorHandler struct {
	Err error `json:"error"`
}

type ResponseClient struct {
	Response string `json:"response"`
}

type ResponseToken struct {
	Token string `json:"jwt_token"`
}
