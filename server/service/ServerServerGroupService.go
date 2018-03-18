package service

import (
	commonM "promise/common/object/model"
	"promise/server/object/dto"
	"promise/server/object/model"
)

// PostServerServerGroup post a server-group.
func PostServerServerGroup(request *dto.PostServerServerGroupRequest) (*model.ServerServerGroup, []commonM.Message) {
	return nil, nil
}

// GetServerServerGroup will get server group by ID.
func GetServerServerGroup(id string) (*model.ServerServerGroup, []commonM.Message) {
	return nil, nil
}

// GetServerServerGroupCollection will get server collection.
func GetServerServerGroupCollection(start int, count int) (*model.ServerServerGroupCollection, []commonM.Message) {
	return nil, nil
}

// DeleteServerServerGroup will delete server group by ID.
func DeleteServerServerGroup(id string) []commonM.Message {
	return nil
}

// DeleteServerServerGroupCollection will delete all the server group except the default "all".
func DeleteServerServerGroupCollection() []commonM.Message {
	return nil
}
