package base

// ClientError represents the error in client operation.
type ClientError interface {
	Status() int
	Body() []byte
	ConnectionError() bool
	Timeout() bool
	LoginFailure() bool
	String() string
}

// NewErrorResponseFromClientError creates error response depends on the client error.
// Return internal error if no suitable one.
func NewErrorResponseFromClientError(clientError ClientError) *ErrorResponse {
	switch {
	case clientError.LoginFailure():
		return NewErrorResponseDeviceCredential()
	case clientError.ConnectionError():
		return NewErrorResponseDeviceConnection()
	case clientError.Timeout():
		return NewErrorResponseDeviceTimeout()
	default:
		return NewErrorResponseInternalError()
	}
}