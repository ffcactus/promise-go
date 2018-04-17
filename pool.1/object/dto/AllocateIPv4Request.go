package dto

import (
	commonDTO "promise/common/object/dto"
)

// AllocateIPv4Request is the DTO to allocate an IP.
type AllocateIPv4Request struct {
	commonDTO.PromiseRequest
	Key *string `json:"Key"`
}
