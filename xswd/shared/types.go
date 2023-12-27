package shared

type JsonRPCVersion string

const V2 JsonRPCVersion = "2.0"

type Event string

const (
	NewBalance    Event = "new_balance"
	NewEntry      Event = "new_entry"
	NewTopoheight Event = "new_topoheight"
)

var Events = []Event{NewBalance, NewEntry, NewTopoheight}
