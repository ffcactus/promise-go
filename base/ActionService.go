package base

// ActionServiceInterface is the interface that an action service should have.
type ActionServiceInterface interface {
	ServiceInterface
	Perform(id string, request ActionRequestInterface) (ResponseInterface, []ErrorResponse)
}
