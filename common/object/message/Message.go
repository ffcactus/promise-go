package message

import (
	"net/http"
	"promise/common/object/constValue"
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

const (
	// MessageInternalError is message ID.
	MessageInternalError = "Promise.Message.InternalError"
	// MessageResourceNotExist is message ID.
	MessageResourceNotExist = "Promise.Message.ResourceNotExist"
	// MessageResourceDuplicate is message ID.
	MessageResourceDuplicate = "Promise.Message.ResourceDuplicate"
	// MessageInvalidRequest is message ID.
	MessageInvalidRequest = "Promise.Message.InvalidRequest"
	// MessageTimeout is message ID.
	MessageTimeout = "Promise.Message.Timeout"
)

// NewMessage create a message with default value.
func NewMessage(category string) Message {
	ret := Message{
		Category:   category,
		CreatedAt:  time.Now(),
		StatusCode: http.StatusBadRequest,
	}
	return ret
}

// NewResourceNotExist return a message that means the resource does not exist.
func NewResourceNotExist() Message {
	ret := NewMessage(constValue.CategoryPromise)
	ret.ID = MessageResourceNotExist
	ret.Severity = constValue.SeverityNormal
	ret.Description = "Resource does not exist."
	ret.Supports = []Support{
		NewSupportResourceNotExist(),
	}
	return ret
}

// NewResourceDuplicate return a message that means resource duplication happened.
func NewResourceDuplicate() Message {
	ret := NewMessage(constValue.CategoryPromise)
	ret.ID = MessageResourceDuplicate
	ret.Severity = constValue.SeverityNormal
	ret.Description = "Resource duplicated."
	ret.Supports = []Support{
		NewSupportResourceDuplicate(),
	}
	return ret
}

// NewBadRequest return a message that means the request is Invalid.
func NewInvalidRequest() Message {
	ret := NewMessage(constValue.CategoryPromise)
	ret.ID = MessageInvalidRequest
	ret.Severity = constValue.SeverityNormal
	ret.Description = "The request is invalid."
	ret.Supports = []Support{
		NewSupportInvalidRequest(),
	}
	return ret
}

// NewInternalError return a message that means there is a internal error happened.
func NewInternalError() Message {
	ret := NewMessage(constValue.CategoryPromise)
	ret.ID = MessageInternalError
	ret.Severity = constValue.SeverityNormal
	ret.Description = "Internal error happened while process the request."
	ret.Supports = []Support{
		NewSupportInternalError(),
	}
	return ret
}

// NewTimeout return a message that means there is a timeout happend.
func NewTimeout() Message {
	ret := NewMessage(constValue.CategoryPromise)
	ret.ID = MessageTimeout
	ret.Severity = constValue.SeverityNormal
	ret.Description = "I/O operation timeout."
	ret.Supports = []Support{
		NewSupportTimeout(),
	}
	return ret
}
