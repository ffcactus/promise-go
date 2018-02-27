package model

import (
	"time"
)

var (
	SeverityNormal   = "Normal"
	SeverityWarning  = "Warning"
	SeverityCritical = "Critical"

	CategoryTask   = "Task"
	CategoryServer = "Server"
	CategoryAuth   = "Auth"
)

// The argument is used to replace a portion in a string. For example, to express
// an error happend on a server, we need replace the {0} in the string below:
// "There is an error on server {0}". The argument here can be:
// { "Type": "URI", "Name": "Server 1", "Value": "/api/v1/server/xxxxxx" }
type Argument struct {
	Type  string
	Name  string
	Value string
}

// The message body.
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

// Message
func NewMessage(category string) Message {
	ret := Message{
		Category:  category,
		CreatedAt: time.Now(),
	}
	return ret
}
