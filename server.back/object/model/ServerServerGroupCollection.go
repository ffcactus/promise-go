package model

// ServerServerGroupMember is the member in collection.
type ServerServerGroupMember struct {
	ID            string
	ServerID      string
	ServerGroupID string
}

// ServerServerGroupCollection is the model of servergroup collection.
type ServerServerGroupCollection struct {
	Start       int64
	Count       int64
	Total       int64
	Members     []ServerServerGroupMember
	NextPageURI string
	PrevPageURI string
}
