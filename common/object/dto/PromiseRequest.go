package dto

// PromiseRequest represents the common request DTO used in Promise project.
type PromiseRequest struct {
}

// PromiseRequestInterface is the interface.
type PromiseRequestInterface interface {
	IsValid() bool
}

// IsValid check if the request is valid.
func (dto *PromiseRequest) IsValid() bool {
	return true
}
