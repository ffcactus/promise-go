package dto

// PostServerRequest The request body of POST server.
type PostServerRequest struct {
	Hostname string `json:"Hostname"`
	Username string `json:"Username"`
	Password string `json:"Password"`
}
