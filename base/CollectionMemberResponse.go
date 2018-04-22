package base

// CollectionMemberResponseInterface is the interface that CollectionMemberResponse have
type CollectionMemberResponseInterface interface {
	Load(i CollectionMemberModelInterface) error
}

// CollectionMemberResponse is the a DTO of a member in a get collection response.
type CollectionMemberResponse struct {
	ID       string `json:"ID"`
	URI      string `json:"URI"`
	Category string `json:"Category"`
}

// Load the data from model.
func (dto *CollectionMemberResponse) Load(m *CollectionMemberModel) error {
	dto.ID = m.ID
	dto.Category = m.Category
	dto.URI = CategoryToURI(m.Category, m.ID)
	return nil
}
