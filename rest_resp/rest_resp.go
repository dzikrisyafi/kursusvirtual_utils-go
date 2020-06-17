package rest_resp

import "net/http"

type ResponseJSON interface {
	Message() string
	Status() int
	Result() interface{}
}

type responseJSON struct {
	RespMessage string      `json:"message"`
	RespStatus  int         `json:"status"`
	RespResult  interface{} `json:"result"`
}

func (r responseJSON) Message() string {
	return r.RespMessage
}

func (r responseJSON) Status() int {
	return r.RespStatus
}

func (r responseJSON) Result() interface{} {
	return r.RespResult
}

func NewStatusOK(message string, result interface{}) ResponseJSON {
	return responseJSON{
		RespMessage: message,
		RespStatus:  http.StatusOK,
		RespResult:  result,
	}
}

func NewStatusCreated(message string, result interface{}) ResponseJSON {
	return responseJSON{
		RespMessage: message,
		RespStatus:  http.StatusCreated,
		RespResult:  result,
	}
}
