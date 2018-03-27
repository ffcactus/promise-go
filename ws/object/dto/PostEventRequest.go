package dto

import (
	"encoding/json"
	"time"
)

// PostEventRequest is the DTO that includes the payload that expected to be send
// from websocket.
type PostEventRequest struct {
	CreatedAt  time.Time       `json:"CreatedAt"`
	Category   string          `json:"Category"`
	Type       string          `json:"Type"`
	ResourceID string          `json:"ResourceID"`
	Data       json.RawMessage `json:"Data"`
}
