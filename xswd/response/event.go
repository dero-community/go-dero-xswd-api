package response

import (
	"github.com/dero-community/go-dero-xswd-api/xswd/rpc"
	"github.com/dero-community/go-dero-xswd-api/xswd/shared"

	"github.com/mitchellh/mapstructure"
)

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

func DecodeEntry(value any) rpc.Entry {
	entry := rpc.Entry{}
	mapstructure.Decode(value, &entry)
	return entry
}

type BalanceEventValue struct {
	Balance uint     `json:"balance"`
	SCID    rpc.Hash `json:"scid"`
}

func DecodeBalance(value any) BalanceEventValue {
	balance := BalanceEventValue{}
	mapstructure.Decode(value, &balance)
	return balance
}

func DecodeHeight(value any) uint {
	return uint(value.(float64))
}
