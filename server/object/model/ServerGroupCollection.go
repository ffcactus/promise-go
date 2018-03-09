package model

// ServerGroupMember is the member in collection.
type ServerGroupMember struct {
	URI  string
	ID   string
	Name string
}

// ServerGroupCollection is the model of server group collection.
type ServerGroupCollection struct {
	Start       int
	Count       int
	Total       int
	Members     []ServerGroupMember
	NextPageURI string
	PrevPageURI string
}
