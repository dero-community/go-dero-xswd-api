package request

import (
	"github.com/dero-community/go-dero-xswd-api/xswd/rpc"
)

func Echo_Request() JSONRPCRequest[DaemonMethod, any] {
	return newJSONRPCRequest(Echo, nil)
}

func GetAddress_Request() JSONRPCRequest[DaemonMethod, any] {
	return newJSONRPCRequest(GetAddress, nil)
}

func GetBalance_Request(p rpc.GetBalance_Params) JSONRPCRequest[DaemonMethod, any] {
	return newJSONRPCRequest(GetBalance, p)
}

func GetHeight_Request(p rpc.GetHeight_Params) JSONRPCRequest[DaemonMethod, any] {
	return newJSONRPCRequest(GetHeight, p)
}

func GetTransferbyTXID_Request(p rpc.Get_Transfer_By_TXID_Params) JSONRPCRequest[DaemonMethod, any] {
	return newJSONRPCRequest(GetTransferbyTXID, p)
}

func GetTransfers_Request(p rpc.Get_Transfers_Params) JSONRPCRequest[DaemonMethod, any] {
	return newJSONRPCRequest(GetTransfers, p)
}

/*func GetTrackedAssets_Request(p rpc.Get_TrackedAssets_Params) JSONRPCRequest[DaemonMethod, any] {
	return newJSONRPCRequest(GetTrackedAssets, p)
}*/

func MakeIntegratedAddress_Request(p rpc.Make_Integrated_Address_Params) JSONRPCRequest[DaemonMethod, any] {
	return newJSONRPCRequest(MakeIntegratedAddress, p)
}

func SplitIntegratedAddress_Request(p rpc.Split_Integrated_Address_Params) JSONRPCRequest[DaemonMethod, any] {
	return newJSONRPCRequest(SplitIntegratedAddress, p)
}

func QueryKey_Request(p rpc.Query_Key_Params) JSONRPCRequest[DaemonMethod, any] {
	return newJSONRPCRequest(QueryKey, p)
}

func Transfer_Request(p rpc.Transfer_Params) JSONRPCRequest[DaemonMethod, any] {
	return newJSONRPCRequest(Transfer, p)
}

func SCInvoke_Request(p rpc.SC_Invoke_Params) JSONRPCRequest[DaemonMethod, any] {
	return newJSONRPCRequest(SCInvoke, p)
}
