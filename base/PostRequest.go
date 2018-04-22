package base

// PostRequestInterface is the interface that PostRequest have.
type PostRequestInterface interface {
	RequestInterface
	ToModel() ModelInterface
}

// // PostRequest is the DTO of a post request.
// type PostRequest struct {
// }
