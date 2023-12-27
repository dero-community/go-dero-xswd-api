package request

type Entity = string

const (
	Daemon Entity = "daemon"
	//node   Entity = "daemon"
	Wallet Entity = "wallet"
)

const (
	DERO_Echo                       DaemonMethod = "DERO.Echo"
	DERO_Ping                       DaemonMethod = "DERO.Ping"
	DERO_GetInfo                    DaemonMethod = "DERO.GetInfo"
	DERO_GetBlock                   DaemonMethod = "DERO.GetBlock"
	DERO_GetBlockHeaderByTopoHeight DaemonMethod = "DERO.GetBlockHeaderByTopoHeight"
	DERO_GetBlockHeaderByHash       DaemonMethod = "DERO.GetBlockHeaderByHash"
	DERO_GetTxPool                  DaemonMethod = "DERO.GetTxPool"
	DERO_GetRandomAddress           DaemonMethod = "DERO.GetRandomAddress"
	DERO_GetTransaction             DaemonMethod = "DERO.GetTransaction"
	DERO_SendRawTransaction         DaemonMethod = "DERO.SendRawTransaction"
	DERO_GetHeight                  DaemonMethod = "DERO.GetHeight"
	DERO_GetBlockCount              DaemonMethod = "DERO.GetBlockCount"
	DERO_GetLastBlockHeader         DaemonMethod = "DERO.GetLastBlockHeader"
	DERO_GetBlockTemplate           DaemonMethod = "DERO.GetBlockTemplate"
	DERO_GetEncryptedBalance        DaemonMethod = "DERO.GetEncryptedBalance"
	DERO_GetSC                      DaemonMethod = "DERO.GetSC"
	DERO_GetGasEstimate             DaemonMethod = "DERO.GetGasEstimate"
	DERO_NameToAddress              DaemonMethod = "DERO.NameToAddress"

	Echo                   WalletMethod = "Echo"
	GetAddress             WalletMethod = "GetAddress"
	GetBalance             WalletMethod = "GetBalance"
	GetHeight              WalletMethod = "GetHeight"
	GetTransferbyTXID      WalletMethod = "GetTransferbyTXID"
	GetTransfers           WalletMethod = "GetTransfers"
	GetTrackedAssets       WalletMethod = "GetTrackedAssets"
	MakeIntegratedAddress  WalletMethod = "MakeIntegratedAddress"
	SplitIntegratedAddress WalletMethod = "SplitIntegratedAddress"
	QueryKey               WalletMethod = "QueryKey"
	Transfer               WalletMethod = "transfer"
	SCInvoke               WalletMethod = "scinvoke"

	Subscribe WalletMethod = "Subscribe"
)
