package client

import (
	log "github.com/sirupsen/logrus"
	"promise/server/client/mock"
	"promise/server/client/redfish"
	"promise/server/object/constValue"
	"promise/server/object/model"
	"strings"
)

// ServerClientInterface is the server client interface.
type ServerClientInterface interface {
	Support() bool
	GetProtocol() string
	GetBasicInfo() (*model.ServerBasicInfo, error)
	CreateManagementAccount(username string, password string) error
	GetProcessors(systemID string) ([]model.Processor, error)
	GetMemory(systemID string) ([]model.Memory, error)
	GetEthernetInterfaces(systemID string) ([]model.EthernetInterface, error)
	GetNetworkInterfaces(systemID string) ([]model.NetworkInterface, error)
	GetStorages(systemID string) ([]model.Storage, error)
	// For chassis info
	GetOemHuaweiBoards(chassisID string) ([]model.OemHuaweiBoard, error)
	GetPower(chassisID string) (*model.Power, error)
	GetThermal(chassisID string) (*model.Thermal, error)
	GetNetworkAdapters(chassisID string) ([]model.NetworkAdapter, error)
	GetDrives(chassisID string) ([]model.Drive, error)
	GetPCIeDevices(chassisID string) ([]model.PCIeDevice, error)
}

// FindBestClient will find the best client for the server.
func FindBestClient(hostname string, username string, password string) ServerClientInterface {
	var client ServerClientInterface
	client = mock.GetInstance(hostname)
	if client.Support() {
		return client
	}
	client = redfish.GetInstance(hostname, username, password, true)
	if client.Support() {
		return client
	}
	log.WithFields(log.Fields{"hostname": hostname}).Warn("FindBestClient(), can't find a client, server address = ", hostname)
	return nil
}

// TODO.
func getServerManagementAccount(server *model.Server) (string, string) {
	if server.OriginUsername != nil && *server.OriginUsername != "" {
		return *server.OriginUsername, *server.OriginPassword
	}
	s := strings.Split(server.Credential, " ")
	if len(s) == 2 {
		return s[0], s[1]
	}
	log.Warn("getServerManagementAccount(), failed to get username and password.")
	return "", ""
}

// GetServerClient will return the server client based on protocol.
func GetServerClient(server *model.Server) ServerClientInterface {
	switch server.Protocol {
	case constValue.RedfishV1:
		username, password := getServerManagementAccount(server)
		return redfish.GetInstance(server.Hostname, username, password, true)
	case constValue.MockProtocol:
		return mock.GetInstance(server.Hostname)
	default:
		return nil
	}
}
