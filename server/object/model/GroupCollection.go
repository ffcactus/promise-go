package model

// GroupMember is the member in collection.
type GroupMember struct {
	URI  string
	ID   string
	Name string
}

// GroupCollection is the model of server group collection.
type GroupCollection struct {
	Start       int
	Count       int
	Total       int
	Members     []GroupMember
	NextPageURI string
	PrevPageURI string
}
