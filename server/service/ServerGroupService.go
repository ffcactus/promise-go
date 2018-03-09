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
