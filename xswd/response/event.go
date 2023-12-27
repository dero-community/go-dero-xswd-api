package response

import "dero-community/go-dero-xswd-api/xswd/shared"

type SubscribeResponse struct {
	JSONRPCResponse
}

type EventMessageResult[T any] struct {
	Event shared.Event `json:"event"`
	Value T            `json:"value"`
}

type EventMessage[T any] struct {
	JSONRPCResponse
	Result EventMessageResult[T] `json:"result"`
}
