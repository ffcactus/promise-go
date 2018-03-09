package service

import (
	commonM "promise/common/object/model"
	"promise/server/db"
	"promise/server/object/dto"
	"promise/server/object/message"
	"promise/server/object/model"
)

// PostServerGroup post a server group.
func PostServerGroup(request *dto.PostServerGroupRequest) (*model.ServerGroup, []commonM.Message) {
	dbImpl := db.GetServerGroupDB()

	posted, exist, err := dbImpl.PostServerGroup(request.ToModel())
	if exist {
		return nil, []commonM.Message{message.NewServerGroupExist()}
	}
	if err != nil {
		return nil, []commonM.Message{message.NewServerInternalError()}
	}
	return posted, nil
}

// GetServerGroup will get server group by ID.
func GetServerGroup(id string) (*model.ServerGroup, []commonM.Message) {
	dbImpl := db.GetServerGroupDB()

	sg := dbImpl.GetServerGroup(id)
	if sg == nil {
		return nil, []commonM.Message{message.NewServerGroupNotExist()}
	}
	return sg, nil
}

// GetServerGroupCollection will get server collection.
func GetServerGroupCollection(start int, count int) (*model.ServerGroupCollection, []commonM.Message) {
	dbImpl := db.GetServerGroupDB()
	ret, err := dbImpl.GetServerGroupCollection(start, count)
	if err != nil {
		return nil, []commonM.Message{message.NewServerInternalError()}
	}
	return ret, nil
}

// DeleteServerGroup will delete server group by ID.
func DeleteServerGroup(id string) []commonM.Message {
	dbImpl := db.GetServerGroupDB()
	exist, err := dbImpl.DeleteServerGroup(id)
	if err != nil {
		return []commonM.Message{message.NewServerInternalError()}
	}
	if !exist {
		return []commonM.Message{message.NewServerGroupNotExist()}
	}
	return nil
}
