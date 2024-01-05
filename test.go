package main

import (
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/url"

	"github.com/dero-community/go-dero-xswd-api/xswd"
	"github.com/dero-community/go-dero-xswd-api/xswd/request"
	"github.com/dero-community/go-dero-xswd-api/xswd/response"
	"github.com/dero-community/go-dero-xswd-api/xswd/rpc"
	"github.com/dero-community/go-dero-xswd-api/xswd/shared"
	"github.com/dero-community/go-dero-xswd-api/xswd/utils"
)

func main() {
	// Enable logging
	xswd.LoggerHandler.SetEnable(true)
	// Set log level of package to debug
	xswd.LoggerOpts.SlogOpts.Level = slog.LevelDebug

	log.Println("Get another wallet address")
	address2 := getOtherWalletAddress()

	log.Println("address:", address2)

	const appName string = "myapp"
	//const host string = "localhost:44326" // default
	const host string = "localhost:40000"

	// create api object
	api := xswd.NewXSWD()

	api.Addr = url.URL{Scheme: "ws", Host: host, Path: "/xswd"}
	api.AppInfo = xswd.AppInfo{
		Id:          utils.GenerateAppId(appName),
		Name:        appName,
		Description: "myapp_description",
		Url:         "http://localhost",
	}

	// initialize connection (creates websocket + authorization process)
	utils.HandleError(api.Initialize())

	//// subscribe to new_entry events to wait for transfers to complete (without callback, we will hande the results with WaitFor)
	//utils.HandleError(api.Subscribe(shared.NewEntry, nil))

	// subscribe to new_entry events to wait for transfers to complete (with callback)
	{
		callback := func(value any) {
			// convert event value to Entry
			entry := response.DecodeEntry(value)

			// we can use entry as is from here...

			// display entry as JSON in console
			if entryJSON, err := json.MarshalIndent(entry, "", "  "); err != nil {
				panic(err)
			} else {
				fmt.Println("new_entry: ", string(entryJSON))
			}
		}

		utils.HandleError(api.Subscribe(shared.NewEntry, callback))
	}

	// Echo
	utils.HandleDataOrError(<-api.Send(
		request.DERO_Echo_Request([]string{"Hello,", "World"}),
	))

	// Ping
	utils.HandleDataOrError(<-api.Send(
		request.DERO_Ping_Request(),
	))

	// GetInfo
	utils.HandleDataOrError(<-api.Send(
		request.DERO_GetInfo_Request(),
	))

	// GetAddress
	utils.HandleDataOrError(<-api.Send(
		request.GetAddress_Request(),
	))

	// create a transfer helper function

	// Transfer 0.1 DERO to another wallet
	resp := utils.Transfer(api, address2, 10000)

	// wait for entry with the same txid
	api.WaitFor(shared.NewEntry,
		func(value any) bool {
			return response.DecodeEntry(value).TXID == resp.Result.TXID
		},
	)

	// Subscribe to new_balance event
	api.Subscribe(shared.NewBalance, func(value any) {
		balance := response.DecodeBalance(value)

		log.Println("MAIN: new_topoheight callback()", balance)
	})

	// Transfer 0.1 DERO to another wallet
	utils.Transfer(api, address2, 10000)

	// Wait for a new balance event (no callback, already set in subscribe)
	api.WaitFor(shared.NewBalance, nil)

	{
		// Subscribe to new_topoheight event
		callback := func(value any) {
			height := response.DecodeHeight(value)
			log.Println("MAIN: new_topoheight callback()", height)
		}
		api.Subscribe(shared.NewTopoheight, callback)
	}

	api.WaitFor(shared.NewTopoheight, nil)

}

// Helper to get another wallet address (using an other instance of XSWD)
func getOtherWalletAddress() string {
	const appName string = "myapp"
	const host string = "localhost:40001"

	// create api object
	api := xswd.NewXSWD()

	api.Addr = url.URL{Scheme: "ws", Host: host, Path: "/xswd"}
	api.AppInfo = xswd.AppInfo{
		Id:          utils.GenerateAppId(appName),
		Name:        appName,
		Description: "myapp_description",
		Url:         "http://localhost",
	}

	// initialize connection (creates websocket + authorization process)
	utils.HandleError(api.Initialize())
	utils.HandleError(api.Subscribe(shared.NewEntry, nil))

	req := request.GetAddress_Request()
	resp := utils.HandleDataOrError(<-api.Send(req))

	resultResponse := response.ResultResponse[rpc.GetAddress_Result]{}
	utils.HandleError(json.Unmarshal(resp, &resultResponse))
	fmt.Println("address:", resultResponse.Result.Address)
	return resultResponse.Result.Address
}
