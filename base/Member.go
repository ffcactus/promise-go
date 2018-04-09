package base

// MemberInterface is the interface a Member should have
type MemberInterface interface {
	GetID() string
	GetCategory() string
}

// Member is the collection member used in Promise project.
type Member struct {
	ID string
	Category string
}

// GetID return the ID.
func (m *Member) GetID() string {
	return m.ID
}

// GetCategory return the category.
func (m *Member) GetCategory() string {
	return m.Category
}
