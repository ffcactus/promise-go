package promise

// RequestInterface is the interface that a PromiseRequest should have.
type RequestInterface interface {
	GetDebugName() string
}

// Request is the request DTO used in Promise project.
type Request struct {
}

// GetDebugName return the name for debug.
func (dto *Request) GetDebugName() string {
	return "NotProvided"
}
