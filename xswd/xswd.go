package xswd

import (
	"dero-community/go-dero-xswd-api/xswd/request"
	"dero-community/go-dero-xswd-api/xswd/response"
	"dero-community/go-dero-xswd-api/xswd/shared"
	"encoding/json"
	"errors"
	"log"
	"net/url"
	"regexp"
	"slices"
	"time"

	"github.com/gorilla/websocket"
)

// XSWD Api status
type Status = string

const (
	Idle        Status = "idle"
	Authorizing Status = "authorizing"
	Connected   Status = "connected"
	Ready       Status = "ready"
	//Fallback    Status = "fallback"
)

// Application information sent during authorization
type AppInfo struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Url         string `json:"url"`
}

// Response from a authorize request
type AuthResponse struct {
	Message  string `json:"message"`
	Accepted bool   `json:"accepted"`
}

// Subscibtion to an event
type Subscription struct {
	// Channel for WaitFor function to pickup messages
	Event chan response.EventMessage[any]
	// Enabled flag
	Enabled bool
	// User callback function
	Callback func(any)
}

// XSWD Api main structure
// Use NewXSWD() to create it
type XSWD struct {
	//FallbackUrl url.URL 	// Fallback HTTP // TODO
	Addr          url.URL                        // Websocket address
	status        Status                         // Status
	connection    *websocket.Conn                // Websocket connection
	response      chan []byte                    // Response channel
	AppInfo       AppInfo                        // Application information that can be edited by the user
	subscriptions map[shared.Event]*Subscription // Subscription map
	nextId        uint                           // Id variable incrementing with each call to SendSync
}

// Creates a XSWD Api struct with default values
// You can edit the Addr & AppInfo fields after
func NewXSWD() *XSWD {

	defaultAddress := "localhost:44326"
	u := url.URL{Scheme: "ws", Host: defaultAddress, Path: "/xswd"}

	xswd := XSWD{
		status:        Idle,
		Addr:          u,
		connection:    nil,
		response:      make(chan []byte),
		subscriptions: make(map[shared.Event]*Subscription),
		nextId:        0,
	}

	for _, event := range shared.Events {
		xswd.subscriptions[event] = &Subscription{
			Enabled:  false,
			Callback: nil,
			Event:    make(chan response.EventMessage[any]),
		}

	}

	return &xswd
}

// Getter for the status of the api
func (xswd XSWD) Status() Status {
	return xswd.status
}

// ! fix for api returning either strings or ints on the Id field of the JsonRPC requests
// Casts an empty string to 0
// Removes double quotes on a numeric ID
func fixMessage(message []byte) []byte {
	m1 := regexp.MustCompile(`"id":""`)
	repl1 := []byte(`"id":0`)
	message = m1.ReplaceAll(message, repl1)
	m2 := regexp.MustCompile(`"id":"(\d+)"`)
	repl2 := []byte(`"id":$1`)
	return m2.ReplaceAll(message, repl2)
}

// Initializes the API
// Mandatory for later calls
func (xswd *XSWD) Initialize() error {
	// Check status
	if xswd.status == "Ready" {
		Logger.Warn("attempt to initialize an api again")
		return errors.New("api is already initialized")
	}

	// Check AppInfo is properly formatted
	if !checkAppInfo(xswd.AppInfo) {
		return errors.New("app info incomplete")
	}

	// Create websocket
	if c, _, err := websocket.DefaultDialer.Dial(xswd.Addr.String(), nil); err != nil {
		return errors.New("failed to initialize: " + err.Error())
	} else {
		xswd.connection = c
	}

	// Authorize
	xswd.status = Authorizing
	if res, err := xswd.authorize(); err != nil || !res {
		if err != nil {
			return errors.New("failed to authorize: " + err.Error())
		} else {
			return errors.New("app Info wrong format")
		}
	}

	xswd.status = Connected
	Logger.Info("authorize accepted")

	// Setup message listener
	go func() {
		for {
			// Read message
			if _, message, err := xswd.connection.ReadMessage(); err != nil {
				Logger.Error("read err:", err)
				return
			} else {
				message = fixMessage(message) //! FIX for id returning as string

				Logger.Debug("recv:", "message", string(message))

				// Attempt to parse event
				if eventMessage, err := parseEvent(message); err == nil {
					// If subscribed to this event
					if xswd.subscriptions[eventMessage.Result.Event].Enabled {
						// Send event message to event channel
						xswd.subscriptions[eventMessage.Result.Event].Event <- eventMessage
						// Call callback function if set up
						if callback := xswd.subscriptions[eventMessage.Result.Event].Callback; callback != nil {
							callback(eventMessage.Result.Value)
						}
					}
				} else {
					// Send mesage to response channel
					xswd.response <- message
				}
			}
		}
	}()
	return nil
}

// Checks application information
func checkAppInfo(appInfo AppInfo) bool {
	return appInfo.Description != "" && appInfo.Id != "" && appInfo.Name != ""
}

// Parse authentification response message
func parseAuth(message []byte) (AuthResponse, error) {
	response := AuthResponse{}
	return response, json.Unmarshal(message, &response)
}

// Parse event message
func parseEvent(message []byte) (response.EventMessage[any], error) {
	resp := response.EventMessage[any]{}
	if err := json.Unmarshal(message, &resp); err != nil {
		Logger.Warn("failed to parse event", "message", string(message), "error", err)
		return resp, err
	} else {
		//? Parse can be successful with empty event field
		// Check if event field is not empty or exists
		if slices.Contains(shared.Events, resp.Result.Event) {
			return resp, err
		}
		Logger.Warn("failed to parse event: event unknown or field is empty", "message", message)
		return resp, errors.New("failed to parse event: event unknown or field is empty")
	}
}

// Close connection
func (xswd *XSWD) Close() {
	if xswd.status != Idle {
		err := xswd.connection.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		if err != nil {
			Logger.Error("write close:", err)
		}
	}
}

// Generic channel stuct to send error or data
type DataOrError[T any] struct {
	Err  error
	Data T
}

// Response or error channel
type ResponseOrError = DataOrError[[]byte]

// Sends a request.
// returns a channel that the user can wait on for a response
func (xswd *XSWD) Send(data request.JSONRPCRequest[request.Method, any]) <-chan ResponseOrError {
	// Check status
	if xswd.status != Connected {
		Logger.Warn("sending without being initialized")
		return nil
	}

	// Make the response or error channel
	r := make(chan ResponseOrError)

	// Put new id on request
	data.Id = xswd.nextId
	xswd.nextId = xswd.nextId + 1

	// Marshal request to text (byte array)
	if jsonData, err := json.Marshal(data); err != nil {
		log.Println(err)
		return nil
	} else {
		Logger.Debug("SendSync send:", "jsonData", data)

		// Send the raw data to the websocket
		xswd.send(jsonData)

		// Setup handler
		go func() {
			defer close(r)

			// Fail helper
			fail := func(xr []byte, err error) {
				Logger.Error("failed to parse:", string(xr), err)
				// Send the message back to the response channel
				xswd.response <- xr
				// Wait for a while
				time.Sleep(500 * time.Millisecond)
			}

			// Listen for responses
			for {
				xr := <-xswd.response
				// Unmarshall response
				resp := response.ResultResponse[any]{}
				if err := json.Unmarshal(xr, &resp); err != nil {
					// Check for an error field in the response
					errorResp := response.ErrorResponse{}
					if err := json.Unmarshal(xr, &errorResp); err != nil {
						//? discard the message
						// TODO check this
						// Send the message back if no error is found
						fail(xr, err)
					} else {
						// Response had an error field
						Logger.Error("error in response:", "error", errorResp.Error)
						// Send it in the response channel and return
						r <- ResponseOrError{Err: errors.New(errorResp.Error.Message)}
						return
					}
				} else {
					// Response was unmarhsalled correctly
					// Check if id corresponds
					if resp.Id != data.Id {
						fail(xr, err)
					} else {
						Logger.Debug("SendSync recv:", "response", resp)
						r <- ResponseOrError{Data: xr}
						return
					}
				}
			}
		}()

		// Return reponse channel
		return r
	}
}

// Subscribe to an Event.
// You can attach a callback function (or nil) that will be called whenever this particular event is fired.
func (xswd XSWD) Subscribe(event shared.Event, callback func(any)) error {
	Logger.Info("subscribing to", "event", event)
	if xswd.subscriptions[event].Enabled {
		Logger.Warn("warning: already subscribed to", "event", event)
	}

	r := xswd.Send(request.SubscribeRequest(event))
	Logger.Info("waiting for response")
	if response := <-r; response.Err != nil {
		Logger.Error("subscribe error:", response.Err)
		return response.Err
	} else {
		Logger.Info("successfully subscribed to", "event", event)
		xswd.subscriptions[event].Enabled = true
		if callback != nil {
			Logger.Info("setting up callback for", "event", event)
			xswd.subscriptions[event].Callback = callback
		}
		return nil
	}

}

// Sends raw data to the websocket (or fallback) //TODO
func (xswd XSWD) send(data []byte) {
	err := xswd.connection.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		Logger.Error("write:", err)
		return
	}
}

// Authorizes the app
func (xswd XSWD) authorize() (bool, error) {
	Logger.Info("auth:", "appInfo", xswd.AppInfo)
	if data, err := json.Marshal(xswd.AppInfo); err != nil {
		Logger.Error("marshal error:", err)
		return false, err
	} else {
		xswd.send(data)
		_, message, err := xswd.connection.ReadMessage()
		if err != nil {
			Logger.Error("read err:", err)
			return false, err
		}
		if authResponse, err := parseAuth(message); err != nil {
			Logger.Error("auth err:", err)
			return false, err
		} else {
			if authResponse.Accepted {
				return true, nil
			} else {
				return false, errors.New(authResponse.Message)
			}
		}
	}
}

// Waits for a given event if the predicate returns true
// A nil predicate function will not filter the events
func (xswd XSWD) WaitFor(event shared.Event, predicate func(any) bool) {
	Logger.Info("waiting for event", "event", event)

	eventMessage := <-xswd.subscriptions[event].Event
	if predicate != nil {

		for !predicate(eventMessage.Result.Value) {
			xswd.subscriptions[event].Event <- eventMessage
			time.Sleep(500 * time.Millisecond)
			eventMessage = <-xswd.subscriptions[event].Event
		}
	}
	Logger.Info("event found", "event", event)
}
