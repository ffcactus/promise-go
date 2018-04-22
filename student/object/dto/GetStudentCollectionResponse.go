package dto

import (
	"promise/base"
	"promise/student/object/model"
)

// GetStudentCollectionResponse is the collection DTO.
type GetStudentCollectionResponse struct {
	base.CollectionResponse
}

// Load will load info from model.
func (dto *GetStudentCollectionResponse) Load(m *base.CollectionModel) error {
	dto.Start = m.Start
	dto.Count = m.Count
	dto.Total = m.Total
	dto.Members = []base.CollectionMemberResponseInterface{}
	for _, v := range m.Members {
		member := StudentCollectionMember{}
		if err := member.Load(v); err != nil {
			return err
		}
		dto.Members = append(dto.Members, &member)
	}
	return nil
}

// StudentCollectionMember is the DTO used in collection response.
type StudentCollectionMember struct {
	base.CollectionMemberResponse
	Name string `json:"Name"`
}

// Load will load info from model.
func (dto *StudentCollectionMember) Load(i base.CollectionMemberModelInterface) error {
	m, ok := i.(*model.StudentCollectionMember)
	if !ok {
		return base.ErrorDataConvert
	}
	dto.CollectionMemberResponse.Load(&m.CollectionMemberModel)
	dto.Name = m.Name
	return nil
}
