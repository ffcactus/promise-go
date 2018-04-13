package base

import (
	"time"
)

// ResponseTemplateInterface is the inteface that a concrete Request should have.
type ResponseTemplateInterface interface {
	GetDebugName() string
	Load(ModelInterface) error
}

// ResponseInterface is the inteface that a Request should have.
type ResponseInterface interface {
	GetDebugName() string
	Load(ModelInterface) error
	GetID() string
	GetCategory() string
}

// Response is the response DTO used in Promise project.
type Response struct {
	TemplateImpl ResponseTemplateInterface `json:"-"`
	ID           string                    `json:"ID"`
	URI          string                    `json:"URI"`
	Category     string                    `json:"Category"`
	CreatedAt    time.Time                 `json:"CreatedAt"`
	UpdatedAt    time.Time                 `json:"UpdatedAt"`
}

// GetDebugName return the name for debug.
func (dto *Response) GetDebugName() string {
	return dto.TemplateImpl.GetDebugName()
}

// GetID returns ID.
func (dto *Response) GetID() string {
	return dto.ID
}

// GetCategory returns Category.
func (dto *Response) GetCategory() string {
	return dto.Category
}

// Load data from model.
func (dto *Response) Load(m ModelInterface) error {
	return dto.TemplateImpl.Load(m)
}

// ResponseLoad load model to DTO.
func ResponseLoad(dto *Response, m *Model) {
	dto.ID = m.ID
	dto.URI = CategoryToURI(m.Category, m.ID)
	dto.Category = m.Category
	dto.CreatedAt = m.CreatedAt
	dto.UpdatedAt = m.UpdatedAt
}

// CollectionMemberResponse is the a DTO in response.
type CollectionMemberResponse struct {
	ID       string `json:"ID"`
	URI      string `json:"URI"`
	Category string `json:"Category"`
}

// Load the data from model.
func (dto *CollectionMemberResponse) Load(m *CollectionMemberModel) {
	dto.ID = m.ID
	dto.Category = m.Category
	dto.URI = CategoryToURI(m.Category, m.ID)
}

// CollectionResponse is the collection response DTO in Promise project.
type CollectionResponse struct {
	Start   int64         `json:"Start"`
	Count   int64         `json:"Count"`
	Total   int64         `json:"Total"`
	Members []interface{} `json:"Members"`
}
