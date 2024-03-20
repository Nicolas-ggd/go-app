package http

type ErrorResponse struct {
	Field string `json:"field"`
	Error string `json:"error"`
	Tag   string `json:"tag"`
	Param string `json:"param"`
}
