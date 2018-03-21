package service

import (
	commomMessage "promise/common/object/message"
	"promise/server/db"
	"promise/server/object/dto"
	"promise/server/object/message"
	"promise/server/object/model"
)

// PostServerServerGroup post a server-group.
func PostServerServerGroup(request *dto.PostServerServerGroupRequest) (*model.ServerServerGroup, []commomMessage.Message) {
	dbImpl := db.GetServerServerGroupInstance()

	posted, exist, err := dbImpl.PostServerServerGroup(request.ToModel())
	if exist {
		return nil, []commomMessage.Message{message.NewServerGroupExist()}
	}
	if err != nil {
		return nil, []commomMessage.Message{message.NewServerInternalError()}
	}
	return posted, nil
}

// GetServerServerGroup will get server group by ID.
func GetServerServerGroup(id string) (*model.ServerServerGroup, []commomMessage.Message) {
	dbImpl := db.GetServerServerGroupInstance()

	ssg := dbImpl.GetServerServerGroup(id)
	if ssg == nil {
		return nil, []commomMessage.Message{message.NewServerGroupNotExist()}
	}
	return ssg, nil
}

// GetServerServerGroupCollection will get server collection.
func GetServerServerGroupCollection(start int, count int, filter string) (*model.ServerServerGroupCollection, []commomMessage.Message) {
	dbImpl := db.GetServerServerGroupInstance()
	ret, err := dbImpl.GetServerServerGroupCollection(start, count, filter)
	if err != nil {
		return nil, []commomMessage.Message{message.NewServerInternalError()}
	}
	return ret, nil
}

// DeleteServerServerGroup will delete server group by ID.
func DeleteServerServerGroup(id string) []commomMessage.Message {
	return nil
}

// DeleteServerServerGroupCollection will delete all the server group except the default "all".
func DeleteServerServerGroupCollection() []commomMessage.Message {
	dbImpl := db.GetServerServerGroupInstance()
	err := dbImpl.DeleteServerServerGroupCollection()
	if err != nil {
		return []commomMessage.Message{message.NewServerInternalError()}
	}
	return nil
}
