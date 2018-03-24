package dto

import (
	"encoding/json"
	"time"
)

// PostEventRequest is the DTO that includes the payload that expected to be send
// from websocket.
type PostEventRequest struct {
	CreatedAt  time.Time
	Category   string
	Type       string
	ResourceID string
	Data       json.RawMessage
}
