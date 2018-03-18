package dto

// PostServerRequest The request body of POST server.
type PostServerRequest struct {
	Address  string `json:"address"`
	Username string `json:"username"`
	Password string `json:"password"`
}
