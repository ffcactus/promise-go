package service

import (
	commonM "promise/common/object/model"
	"promise/server/object/dto"
	"promise/server/object/message"
	"promise/server/object/model"
	"promise/server/db"
)

// PostServerServerGroup post a server-group.
func PostServerServerGroup(request *dto.PostServerServerGroupRequest) (*model.ServerServerGroup, []commonM.Message) {
	dbImpl := db.GetServerServerGroupInstance()

	posted, exist, err := dbImpl.PostServerServerGroup(request.ToModel())
	if exist {
		return nil, []commonM.Message{message.NewServerGroupExist()}
	}
	if err != nil {
		return nil, []commonM.Message{message.NewServerInternalError()}
	}
	return posted, nil
}

// GetServerServerGroup will get server group by ID.
func GetServerServerGroup(id string) (*model.ServerServerGroup, []commonM.Message) {
	dbImpl := db.GetServerServerGroupInstance()

	ssg := dbImpl.GetServerServerGroup(id)
	if ssg == nil {
		return nil, []commonM.Message{message.NewServerGroupNotExist()}
	}
	return ssg, nil
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
	dbImpl := db.GetServerServerGroupInstance()
	err := dbImpl.DeleteServerServerGroupCollection()
	if err != nil {
		return []commonM.Message{message.NewServerInternalError()}
	}
	return nil
}
