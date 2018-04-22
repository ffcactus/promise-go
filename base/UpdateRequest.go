package base

// UpdateRequestInterface is the interface that UpdateRequest have.
type UpdateRequestInterface interface {
	RequestInterface
	UpdateModel(currrent ModelInterface) error
}
