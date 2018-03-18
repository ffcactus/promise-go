package service

import (
	log "github.com/sirupsen/logrus"
	commonM "promise/common/object/model"
	"promise/server/db"
	"promise/server/object/dto"
	"promise/server/object/message"
	"promise/server/object/model"
)

// CreateDefaultGroup will create the default server group.
func CreateDefaultGroup() {
	var request dto.PostGroupRequest
	request.Name = "all"
	request.Description = "The default server group that includes all the servers."
	dbImpl := db.GetGroupDB()
	_, exist, err := dbImpl.PostGroup(request.ToModel())
	if exist {
		log.Debug("The default server group exist.")
	}
	if err != nil {
		log.Fatal("Failed to create default server group.")
	} else {
		log.Info("Default server group created.")
	}

}

// PostGroup post a server group.
func PostGroup(request *dto.PostGroupRequest) (*model.Group, []commonM.Message) {
	dbImpl := db.GetGroupDB()

	posted, exist, err := dbImpl.PostGroup(request.ToModel())
	if exist {
		return nil, []commonM.Message{message.NewGroupExist()}
	}
	if err != nil {
		return nil, []commonM.Message{message.NewServerInternalError()}
	}
	return posted, nil
}

// GetGroup will get server group by ID.
func GetGroup(id string) (*model.Group, []commonM.Message) {
	dbImpl := db.GetGroupDB()

	sg := dbImpl.GetGroup(id)
	if sg == nil {
		return nil, []commonM.Message{message.NewGroupNotExist()}
	}
	return sg, nil
}

// GetGroupCollection will get server collection.
func GetGroupCollection(start int, count int) (*model.GroupCollection, []commonM.Message) {
	dbImpl := db.GetGroupDB()
	ret, err := dbImpl.GetGroupCollection(start, count)
	if err != nil {
		return nil, []commonM.Message{message.NewServerInternalError()}
	}
	return ret, nil
}

// DeleteGroup will delete server group by ID.
func DeleteGroup(id string) []commonM.Message {
	dbImpl := db.GetGroupDB()
	exist, err := dbImpl.DeleteGroup(id)
	if err != nil {
		return []commonM.Message{message.NewServerInternalError()}
	}
	if !exist {
		return []commonM.Message{message.NewGroupNotExist()}
	}
	return nil
}

// DeleteGroupCollection will delete all the server group except the default "all".
func DeleteGroupCollection() []commonM.Message {
	dbImpl := db.GetGroupDB()
	err := dbImpl.DeleteGroupCollection()
	if err != nil {
		return []commonM.Message{message.NewServerInternalError()}
	}
	return nil
}
