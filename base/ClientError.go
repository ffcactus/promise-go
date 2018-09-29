package base

// ClientError represents the error in client operation.
type ClientError interface {
	error
	Status() int
	Body() []byte
	ConnectionError() bool
	Timeout() bool
	LoginFailure() bool
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
