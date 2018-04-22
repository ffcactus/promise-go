package base

// UpdateRequestInterface is the interface that UpdateRequest have.
type UpdateRequestInterface interface {
	RequestInterface
	UpdateModel(currrent ModelInterface) error
}

// // UpdateRequest is the DTO of a post request.
// type UpdateRequest struct {
// 	Request
// }
