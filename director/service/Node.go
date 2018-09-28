package service

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/director/object/dto"
	"promise/director/object/model"
)

var (
	_client *client.Client
)

// The init of service package.
func init() {
	var err error
	if _client, err = client.NewClientWithOpts(
		// TODO do not hard write the default gwbridge IP.
		client.WithHost("http://172.18.0.1:2376"),
		client.WithVersion("1.35"),
	); err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("Service failed to create client.")
	}
}

// Node is the service
type Node struct {
	base.CRUDService
}

// Category returns the category of this service.
func (s *Node) Category() string {
	return base.CategoryNode
}

// Response creates a new response DTO.
func (s *Node) Response() base.GetResponseInterface {
	return new(dto.GetNodeResponse)
}

// DB returns the DB implementation.
// DB is not need.
func (s *Node) DB() base.DBInterface {
	return nil
}

// GetCollection get the Node collection.
func (s *Node) GetCollection(start int64, count int64, filter string) (*base.CollectionModel, []base.ErrorResponse) {
	if _client == nil {
		return nil, []base.ErrorResponse{*base.NewErrorResponseInternalError()}
	}
	if start != 0 || count != -1 || filter != "" {
		return nil, []base.ErrorResponse{*base.NewErrorResponseInvalidRequest()}
	}
	nodes, err := _client.NodeList(context.Background(), types.NodeListOptions{})
	if err != nil {
		panic(err)
	}
	collection := base.CollectionModel{
		Start: 0,
		Count: int64(len(nodes)),
		Total: int64(len(nodes)),
	}
	for _, node := range nodes {
		n := model.NodeCollectionMember{}
		n.ID = node.ID
		n.Category = base.CategoryNode
		n.Hostname = node.Description.Hostname
		n.Status = string(node.Status.State)
		n.Availibility = string(node.Spec.Availability)
		if node.ManagerStatus.Leader {
			n.ManagerStatus = "Leader"
		} else {
			n.ManagerStatus = string(node.ManagerStatus.Reachability)
		}
		collection.Members = append(collection.Members, &n)
	}
	return &collection, nil
}
