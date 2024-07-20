package domain

type ErrorHandler struct {
	Err error `json:"err"`
}

type ResponseClient struct {
	Response string `json:"response"`
}
