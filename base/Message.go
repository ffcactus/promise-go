package base

import (
	"net/http"
	"time"
)

const (
	// SeverityNormal is the message severity level.
	SeverityNormal = "Normal"
	// SeverityWarning is the message severity level.
	SeverityWarning = "Warning"
	// SeverityCritical is the message severity level.
	SeverityCritical = "Critical"
)

const (
	// MessageInternalError is message ID.
	MessageInternalError = "Promise.Message.InternalError"
	// MessageNotExist is message ID.
	MessageNotExist = "Promise.Message.NotExist"
	// MessageDuplicate is message ID.
	MessageDuplicate = "Promise.Message.Duplicate"
	// MessageInvalidRequest is message ID.
	MessageInvalidRequest = "Promise.Message.InvalidRequest"
	// MessageTimeout is message ID.
	MessageTimeout = "Promise.Message.Timeout"
	// MessageTransactionError is message ID.
	MessageTransactionError = "Promise.Message.TransactionError"
	// MessageUnknownFilterName is message ID.
	MessageUnknownFilterName = "Promise.Message.UnknownFilterName"

	// MessageTaskNoStep is message ID.
	MessageTaskNoStep = "Task.Message.NoStep"
)

const (
	// SupportNotExist is Support ID.
	SupportNotExist = "Promise.Support.NotExist"
	// SupportDuplicate is Support ID.
	SupportDuplicate = "Promise.Support.Duplicate"
	// SupportInvalidRequest is Support ID.
	SupportInvalidRequest = "Promise.Support.InvalidRequest"
	// SupportInternalError is Support ID.
	SupportInternalError = "Promise.Support.InternalError"
	// SupportTimeout is Support ID.
	SupportTimeout = "Promise.Support.Timeout"
	// SupportTransactionError is Support ID.
	SupportTransactionError = "Promise.Support.TransactionError"
)

// Argument is used to replace a portion in a string. For example, to express
// an error happend on a server, we need replace the {0} in the string below:
// "There is an error on server {0}". The argument here can be:
// { "Type": "URI", "Name": "Server 1", "Value": "/api/v1/server/xxxxxx" }
type Argument struct {
	Type  string `json:"Type"`
	Name  string `json:"Name"`
	Value string `json:"Value"`
}

// Support tells how to solve a problem.
type Support struct {
	ID                string     `json:"ID"` // The unique ID within a micro service.
	Reason            string     `json:"Reason"`
	ReasonArguments   []Argument `json:"ReasonArguments"`
	Solution          string     `json:"Solution"`
	SolutionArguments []Argument `json:"SolutionArguments"`
}

// Message is the message used in Promise project.
type Message struct {
	ID          string     `json:"ID"` // The unique ID within a micro service.
	StatusCode  int        `json:"-"`  // The HTTP status code along with this message.
	Severity    string     `json:"Severity"`
	CreatedAt   time.Time  `json:"CreatedAt"`
	Description string     `json:"Description"`
	Arguments   []Argument `json:"Arguments"` // The arguments in the description.
	Supports    []Support  `json:"Supports"`
}

// NewMessage create a message with default value.
func NewMessage() Message {
	ret := Message{
		CreatedAt:  time.Now(),
		StatusCode: http.StatusBadRequest,
	}
	return ret
}

// NewMessageNotExist return a message that means the resource does not exist.
func NewMessageNotExist() Message {
	ret := NewMessage()
	ret.ID = MessageNotExist
	ret.Severity = SeverityNormal
	ret.Description = "Resource does not exist."
	ret.Supports = []Support{
		NewSupportNotExist(),
	}
	return ret
}

// NewMessageDuplicate return a message that means resource duplication happened.
func NewMessageDuplicate() Message {
	ret := NewMessage()
	ret.ID = MessageDuplicate
	ret.Severity = SeverityNormal
	ret.Description = "Resource duplicated."
	ret.Supports = []Support{
		NewSupportDuplicate(),
	}
	return ret
}

// NewMessageInvalidRequest return a message that means the request is Invalid.
func NewMessageInvalidRequest() Message {
	ret := NewMessage()
	ret.ID = MessageInvalidRequest
	ret.Severity = SeverityNormal
	ret.Description = "The request is invalid."
	ret.Supports = []Support{
		NewSupportInvalidRequest(),
	}
	return ret
}

// NewMessageInternalError return a message that means there is a internal error happened.
func NewMessageInternalError() Message {
	ret := NewMessage()
	ret.ID = MessageInternalError
	ret.Severity = SeverityNormal
	ret.Description = "Internal error happened while process the request."
	ret.Supports = []Support{
		NewSupportInternalError(),
	}
	return ret
}

// NewMessageTimeout return a message that means there is a timeout happend.
func NewMessageTimeout() Message {
	ret := NewMessage()
	ret.ID = MessageTimeout
	ret.Severity = SeverityNormal
	ret.Description = "I/O operation timeout."
	ret.Supports = []Support{
		NewSupportTimeout(),
	}
	return ret
}

// NewMessageTransactionError return a message that means transaction error.
func NewMessageTransactionError() Message {
	ret := NewMessage()
	ret.ID = MessageTransactionError
	ret.Severity = SeverityNormal
	ret.Description = "Transaction error."
	ret.Supports = []Support{
		NewSupportTransactionError(),
	}
	return ret
}

// NewMessageUnknownFilterName returns a new message.
func NewMessageUnknownFilterName() Message {
	ret := NewMessage()
	ret.ID = MessageUnknownFilterName
	ret.Severity = SeverityNormal
	ret.Description = "Unknown filter name."
	ret.Supports = []Support{}
	return ret
}

// NewMessageTaskNoStep returns a new message.
func NewMessageTaskNoStep() Message {
	ret := NewMessage()
	ret.ID = MessageTaskNoStep
	ret.Severity = SeverityNormal
	ret.Description = "No task steps included.."
	ret.Supports = []Support{}
	return ret
}

// NewSupport create a new Support.
func NewSupport() Support {
	ret := Support{}
	ret.ReasonArguments = make([]Argument, 0)
	ret.SolutionArguments = make([]Argument, 0)
	return ret
}

// NewSupportNotExist will return a support message.
func NewSupportNotExist() Support {
	ret := NewSupport()
	ret.ID = SupportNotExist
	ret.Reason = "There is no resource match the identifier."
	ret.Solution = "Verify the identifier and try again."
	return ret
}

// NewSupportDuplicate will return a support message.
func NewSupportDuplicate() Support {
	ret := NewSupport()
	ret.ID = SupportDuplicate
	ret.Reason = "The resource duplication happend inside."
	ret.Solution = "Stop create the duplicated resource."
	return ret
}

// NewSupportInvalidRequest will return a support message.
func NewSupportInvalidRequest() Support {
	ret := NewSupport()
	ret.ID = SupportInvalidRequest
	ret.Reason = "The request is invalid."
	ret.Solution = "Verify the request and try again."
	return ret
}

// NewSupportInternalError will return a support message.
func NewSupportInternalError() Support {
	ret := NewSupport()
	ret.ID = SupportInternalError
	ret.Reason = "Internal error."
	ret.Solution = "Contact support."
	return ret
}

// NewSupportTimeout will return a support message.
func NewSupportTimeout() Support {
	ret := NewSupport()
	ret.ID = SupportTimeout
	ret.Reason = "I/O operation timeout."
	ret.Solution = "Try again later."
	return ret
}

// NewSupportTransactionError will return a support message.
func NewSupportTransactionError() Support {
	ret := NewSupport()
	ret.ID = SupportTransactionError
	ret.Reason = "DB operation failed."
	ret.Solution = "Try again later or contact support."
	return ret
}


