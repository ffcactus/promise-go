package service

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	. "promise/server/client"
	"promise/server/object/dto"
	. "promise/server/object/model"
)

func Probe(request *dto.PostServerRequest) (*ServerBasicInfo, error) {
	client := FindBestClient(request.Address, request.Username, request.Password)
	if client == nil {
		log.Fatal("Probe() failed, failed to get the client for the server, address = ", request.Address)
		return nil, fmt.Errorf("Failed to get server client.")
	}

	serverBasicInfo, err := client.GetBasicInfo()
	if err != nil {
		log.Fatal("Probe() failed, failed to get basic info, address = ", request.Address, ", error = ", err)
	}

	serverBasicInfo.Address = request.Address
	serverBasicInfo.OriginUsername = &request.Username
	serverBasicInfo.OriginPassword = &request.Password
	serverBasicInfo.Name = "Name???"
	serverBasicInfo.Description = "Description???"
	return serverBasicInfo, nil
}
