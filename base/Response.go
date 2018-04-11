package base

import (
	"promise/apps"
	"time"
)

// ResponseTemplateInterface is the inteface that a concrete Request should have.
type ResponseTemplateInterface interface {
	GetDebugName() string
	Load(ModelInterface) error
	GetID() string
	GetURI() string
	GetCategory() string
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
}

// ResponseInterface is the inteface that a Request should have.
type ResponseInterface interface {
	GetDebugName() string
	Load(ModelInterface) error
	GetID() string
	GetURI() string
	GetCategory() string
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
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

// GetURI returns URI.
func (dto *Response) GetURI() string {
	return dto.URI
}

// GetCategory returns Category.
func (dto *Response) GetCategory() string {
	return dto.Category
}

// GetCreatedAt returns CreatedAt.
func (dto *Response) GetCreatedAt() time.Time {
	return dto.CreatedAt
}

// GetUpdatedAt returns UpdatedAt.
func (dto *Response) GetUpdatedAt() time.Time {
	return dto.UpdatedAt
}

// Load data from model.
func (dto *Response) Load(m ModelInterface) error {
	return dto.TemplateImpl.Load(m)
}

// ResponseLoad load model to DTO.
func ResponseLoad(dto *Response, m *Model) {
	dto.ID = m.ID
	dto.URI = apps.CategoryToURI(m.Category, m.ID)
	dto.Category = m.Category
	dto.CreatedAt = m.CreatedAt
	dto.UpdatedAt = m.UpdatedAt
}

// MemberResponseTemplateInterface is the interface that a concrete one should have.
type MemberResponseTemplateInterface interface {
	Load(interface{}) error
}

// MemberResponseInterface is the interface that a member response should have.
type MemberResponseInterface interface {
	Load(interface{}) error
}

// MemberResponse is the a DTO in response.
type MemberResponse struct {
	TemplateImpl MemberResponseTemplateInterface `json:"-"`
	ID           string                    `json:"ID"`
	URI          string                    `json:"URI"`
	Category     string                    `json:"Category"`
}

// Load the data from model.
func (dto *MemberResponse) Load(i MemberModelInterface) error {
	return dto.TemplateImpl.Load(i)
}

// CollectionResponseTemplateInterface is the interface that 
// a concrete collection response should have.
type CollectionResponseTemplateInterface interface {
	Load([]MemberModelInterface) error
	NewMemberResponse() MemberResponseInterface
}

// CollectionResponseInterface is the interface that a collection
// response should have.
// type CollectionResponseInterface interface {
// 	Load([]MemberModelInterface) error
// 	SetStart(int64)
// 	SetCount(int64)
// 	SetTotal(int64)
// }

// CollectionResponse is the collection response DTO in Promise project.
type CollectionResponse struct {
	TemplateImpl CollectionResponseTemplateInterface
	Start       int64              `json:"Start"`
	Count       int64              `json:"Count"`
	Total       int64              `json:"Total"`
	Members		[]MemberResponseInterface `json:"Members"`
}

// Load data from model.
func (dto *CollectionResponse) Load(m *CollectionModel) error {
	dto.Start = m.Start
	dto.Count = m.Count
	dto.Total = m.Total
	for _, v := range m.Members {
		vv := dto.TemplateImpl.NewMemberResponse()
		vv.Load(v)
		dto.Members = append(dto.Members, vv)
	}
	return nil
}