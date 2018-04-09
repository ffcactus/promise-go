package base

// CollectionInterface is the interface a Collection should have.
type CollectionInterface interface {
	GetStart() int64
	GetCount() int64
	GetTotal() int64
	GetMembers() []MemberInterface
}

// Collection is the collection used in Promise project.
type Collection struct {
	Start int64
	Count int64
	Total int64
	Members []MemberInterface
}

// GetStart return the start.
func (m *Collection) GetStart() int64 {
	return m.Start
}

// GetCount return the count.
func (m *Collection) GetCount() int64 {
	return m.Count
}

// GetTotal return the total.
func (m *Collection) GetTotal() int64 {
	return m.Total
}

// GetMembers return the members.
func (m *Collection) GetMembers() []MemberInterface {
	return m.Members
}
