package base

// ContextInterface is the interface that a Context should have.
type ContextInterface interface {
	Request() RequestInterface
}
