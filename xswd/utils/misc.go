package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"log"

	"github.com/dero-community/go-dero-xswd-api/xswd"
	"github.com/dero-community/go-dero-xswd-api/xswd/request"
	"github.com/dero-community/go-dero-xswd-api/xswd/response"
	"github.com/dero-community/go-dero-xswd-api/xswd/rpc"
)

// Transfer a DERO amount to an other wallet
func Transfer(api *xswd.XSWD, to string, amount uint64) response.ResultResponse[rpc.Transfer_Result] {
	// Build transfer request
	tranferRequest := request.Transfer_Request(
		rpc.Transfer_Params{
			Transfers: []rpc.Transfer{
				{
					Destination: to,
					Amount:      amount,
				},
			},
		})

	// Send it
	dataOrErrorResponse := <-api.Send(tranferRequest)

	// Handle errors
	raw_response := HandleDataOrError(dataOrErrorResponse)

	// Parse response
	resp := response.ResultResponse[rpc.Transfer_Result]{}
	HandleError(json.Unmarshal(raw_response, &resp))

	// Return response
	return resp
}

// Handle error
func HandleError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// Handle error in xswd.DataOrError struct
func HandleDataOrError[T any](d xswd.DataOrError[T]) T {
	if d.Err != nil {
		log.Fatalln(d.Err)
	}
	return d.Data
}

// generate 64 alphanum character APP ID from app name with sha256
func GenerateAppId(name string) string {
	running_hash := sha256.New()     // type hash.Hash
	running_hash.Write([]byte(name)) // data is []byte
	sum := running_hash.Sum([]byte(name))
	return hex.EncodeToString(sum)[0:64]
}
