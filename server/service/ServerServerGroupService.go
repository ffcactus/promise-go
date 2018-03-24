package service

import (
	commonMessage "promise/common/object/message"
	"promise/server/db"
	"promise/server/object/constError"
	"promise/server/object/dto"
	"promise/server/object/message"
	"promise/server/object/model"
)

// PostServerServerGroup post a server-group.
func PostServerServerGroup(request *dto.PostServerServerGroupRequest) (*model.ServerServerGroup, []commonMessage.Message) {
	dbImpl := db.GetServerServerGroupInstance()

	posted, exist, err := dbImpl.PostServerServerGroup(request.ToModel())
	if exist {
		return nil, []commonMessage.Message{commonMessage.NewResourceDuplicate()}
	}
	if err != nil {
		return nil, []commonMessage.Message{commonMessage.NewInternalError()}
	}
	return posted, nil
}

// GetServerServerGroup will get server group by ID.
func GetServerServerGroup(id string) (*model.ServerServerGroup, []commonMessage.Message) {
	dbImpl := db.GetServerServerGroupInstance()

	ssg := dbImpl.GetServerServerGroup(id)
	if ssg == nil {
		return nil, []commonMessage.Message{commonMessage.NewResourceNotExist()}
	}
	return ssg, nil
}

// GetServerServerGroupCollection will get server collection.
func GetServerServerGroupCollection(start int, count int, filter string) (*model.ServerServerGroupCollection, []commonMessage.Message) {
	dbImpl := db.GetServerServerGroupInstance()
	ret, err := dbImpl.GetServerServerGroupCollection(start, count, filter)
	if err != nil {
		return nil, []commonMessage.Message{commonMessage.NewInternalError()}
	}
	return ret, nil
}

// DeleteServerServerGroup will delete server group by ID.
func DeleteServerServerGroup(id string) []commonMessage.Message {
	dbImpl := db.GetServerServerGroupInstance()
	previous, err := dbImpl.DeleteServerServerGroup(id)
	if err != nil && err.Error() == constError.ErrorDeleteDefaultServerServerGroup.Error() {
		return []commonMessage.Message{message.NewDeleteDefaultServerServerGroup()}
	}
	if previous == nil {
		return []commonMessage.Message{commonMessage.NewResourceNotExist()}
	}
	if err != nil {
		return []commonMessage.Message{commonMessage.NewInternalError()}
	}

	return nil
}

// DeleteServerServerGroupCollection will delete all the server group except the default "all".
func DeleteServerServerGroupCollection() []commonMessage.Message {
	dbImpl := db.GetServerServerGroupInstance()
	err := dbImpl.DeleteServerServerGroupCollection()
	if err != nil {
		return []commonMessage.Message{commonMessage.NewInternalError()}
	}
	return nil
}
