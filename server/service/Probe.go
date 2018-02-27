package service

import (
	. "promise/server/client"
	"promise/server/object/dto"
	. "promise/server/object/model"
	"fmt"
	"github.com/astaxie/beego"
)

func Probe(request *dto.PostServerRequest) (*ServerBasicInfo, error) {
	client := FindBestClient(request.Address, request.Username, request.Password)
	if client == nil {
		beego.Critical("Probe() failed, failed to get the client for the server, address = ", request.Address)
		return nil, fmt.Errorf("Failed to get server client.")
	}

	serverBasicInfo, err := client.GetBasicInfo()
	if err != nil {
		beego.Critical("Probe() failed, failed to get basic info, address = ", request.Address, ", error = ", err)
	}

	serverBasicInfo.Address = request.Address
	serverBasicInfo.OriginUsername = &request.Username
	serverBasicInfo.OriginPassword = &request.Password
	serverBasicInfo.Name = "Name???"
	serverBasicInfo.Description = "Description???"
	return serverBasicInfo, nil
}
