package model

// PromiseCollectionMemberInterface is the promise collection member interface.
type PromiseCollectionMemberInterface interface {
	GetID() string
	GetCategory() string
}

// PromiseCollectionMember is the member in resource collection.
type PromiseCollectionMember struct {
	ID        string
	Category  string
}

// GetID return the ID.
func (m *PromiseCollectionMember) GetID() string {
	return m.ID
}

// GetCategory return the category.
func (m *PromiseCollectionMember) GetCategory() string {
	return m.Category
}
