package request

import "github.com/dero-community/go-dero-xswd-api/xswd/shared"

type SubscribeParams struct {
	Event shared.Event `json:"event"`
}

func SubscribeRequest(event shared.Event) JSONRPCRequest[WalletMethod, any] {
	return newJSONRPCRequest(Subscribe, SubscribeParams{
		Event: event,
	})
}
