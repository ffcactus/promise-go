package model

// PromiseCollection is the resource collection.
type PromiseCollection struct {
	Start int64
	Count int64
	Total int64
	Members []PromiseCollectionMemberInterface
	NextPageURI string
	PrevPageURI string
}