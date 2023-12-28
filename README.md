# go-dero-xswd-api

Go library to interface with the XSWD Protocol on DERO wallets.

## Usage


### Getting started

```go
package main

import (
	"dero-community/go-dero-xswd-api/xswd"
	"dero-community/go-dero-xswd-api/xswd/utils"
)

func main() {

  // create api object
  api := xswd.NewXSWD()

  const appName string = "myapp"
  const host string = "localhost:44326" // default

  // edit app info
  api.AppInfo = xswd.AppInfo{
    Id:          utils.GenerateAppId(appName),
    Name:        appName,
    Description: "myapp_description",
    Url:         "http://localhost",
  }

  // initialize
  // will open the websocket connection and authorize the app
  if err := api.Initialize(); err != nil {
    panic(err)
  }
}
```

### Subscribe to events

#### New Entry
```go
// Subscribe to new_entry event without callback
if err := api.Subscribe(shared.NewEntry, nil); err != nil {
  panic(err)
}
```

```go
// Subscribe to new_entry event with a callback
callback := func(value any) {
  // convert event value to Entry
  entry := response.DecodeEntry(value)

  // we can use entry as is from here...

  // [OPTIONAL] display entry as JSON in console
  if entryJSON, err := json.MarshalIndent(entry, "", "  "); err != nil {
    panic(err)
  } else {
    fmt.Println("new_entry: ", string(entryJSON))
  }
}

if err := api.Subscribe(shared.NewEntry, nil); err != nil {
  panic(err)
}
```


#### New Balance

```go
// Subscribe to new_topoheight event without callback
if err := api.Subscribe(shared.NewBalance, nil); err != nil {
  panic(err)
}
```

```go
// Subscribe to new_topoheight event with a callback
callback := func(value any) {
  balance := response.DecodeBalance(value)

  fmt.Println("new_balance: ", balance.Balance)
}

if err := api.Subscribe(shared.NewBalance, callback); err != nil {
  panic(err)
}

```


#### New Topoheight

```go
// Subscribe to new_topoheight event without callback
if err := api.Subscribe(shared.NewTopoheight, nil); err != nil {
  panic(err)
}
```

```go
// Subscribe to new_topoheight event with a callback
callback := func(value any) {
  height := response.DecodeHeight(value)

  fmt.Println("new_topoheight: ", height)
}

if err := api.Subscribe(shared.NewTopoheight, callback); err != nil {
  panic(err)
}

```

### Wait for events

Once subscribed to an event, you can wait for it with a predicate function to filter out some events.

```go
// Wait for new balance
api.WaitFor(shared.NewBalance, nil)
```

```go
// Make a transfer of 0.1 DERO 
// (Transfer is a utility function from the "dero-community/go-dero-xswd-api/xswd/utils" module)
resp := utils.Transfer(api, recipient_address, 10000)

// Create a predicate function that checks if the entry's TXID is the same as the transfer's TXID
predicate := func(value any) bool {
  // Decode entry
  entry := response.DecodeEntry(value)

  // Check txid is the one we expect
  return entry.TXID == resp.Result.TXID
}

// Wait for new entry satisfying the predicate
api.WaitFor(shared.NewEntry, predicate)
```

### Node / Wallet Methods

Requests can be built using helper functions in the `dero-community/go-dero-xswd-api/xswd/request` module.

Lets see an example with the `GetAddress` method

#### Request
```go
// Build request (get address has no arguments)
req := request.GetAddress_Request()
// Send request, we get a channel on which we can wait for the response
responseChannel := api.Send(req)
// Wait for the response. Can be a result or an error
responseDataOrError := <-responseChannel
// Handle the error and get the raw response 
resp := utils.HandleDataOrError(responseDataOrError)
```
Can be written shorter as:
```go
req := request.GetAddress_Request()
resp := utils.HandleDataOrError(<-api.Send(req))
```

#### Parse the response

```go
// Parse the response
resultResponse := response.ResultResponse[rpc.GetAddress_Result]{}
utils.HandleError(json.Unmarshal(resp, &resultResponse))

// Print address
fmt.Println("address:", resultResponse.Result.Address)
```


## Development & Tests

To try it out simply run the DERO `simulator` with `--use-xswd` parameter and launch test with `go run test.go`. This file contains multiple usage examples.
