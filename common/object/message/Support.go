package message

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

// Support tells how to solve a problem.
type Support struct {
	ID                string // The unique ID within a micro service.
	Reason            string
	ReasonArguments   []Argument
	Solution          string
	SolutionArguments []Argument
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
