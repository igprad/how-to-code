package error

type error struct {
	ErrorCode    string `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

func ToErrorResponse(ec string, em string) error {
	return error{ec, em}
}
