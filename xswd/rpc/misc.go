package rpc

const HashLength = 32

// type Hash [HashLength]byte
type Hash string

type DataType string

const (
	DataString  DataType = "S"
	DataInt64   DataType = "I"
	DataUint64  DataType = "U"
	DataFloat64 DataType = "F"
	DataHash    DataType = "H" // a 256 bit hash (basically sha256 of 32 bytes long)
	DataAddress DataType = "A" // dero address represented in 33 bytes
	DataTime    DataType = "T"
)

// type  DataType byte
type Argument struct {
	Name     string      `json:"name"`     // string name must be atleast 1 byte
	DataType DataType    `json:"datatype"` // Type must one of the known data types
	Value    interface{} `json:"value"`    // value should be as per type
}

type Arguments []Argument
