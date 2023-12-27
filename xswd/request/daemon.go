package request

import (
	"dero-community/go-dero-xswd-api/xswd/rpc"
)

func DERO_Echo_Request(p []string) JSONRPCRequest[DaemonMethod, any] {
	return newJSONRPCRequest(DERO_Echo, p)
}

func DERO_Ping_Request() JSONRPCRequest[DaemonMethod, any] {
	return newJSONRPCRequest(DERO_Ping, nil)
}

func DERO_GetInfo_Request() JSONRPCRequest[DaemonMethod, any] {
	return newJSONRPCRequest(DERO_GetInfo, nil)
}

func DERO_GetBlock_Request(p rpc.GetBlock_Params) JSONRPCRequest[DaemonMethod, any] {
	return newJSONRPCRequest(DERO_GetInfo, p)
}

func DERO_GetBlockHeaderByTopoHeight_Request(p rpc.GetBlockHeaderByTopoHeight_Params) JSONRPCRequest[DaemonMethod, any] {
	return newJSONRPCRequest(DERO_GetBlockHeaderByTopoHeight, p)
}

func DERO_GetBlockHeaderByHash_Request(p rpc.GetBlockHeaderByHash_Params) JSONRPCRequest[DaemonMethod, any] {
	return newJSONRPCRequest(DERO_GetBlockHeaderByHash, p)
}

func DERO_GetTxPool_Request(p rpc.GetTxPool_Params) JSONRPCRequest[DaemonMethod, any] {
	return newJSONRPCRequest(DERO_GetTxPool, p)
}

func DERO_GetRandomAddress_Request(p rpc.GetRandomAddress_Params) JSONRPCRequest[DaemonMethod, any] {
	return newJSONRPCRequest(DERO_GetRandomAddress, p)
}

func DERO_GetTransaction_Request(p rpc.GetTransaction_Params) JSONRPCRequest[DaemonMethod, any] {
	return newJSONRPCRequest(DERO_GetTransaction, p)
}

func DERO_SendRawTransaction_Request(p rpc.SendRawTransaction_Params) JSONRPCRequest[DaemonMethod, any] {
	return newJSONRPCRequest(DERO_SendRawTransaction, p)
}

func DERO_GetHeight_Request(p rpc.GetHeight_Params) JSONRPCRequest[DaemonMethod, any] {
	return newJSONRPCRequest(DERO_GetHeight, p)
}

func DERO_GetBlockCount_Request(p rpc.GetBlockCount_Params) JSONRPCRequest[DaemonMethod, any] {
	return newJSONRPCRequest(DERO_GetBlockCount, p)
}

func DERO_GetLastBlockHeader_Request(p rpc.GetLastBlockHeader_Params) JSONRPCRequest[DaemonMethod, any] {
	return newJSONRPCRequest(DERO_GetLastBlockHeader, p)
}

func DERO_GetBlockTemplate_Request(p rpc.GetBlockTemplate_Params) JSONRPCRequest[DaemonMethod, any] {
	return newJSONRPCRequest(DERO_GetBlockTemplate, p)
}

func DERO_GetEncryptedBalance_Request(p rpc.GetEncryptedBalance_Params) JSONRPCRequest[DaemonMethod, any] {
	return newJSONRPCRequest(DERO_GetEncryptedBalance, p)
}

func DERO_GetSC_Request(p rpc.GetSC_Params) JSONRPCRequest[DaemonMethod, any] {
	return newJSONRPCRequest(DERO_GetSC, p)
}

func DERO_GetGasEstimate_Request(p rpc.GasEstimate_Params) JSONRPCRequest[DaemonMethod, any] {
	return newJSONRPCRequest(DERO_GetGasEstimate, p)
}

func DERO_NameToAddress_Request(p rpc.NameToAddress_Params) JSONRPCRequest[DaemonMethod, any] {
	return newJSONRPCRequest(DERO_NameToAddress, p)
}
