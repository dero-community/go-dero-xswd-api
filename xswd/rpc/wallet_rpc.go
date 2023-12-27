// Copyright 2017-2021 DERO Project. All rights reserved.
// Use of this source code in any form is governed by RESEARCH license.
// license can be found in the LICENSE file.
// GPG: 0F39 E425 8C65 3947 702A  8234 08B2 0360 A03A 9DE8
//
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY
// EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL
// THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO,
// PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
// INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT,
// STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF
// THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

// this package contains only struct definitions required to implement wallet rpc (compatible with C daemon)
// in order to avoid the dependency on block chain by any package requiring access to rpc
// and other structures
// having the structures was causing the build times of explorer/wallet to be more than couple of secs
// so separated the logic from the structures

package rpc

import (
	"fmt"
	"math/big"
	"time"
)

// these structures are completely decoupled from blockchain and live only within the wallet
// all inputs and outputs which modify balance are presented by this structure
type Entry struct {
	Height         uint64    `json:"height"`
	TopoHeight     int64     `json:"topoheight"`
	BlockHash      string    `json:"blockhash"`
	MinerReward    uint64    `json:"minerreward"`
	TransactionPos int       `json:"tpos"` // pos within block is negative -1 for coinbase
	Pos            int       `json:"pos"`  // pos within transaction
	Coinbase       bool      `json:"coinbase"`
	Incoming       bool      `json:"incoming"`
	TXID           string    `json:"txid"`
	Destination    string    `json:"destination"`
	Burn           uint64    `json:"burn,omitempty"`
	Amount         uint64    `json:"amount"`
	Fees           uint64    `json:"fees"`
	Proof          string    `json:"proof"` // can be used to prove if available
	Status         byte      `json:"status"`
	Time           time.Time `json:"time"`
	EWData         string    `json:"ewdata"` // encrypted wallet balance at that point in time

	Data []byte `json:"data"` // data  is entire decrypted dump

	PayloadType  byte      `json:"payloadtype"`
	Payload      []byte    `json:"payload"`
	PayloadError string    `json:"payloaderror,omitempty"`
	Payload_RPC  Arguments `json:"payload_rpc,omitempty"`

	// these fields are only valid based on payload type  and if payload could be successfully parsed and will by default be equal to zero values
	Sender          string `json:"sender"`
	DestinationPort uint64 `json:"dstport"`
	SourcePort      uint64 `json:"srcport"`
}

// never do any division operation on money due to floating point issues
// newbies, see type the next in python interpretor "3.33-3.13"
func FormatMoney(amount uint64) string {
	return FormatMoneyPrecision(amount, 5) // default is 5 precision after floating point
}

// 0
func FormatMoney0(amount uint64) string {
	return FormatMoneyPrecision(amount, 0)
}

// 5 precision
func FormatMoney5(amount uint64) string {
	return FormatMoneyPrecision(amount, 5)
}

// 8 precision
func FormatMoney8(amount uint64) string {
	return FormatMoneyPrecision(amount, 8)
}

// 12 precision
func FormatMoney12(amount uint64) string {
	return FormatMoneyPrecision(amount, 12) // default is 8 precision after floating point
}

// format money with specific precision
func FormatMoneyPrecision(amount uint64, precision int) string {
	hard_coded_decimals := new(big.Float).SetInt64(100000)
	float_amount, _, _ := big.ParseFloat(fmt.Sprintf("%d", amount), 10, 0, big.ToZero)
	result := new(big.Float)
	result.Quo(float_amount, hard_coded_decimals)
	return result.Text('f', precision) // 5 is display precision after floating point
}

type (
	GetBalance_Params struct {
		SCID Hash `json:"scid"`
	} // no params
	GetBalance_Result struct {
		Balance          uint64 `json:"balance"`
		Unlocked_Balance uint64 `json:"unlocked_balance"`
	}
)

type (
	GetAddress_Params struct{} // no params
	GetAddress_Result struct {
		Address string `json:"address"`
	}
)

type (
	GetHeight_Params struct{} // no params
	GetHeight_Result struct {
		Height uint64 `json:"height"`
	}
)

// return type is string
type (
	Transfer struct {
		SCID        Hash      `json:"scid,omitempty"`
		Destination string    `json:"destination,omitempty"`
		Amount      uint64    `json:"amount,omitempty"`
		Burn        uint64    `json:"burn,omitempty"`
		Payload_RPC Arguments `json:"payload_rpc,omitempty"`
	}

	Transfer_Params struct {
		Transfers []Transfer `json:"transfers,omitempty"`
		SC_Code   string     `json:"sc,omitempty"`
		SC_Value  uint64     `json:"sc_value,omitempty"`
		SC_ID     string     `json:"scid,omitempty"`
		SC_RPC    Arguments  `json:"sc_rpc,omitempty"`
		Ringsize  uint64     `json:"ringsize,omitempty"`
		Fees      uint64     `json:"fees,omitempty"`
		Signer    string     `json:"signer,omitempty"` // only used for gas estimation
	}
	Transfer_Result struct {
		TXID string `json:"txid,omitempty"`
	}
)

type (
	SC_Invoke_Params struct {
		SC_ID            string    `json:"scid,omitempty"`
		SC_RPC           Arguments `json:"sc_rpc,omitempty"`
		SC_DERO_Deposit  uint64    `json:"sc_dero_deposit,omitempty"`
		SC_TOKEN_Deposit uint64    `json:"sc_token_deposit,omitempty"`
		Ringsize         uint64    `json:"ringsize,omitempty"`
	}
)

type (
	Get_Transfers_Params struct {
		SCID            Hash   `json:"scid,omitempty"`
		Coinbase        bool   `json:"coinbase,omitempty"`
		In              bool   `json:"in,omitempty"`
		Out             bool   `json:"out,omitempty"`
		Min_Height      uint64 `json:"min_height,omitempty"`
		Max_Height      uint64 `json:"max_height,omitempty"`
		Sender          string `json:"sender,omitempty"`
		Receiver        string `json:"receiver,omitempty"`
		DestinationPort uint64 `json:"dstport,omitempty"`
		SourcePort      uint64 `json:"srcport,omitempty"`
	}
	Get_Transfers_Result struct {
		Entries []Entry `json:"entries,omitempty"`
	}
)

// Get_Bulk_Payments
type (
	Get_Bulk_Payments_Params struct {
		Payment_IDs      []string `json:"payment_ids"`
		Min_block_height uint64   `json:"min_block_height"`
	}
	Get_Bulk_Payments_Result struct {
	}
)

// query_key
type (
	Query_Key_Params struct {
		Key_type string `json:"key_type"`
	}
	Query_Key_Result struct {
		Key string `json:"key"`
	}
)

// make_integrated_address_handler
type (
	Make_Integrated_Address_Params struct {
		Address     string    `json:"address"` // if its empty we assume wallets address
		Payload_RPC Arguments `json:"payload_rpc"`
	}
	Make_Integrated_Address_Result struct {
		Integrated_Address string    `json:"integrated_address"`
		Payload_RPC        Arguments `json:"payload_rpc"`
	}
)

// split_integrated_address_handler
type (
	Split_Integrated_Address_Params struct {
		Integrated_Address string `json:"integrated_address"`
	}
	Split_Integrated_Address_Result struct {
		Address     string    `json:"address"`
		Payload_RPC Arguments `json:"payload_rpc"`
	}
)

// Get_Transfer_By_TXID
type (
	Get_Transfer_By_TXID_Params struct {
		SCID Hash   `json:"scid"`
		TXID string `json:"txid,omitempty"`
	}
	Get_Transfer_By_TXID_Result struct {
		SCID  Hash  `json:"scid,omitempty"`
		Entry Entry `json:"entry,omitempty"`
	}
)
