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

// MessageInterface is the interface that a Promise message should have.
type MessageInterface interface {
	GetID() string
	GetStatusCode() int
	GetSeverity() string
	GetCategory() string
	GetCreatedAt() time.Time
	GetDescription() string
	GetArguments() []Argument
	GetSupports() []Support
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

// GetID return the ID of this message.
func (o *Message) GetID() string {
	return o.ID
}

// GetStatusCode return the StatusCode of this message.
func (o *Message) GetStatusCode() int {
	return o.StatusCode
}

// GetSeverity return the Severity of this message.
func (o *Message) GetSeverity() string {
	return o.Severity
}

// GetCategory return the Category of this message.
func (o *Message) GetCategory() string {
	return o.Category
}

// GetCreatedAt return the CreatedAt of this message.
func (o *Message) GetCreatedAt() time.Time {
	return o.CreatedAt
}

// GetDescription return the Description of this message.
func (o *Message) GetDescription() string {
	return o.Description
}

// GetArguments return the Arguments of this message.
func (o *Message) GetArguments() []Argument {
	return o.Arguments
}

// GetSupports return the Supports of this message.
func (o *Message) GetSupports() []Support {
	return o.Supports
}

// NewMessageInvalidRequest return a new message.
func NewMessageInvalidRequest() MessageInterface {
	return new(Message)
}

// NewMessageResourceDuplicate return a new message.
func NewMessageResourceDuplicate() MessageInterface {
	return new(Message)
}

// NewMessageTransactionError return a new message.
func NewMessageTransactionError() MessageInterface {
	return new(Message)
}
