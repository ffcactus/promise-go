package base

import (
	"time"
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

// Support tells how to solve a problem.
type Support struct {
	ID                string // The unique ID within a micro service.
	Reason            string
	ReasonArguments   []Argument
	Solution          string
	SolutionArguments []Argument
}

// Message is the message used in Promise project.
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

// NewMessageInvalidRequest return a new message.
func NewMessageInvalidRequest() Message {
	return Message{}
}

// NewMessageResourceDuplicate return a new message.
func NewMessageResourceDuplicate() Message {
	return Message{}
}

// NewMessageTransactionError return a new message.
func NewMessageTransactionError() Message {
	return Message{}
}

// NewMessageNotExist return a new message.
func NewMessageNotExist() Message {
	return Message{}
}
