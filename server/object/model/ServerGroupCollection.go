package model

// ServerGroupMember is the member in collection.
type ServerGroupMember struct {
	ID   string
	Name string
}

// ServerGroupCollection is the model of servergroup collection.
type ServerGroupCollection struct {
	Start       int64
	Count       int64
	Total       int64
	Members     []ServerGroupMember
	NextPageURI string
	PrevPageURI string
}
