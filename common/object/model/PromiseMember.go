package model

// PromiseMemberInterface is the promise collection member interface.
type PromiseMemberInterface interface {
	GetID() string
	GetCategory() string
}

// PromiseMember is the member in resource collection.
type PromiseMember struct {
	ID       string
	Category string
}

// GetID return the ID.
func (m *PromiseMember) GetID() string {
	return m.ID
}

// GetCategory return the category.
func (m *PromiseMember) GetCategory() string {
	return m.Category
}
