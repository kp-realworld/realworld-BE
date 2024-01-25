package types

type ErrorResponse struct {
	Errors ErrorDetail `json:"error"`
}

type ErrorDetail struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
