package dto

type CollectionElement struct {
	Id string `json:"@odata.id"`
}

type Collection struct {
	Count   int                 `json:"Members@odata.count"`
	Members []CollectionElement `json:"Members"`
}
