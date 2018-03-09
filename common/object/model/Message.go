package model

import (
	"net/http"
	"time"
)

var (
	// SeverityNormal is the message severity level.
	SeverityNormal = "Normal"
	// SeverityWarning is the message severity level.
	SeverityWarning = "Warning"
	// SeverityCritical is the message severity level.
	SeverityCritical = "Critical"

	// CategoryTask is the message category.
	CategoryTask = "Task"
	// CategoryServer is the message category.
	CategoryServer = "Server"
	// CategoryAuth is the message category.
	CategoryAuth = "Auth"
)

// Argument is used to replace a portion in a string. For example, to express
// an error happend on a server, we need replace the {0} in the string below:
// "There is an error on server {0}". The argument here can be:
// { "Type": "URI", "Name": "Server 1", "Value": "/api/v1/server/xxxxxx" }
type Argument struct {
	Type  string
	Name  string
	Value string
}

// Message is the message body.
type Message struct {
	ID          string // The unique ID within a micro service.
	StatusCode  int    // The HTTP status code along with this message.
	Severity    string
	Category    string // Each micro service should belong one and only one category.
	CreatedAt   time.Time
	Description string
	Arguments   []Argument // The arguments in the description.
	Supports    []Support
}

// NewMessage create a message with default value.
func NewMessage(category string) Message {
	ret := Message{
		Category:   category,
		CreatedAt:  time.Now(),
		StatusCode: http.StatusBadRequest,
	}
	return ret
}
