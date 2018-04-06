package model

// IPv4PoolMember is the member in collection.
type IPv4PoolMember struct {
	ID   string
	Name string
}

// IPv4PoolCollection is the model of collection.
type IPv4PoolCollection struct {
	Start       int64
	Count       int64
	Total       int64
	Members     []IPv4PoolMember
	NextPageURI string
	PrevPageURI string
}
