package model

// IPv4PoolMember is the member in collection.
type IPv4PoolMember struct {
	ID   string
	Name string
}

// IPv4PoolCollection is the model of collection.
type IPv4PoolCollection struct {
	Start       int
	Count       int
	Total       int
	Members     []IPv4PoolMember
	NextPageURI string
	PrevPageURI string
}
