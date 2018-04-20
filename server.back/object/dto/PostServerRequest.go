package dto

import (
	commonDTO "promise/common/object/dto"
	commonMessage "promise/common/object/message"
)

// PostServerRequest The request body of POST server.
type PostServerRequest struct {
	commonDTO.PromiseRequest
	Hostname string `json:"Hostname"`
	Username string `json:"Username"`
	Password string `json:"Password"`
}

// Validate the request.
func (dto *PostServerRequest) Validate() *commonMessage.Message {
	return nil
}
