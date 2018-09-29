package base

// ServiceError is the imterface that an error happend in the REST call between service should support.
type ServiceError interface {
	Category() string
	Status() int
	Timeout() bool
	ErrorResponse() ErrorResponse
}
