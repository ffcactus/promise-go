package service

import (
	commonM "promise/common/object/model"
	"promise/server/db"
	"promise/server/object/dto"
	"promise/server/object/model"
	"promise/server/object/message"
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
