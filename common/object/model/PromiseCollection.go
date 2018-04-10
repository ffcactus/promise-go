package model

// PromiseCollectionInterface is the Promise collection interface.
type PromiseCollectionInterface interface {
	GetStart() int64
	GetCount() int64
	GetTotal() int64
	GetMembers() []PromiseMemberInterface
}

// PromiseCollection is the resource collection.
type PromiseCollection struct {
	Start   int64
	Count   int64
	Total   int64
	Members []PromiseMemberInterface
}

// GetStart return the start.
func (m *PromiseCollection) GetStart() int64 {
	return m.Start
}

// GetCount return the count.
func (m *PromiseCollection) GetCount() int64 {
	return m.Count
}

// GetTotal return the total.
func (m *PromiseCollection) GetTotal() int64 {
	return m.Total
}

// GetMembers return the members.
func (m *PromiseCollection) GetMembers() []PromiseMemberInterface {
	return m.Members
}
