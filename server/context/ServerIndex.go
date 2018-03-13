package context

import (
	"fmt"
	"net/http"
	"promise/server/object/dto"
	"promise/server/util"

	log "github.com/sirupsen/logrus"
)

// RootCause RootCause object in ElasticSearch
type RootCause struct {
	Type      string `json:"type"`
	Reason    string `json:"reason"`
	IndexUUID string `json:"index_uuid"`
	Index     string `json:"index"`
}

// ESError object in ElasticSearch
type ESError struct {
	Type      string      `json:"type"`
	RootCause []RootCause `json:"root_cause"`
	Reason    string      `json:"reason"`
	IndexUUID string      `json:"index_uuid"`
	Index     string      `json:"index"`
}

// ESErrorResponse object in ElasticSearch
type ESErrorResponse struct {
	Error  ESError `json:"error"`
	Status int     `json:"status"`
}

// ESIndexResponse object in ElasticSearch
type ESIndexResponse struct {
	ESErrorResponse
	Result  string `json:"result"`
	Created bool   `json:"created"`
	Index   string `json:"_index"`
	Type    string `json:"_type"`
	ID      string `json:"_id"`
	Version int    `json:"_version"`
}

// ESCreateIndexResponse object in ElasticSearch
type ESCreateIndexResponse struct {
	ESErrorResponse
	Acknowledged       bool   `json:"acknowledged"`
	ShardsAcknowledged bool   `json:"shards_acknowledged"`
	Index              string `json:"index"`
}

// ESDeleteIndexResponse object in ElasticSearch
type ESDeleteIndexResponse struct {
	ESErrorResponse
	Acknowledged bool `json:"acknowledged"`
}

// ServerIndexInterface Server index interface.
type ServerIndexInterface interface {
	Init() bool
	Clean() bool
	IndexServer(s *dto.GetServerResponse) error
}

// ServerIndex Server index implementation.
type ServerIndex struct {
	client util.RESTClient
}

// CreateServerIndex Create server index object.
func CreateServerIndex() *ServerIndex {
	return &ServerIndex{
		client: util.RESTClient{
			Client:       &http.Client{},
			Address:      "http://localhost:9200",
			Username:     "",
			Password:     "",
			UseBasicAuth: false,
		},
	}
}

// Init Init server index.
func (index *ServerIndex) Init() bool {
	resp := ESCreateIndexResponse{}
	if _, err := index.client.PutObject("/promise", nil, &resp); err != nil {
		log.Info("Failed to init ElasticSearch, HTTP operation failed error = ", err)
		return false
	}
	if resp.Status == 400 {
		log.Info("ElasticSearch already initialized.")
		return true
	}
	if !resp.Acknowledged {
		log.Info("Failed to init ElasticSearch, result = ", resp)
		return false
	}
	return true
}

// Clean Clean server index.
func (index *ServerIndex) Clean() bool {
	resp := ESDeleteIndexResponse{}
	if _, err := index.client.DeleteObject("/promise", nil, &resp); err != nil {
		log.Info("Failed to clean ElasticSearch, HTTP operation failed error = ", err)
		return false
	}
	if resp.Status == 404 {
		log.Info("ElasticSearch already cleaned.")
		return true
	}
	if !resp.Acknowledged {
		log.Info("Failed to clean ElasticSearch, result = ", resp)
		return false
	}
	return true
}

// IndexServer index a server.
func (index *ServerIndex) IndexServer(s *dto.GetServerResponse) error {
	resp := new(ESIndexResponse)
	if _, err := index.client.PutObject("/promise/server/"+s.ID, s, resp); err != nil {
		log.Warn("Failed to index server to ElasticSearch, server ID = ", s.ID, "HTTP operation faild error = ", err)
		return err
	}
	if resp.Created && resp.Result == "created" {
		log.Info("ElasticSearch put server done, server created, server ID = ", s.ID)
		return nil
	} else if !resp.Created && resp.Result == "updated" {
		log.Info("ElasticSearch put server done, server updated, server ID = ", s.ID)
		return nil
	}
	log.Info("Failed to index server to ElasticSearch, server ID = ", s.ID, ", result = ", resp)
	return fmt.Errorf("failed to index server %s to ElasticSearch, created = %#v, result = %#v", s.ID, resp.Created, resp.Result)

}
