package helpers

type Response struct {
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
	Error   error       `json:"error"`
}

func ResponseJSON(message string, code int, data interface{}, error error) Response {

	response := Response{
		Message: message,
		Status:  code,
		Data:    data,
		Error:   error,
	}

	return response
}
