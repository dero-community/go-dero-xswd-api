package request

import "dero-community/go-dero-xswd-api/xswd/shared"

type Method string
type DaemonMethod = Method
type WalletMethod = Method

type JSONRPCRequest[M Method, T any] struct {
	Jsonrpc shared.JsonRPCVersion `json:"jsonrpc"`
	Method  M                     `json:"method"`
	Params  T                     `json:"params"`
	Id      uint                  `json:"id"`
}

func newJSONRPCRequest(method Method, p any) JSONRPCRequest[DaemonMethod, any] {
	return JSONRPCRequest[DaemonMethod, any]{
		Jsonrpc: shared.V2,
		Method:  method,
		Params:  p,
	}
}
