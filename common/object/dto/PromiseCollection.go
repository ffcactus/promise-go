package dto

import (
	"promise/common/object/model"
)

// PromiseCollectionResponse is the Promise collection response DTO.
type PromiseCollectionResponse struct {
	Start       int64            `json:"Start"`
	Count       int64            `json:"Count"`
	Total       int64            `json:"Total"`
	NextPageURI *string          `json:"NextPageURI,omitempty"`
	PrevPageURI *string          `json:"PrevPageURI,omitempty"`
}

// Load will load from model.
func (dto *PromiseCollectionResponse) Load(m model.PromiseCollectionInterface) {
	dto.Start = m.GetStart()
}