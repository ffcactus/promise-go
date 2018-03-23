package dto

import (
	"encoding/json"
)
// PostEventRequest is the DTO that includes the payload that expected to be send
// from websocket.
type PostEventRequest struct {
	Category string
	Type string
	ResourceID string
	Data json.RawMessage
}