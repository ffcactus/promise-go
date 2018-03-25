package service

import (
	log "github.com/sirupsen/logrus"
	commonMessage "promise/common/object/message"
	wsSDK "promise/sdk/ws"
	"promise/server/db"
	"promise/server/object/constError"
	"promise/server/object/dto"
	"promise/server/object/message"
	"promise/server/object/model"
)

// CreateDefaultServerGroup will create the default server group.
func CreateDefaultServerGroup() {
	var request dto.PostServerGroupRequest
	request.Name = "all"
	request.Description = "The default servergroup that includes all the servers."
	dbImpl := db.GetServerGroupDB()
	sg, exist, err := dbImpl.PostServerGroup(request.ToModel())
	if exist {
		log.Debug("The default servergroup exist.")
	}
	if err != nil {
		log.Fatal("Failed to create default servergroup.")
	} else {
		wsSDK.DispatchServerGroupCreate(sg)
		log.Info("Default servergroup created.")
	}
	db.DefaultServerGroupID = sg.ID
}

// PostServerGroup post a server group.
func PostServerGroup(request *dto.PostServerGroupRequest) (*model.ServerGroup, []commonMessage.Message) {
	dbImpl := db.GetServerGroupDB()

	posted, exist, err := dbImpl.PostServerGroup(request.ToModel())
	if exist {
		return nil, []commonMessage.Message{commonMessage.NewResourceDuplicate()}
	}
	if err != nil {
		return nil, []commonMessage.Message{commonMessage.NewInternalError()}
	}
	wsSDK.DispatchServerGroupCreate(posted)
	return posted, nil
}

// GetServerGroup will get server group by ID.
func GetServerGroup(id string) (*model.ServerGroup, []commonMessage.Message) {
	dbImpl := db.GetServerGroupDB()

	sg := dbImpl.GetServerGroup(id)
	if sg == nil {
		return nil, []commonMessage.Message{commonMessage.NewResourceNotExist()}
	}
	return sg, nil
}

// GetServerGroupCollection will get server collection.
func GetServerGroupCollection(start int, count int, filter string) (*model.ServerGroupCollection, []commonMessage.Message) {
	dbImpl := db.GetServerGroupDB()
	ret, err := dbImpl.GetServerGroupCollection(start, count, filter)
	if err != nil {
		return nil, []commonMessage.Message{commonMessage.NewInternalError()}
	}
	return ret, nil
}

// DeleteServerGroup will delete server group by ID.
func DeleteServerGroup(id string) []commonMessage.Message {
	dbImpl := db.GetServerGroupDB()
	previous, err := dbImpl.DeleteServerGroup(id)
	if err != nil && err.Error() == constError.ErrorDeleteDefaultServerGroup.Error() {
		return []commonMessage.Message{message.NewDeleteDefaultServerGroup()}
	}
	if previous == nil {
		return []commonMessage.Message{commonMessage.NewResourceNotExist()}
	}
	if err != nil {
		return []commonMessage.Message{commonMessage.NewInternalError()}
	}
	wsSDK.DispatchServerGroupDelete(previous.ID)
	return nil
}

// DeleteServerGroupCollection will delete all the server group except the default "all".
func DeleteServerGroupCollection() []commonMessage.Message {
	dbImpl := db.GetServerGroupDB()
	err := dbImpl.DeleteServerGroupCollection()
	if err != nil {
		return []commonMessage.Message{commonMessage.NewInternalError()}
	}
	return nil
}
