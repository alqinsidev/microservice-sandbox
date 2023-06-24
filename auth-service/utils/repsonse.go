package utils

type Success struct {
	Data    interface{} `json:"data"`
	Status  int         `json:"status"`
	Message string      `json:"message"`
}

func SuccessResponse(data interface{}, status int, message string) *Success {
	if message == "" {
		message = "success"
	}
	return &Success{
		Data:    data,
		Status:  status,
		Message: message,
	}
}

type Error struct {
	Status  int    `json:"statusCode"`
	Message string `json:"message"`
}

func ErrorResponse(status int, message string) *Error {
	return &Error{
		Status:  status,
		Message: message,
	}
}
