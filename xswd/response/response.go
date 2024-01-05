package response

import (
	"github.com/dero-community/go-dero-xswd-api/xswd/shared"
)

type ResponseType string

const (
	Authorization ResponseType = "authorization"
	Event         ResponseType = "event"
	Command       ResponseType = "command"
)

type JSONRPCResponse struct {
	Jsonrpc shared.JsonRPCVersion `json:"jsonrpc"`
	Id      uint                  `json:"id"`
}

type ResultResponse[T any] struct {
	JSONRPCResponse
	Result T `json:"result"`
}

type ErrorCodeAndMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	JSONRPCResponse
	Error ErrorCodeAndMessage `json:"error"`
}

type Response struct {
	Type ResponseType
	Data JSONRPCResponse
}
