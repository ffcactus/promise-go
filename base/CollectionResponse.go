package base

// CollectionResponse is the DTO of a get collection response.
type CollectionResponse struct {
	Start   int64                               `json:"Start"`
	Count   int64                               `json:"Count"`
	Total   int64                               `json:"Total"`
	Members []CollectionMemberResponseInterface `json:"Members"`
}
