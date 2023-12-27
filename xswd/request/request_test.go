package request

import (
	"dero-community/go-dero-xswd-api/xswd/shared"
	"encoding/json"
	"testing"
)

func TestSubscribeMarshal(t *testing.T) {

	sub := JSONRPCRequest[WalletMethod, SubscribeParams]{
		Jsonrpc: shared.V2,
		Method:  Subscribe,
		Id:      0,
		Params: SubscribeParams{
			Event: shared.NewBalance,
		},
	}

	if v, err := json.Marshal(sub); err != nil {
		t.Fatal(err)
	} else {
		const expected string = `{"jsonrpc":"2.0","method":"Subscribe","params":{"event":"new_balance"},"id":0}`
		if string(v) != expected {
			t.Fatal("\nexpected", expected, "\ngot", string(v))
		}
	}

}

func TestDEROEchoRequest(t *testing.T) {
	p := []string{"Hello,", "world!"}
	sub := DERO_Echo_Request(p)

	if v, err := json.Marshal(sub); err != nil {
		t.Fatal(err)
	} else {
		const expected string = `{"jsonrpc":"2.0","method":"DERO.Echo","params":["Hello,","world!"],"id":0}`
		if string(v) != expected {
			t.Fatal("\nexpected", expected, "\ngot", string(v))
		}
	}

}
