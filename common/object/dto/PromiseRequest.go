package dto

// PromiseRequest represents the common request DTO used in Promise project.
type PromiseRequest struct {
}

// PromiseRequestInterface is the interface.
type PromiseRequestInterface interface {
	Validate() error
}

// Validate the request.
func (dto *PromiseRequest) Validate() error {
	return nil
}
