package service

import (
	commonM "promise/common/object/model"
	"promise/server/db"
	"promise/server/object/dto"
	"promise/server/object/model"
)

// PostServerGroup post a server group.
func PostServerGroup(request *dto.PostServerGroupRequest) (*model.ServerGroup, []commonM.Message) {
	dbImpl := db.GetServerGroupDB()

	posted, exist, err := dbImpl.PostServerGroup(request.ToModel())
	if exist {
		return nil, []commonM.Message{model.NewServerGroupExist()}
	}
	if err != nil {
		return nil, []commonM.Message{model.NewInternalError()}
	}
	return posted, nil
}
