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
	// MessageUnknownPropertyValue is message ID.
	MessageUnknownPropertyValue = "Promise.Message.UnknownPropertyValue"
	// MessageTimeout is message ID.
	MessageTimeout = "Promise.Message.Timeout"
	// MessageTransactionError is message ID.
	MessageTransactionError = "Promise.Message.TransactionError"
	// MessageUnknownFilterName is message ID.
	MessageUnknownFilterName = "Promise.Message.UnknownFilterName"

	// --- For Task ---

	// MessageTaskNoStep is message ID.
	MessageTaskNoStep = "Task.Message.NoStep"

	// --- For IPv4 ---

	// MessageIPv4PoolEmpty is message ID
	MessageIPv4PoolEmpty = "IPv4.Message.IPv4PoolEmpty"
	// MessageIPv4AddressNotExist is message ID.
	MessageIPv4AddressNotExist = "IPv4.Message.AddressNotExist"
	// MessageIPv4FormateError is message ID.
	MessageIPv4FormateError = "IPv4.Message.IPv4FormatError"
	// MessageIPv4RangeEndAddressError is message ID.
	MessageIPv4RangeEndAddressError = "IPv4.Message.IPv4RangeEndAddressError"
	// MessageIPv4RangeSizeError is message ID.
	MessageIPv4RangeSizeError = "IPv4.Message.IPv4RangeSizeError"
	// MessageIPv4RangeCountError is message ID.
	MessageIPv4RangeCountError = "IPv4.Message.IPv4RangeCountError"
	// MessageIPv4NotAllocatedError is message ID.
	MessageIPv4NotAllocatedError = "IPv4.Message.IPv4NotAllocatedError"
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

// NewMessageNotExist returns a message that means the resource does not exist.
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

// NewMessageDuplicate returns a message that means resource duplication happened.
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

// NewMessageInvalidRequest returns a message that means the request is Invalid.
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

// NewMessageUnknownPropertyValue returns a message that means the property in the request have unknown value.
func NewMessageUnknownPropertyValue() Message {
	ret := NewMessage()
	ret.ID = MessageUnknownPropertyValue
	ret.Severity = SeverityNormal
	ret.Description = "Unknown property value."
	ret.Supports = []Support{}
	return ret
}

// NewMessageInternalError returns a message that means there is a internal error happened.
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

// NewMessageTimeout returns a message that means there is a timeout happend.
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

// NewMessageTransactionError returns a message that means transaction error.
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

// --- IPv4 ---

// NewMessageIPv4PoolEmpty will return a message.
func NewMessageIPv4PoolEmpty() Message {
	ret := NewMessage()
	ret.ID = MessageIPv4PoolEmpty
	ret.Severity = SeverityWarning
	ret.Description = "No more IPv4 address can be allocated."
	return ret
}

// NewMessageIPv4AddressNotExist will return a message.
func NewMessageIPv4AddressNotExist() Message {
	ret := NewMessage()
	ret.ID = MessageIPv4AddressNotExist
	ret.Severity = SeverityWarning
	ret.Description = "The address does not exist in this pool."
	return ret
}

// NewMessageIPv4FormatError will return a message.
func NewMessageIPv4FormatError() Message {
	ret := NewMessage()
	ret.ID = MessageIPv4FormateError
	ret.Severity = SeverityWarning
	ret.Description = "Unknown IPv4 format."
	return ret
}

// NewMessageIPv4RangeEndAddressError will return a message.
func NewMessageIPv4RangeEndAddressError() Message {
	ret := NewMessage()
	ret.ID = MessageIPv4RangeEndAddressError
	ret.Severity = SeverityWarning
	ret.Description = "The end address in a range should equal or big then start address"
	return ret
}

// NewMessageIPv4RangeSizeError will return a message.
func NewMessageIPv4RangeSizeError() Message {
	ret := NewMessage()
	ret.ID = MessageIPv4RangeSizeError
	ret.Severity = SeverityWarning
	ret.Description = "The number of addresses in a range should not more than 256."
	return ret
}

// NewMessageIPv4RangeCountError will return a message.
func NewMessageIPv4RangeCountError() Message {
	ret := NewMessage()
	ret.ID = MessageIPv4RangeCountError
	ret.Severity = SeverityWarning
	ret.Description = "IPv4 pool should contain one range at least."
	return ret
}

// NewMessageIPv4NotAllocatedError will return a message.
func NewMessageIPv4NotAllocatedError() Message {
	ret := NewMessage()
	ret.ID = MessageIPv4NotAllocatedError
	ret.Severity = SeverityWarning
	ret.Description = "IP is not allocated before."
	return ret
}

// NewSupport create a new Support.
func NewSupport() Support {
	ret := Support{}
	ret.ReasonArguments = make([]Argument, 0)
	ret.SolutionArguments = make([]Argument, 0)
	return ret
}

// NewSupportNotExist will returns a support message.
func NewSupportNotExist() Support {
	ret := NewSupport()
	ret.ID = SupportNotExist
	ret.Reason = "There is no resource match the identifier."
	ret.Solution = "Verify the identifier and try again."
	return ret
}

// NewSupportDuplicate will returns a support message.
func NewSupportDuplicate() Support {
	ret := NewSupport()
	ret.ID = SupportDuplicate
	ret.Reason = "The resource duplication happend inside."
	ret.Solution = "Stop create the duplicated resource."
	return ret
}

// NewSupportInvalidRequest will returns a support message.
func NewSupportInvalidRequest() Support {
	ret := NewSupport()
	ret.ID = SupportInvalidRequest
	ret.Reason = "The request is invalid."
	ret.Solution = "Verify the request and try again."
	return ret
}

// NewSupportInternalError will returns a support message.
func NewSupportInternalError() Support {
	ret := NewSupport()
	ret.ID = SupportInternalError
	ret.Reason = "Internal error."
	ret.Solution = "Contact support."
	return ret
}

// NewSupportTimeout will returns a support message.
func NewSupportTimeout() Support {
	ret := NewSupport()
	ret.ID = SupportTimeout
	ret.Reason = "I/O operation timeout."
	ret.Solution = "Try again later."
	return ret
}

// NewSupportTransactionError will returns a support message.
func NewSupportTransactionError() Support {
	ret := NewSupport()
	ret.ID = SupportTransactionError
	ret.Reason = "DB operation failed."
	ret.Solution = "Try again later or contact support."
	return ret
}
