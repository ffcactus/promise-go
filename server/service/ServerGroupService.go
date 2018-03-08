package service

import (
	"promise/server/object/dto"
	"promise/server/object/model"
	commonM "promise/common/object/model"
	"github.com/astaxie/beego"
)

// PostServerGroup post a server group.
func PostServerGroup(request *dto.PostServerGroupRequest)(*model.ServerGroup, []commonM.Message) {
	beego.Info("Post server group", request.Name)
	return nil, nil
}