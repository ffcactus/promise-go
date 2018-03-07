package dto

import (
	"promise/server/object/model"
)

type ServerMember struct {
	PageURI string `json:"URI"`
	Name    string `json:"Name"`
	State   string `json:"State"`
	Health  string `json:"Health"`
}

type GetServerCollectionResponse struct {
	Start       int            `json:"Start"`
	Count       int            `json:"Count"`
	Total       int            `json:"Total"`
	Members     []ServerMember `json:"Members"`
	NextPageURI *string        `json:"NextPageURI,omitempty"`
	PrevPageURI *string        `json:"PrevPageURI,omitempty"`
}

func (this *GetServerCollectionResponse) Load(m *model.ServerCollection) {
	this.Start = m.Start
	this.Count = m.Count
	this.Total = m.Total
	this.Members = make([]ServerMember, 0)
	for i, _ := range m.Members {
		this.Members = append(this.Members, ServerMember{
			PageURI: m.Members[i].URI,
			Name:    m.Members[i].Name,
			State:   m.Members[i].State,
			Health:  m.Members[i].Health,
		})
	}
	if m.NextPageURI != "" {
		this.NextPageURI = &m.NextPageURI
	}
	if m.PrevPageURI != "" {
		this.PrevPageURI = &m.PrevPageURI
	}
}
