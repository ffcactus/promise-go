package base

// PostRequestInterface is the interface that PostRequest have.
type PostRequestInterface interface {
	RequestInterface
	ToModel() ModelInterface
}
