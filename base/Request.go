package base

// RequestInterface is the interface that  Request should have.
type RequestInterface interface {
	NewInstance() RequestInterface // Create a new instance of the request.
	IsValid() *ErrorResponse       // Check if the request is valid.
}
