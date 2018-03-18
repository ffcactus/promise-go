package model

// ServerMember Server member in collection.
type ServerMember struct {
	ID     string
	Name   string
	State  string
	Health string
}

// ServerCollection server collection.
type ServerCollection struct {
	Start       int
	Count       int
	Total       int
	Members     []ServerMember
	NextPageURI string
	PrevPageURI string
}
