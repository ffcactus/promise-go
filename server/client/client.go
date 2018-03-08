package client

import (
	"github.com/astaxie/beego"
	"promise/server/client/mock"
	"promise/server/client/redfish"
	. "promise/server/object/model"
	"strings"
)

type ServerClientInterface interface {
	Support() bool
	GetProtocol() *string
	GetBasicInfo() (*ServerBasicInfo, error)
	CreateManagementAccount(username string, password string) error
	GetProcessors(systemID string) ([]Processor, error)
	GetMemory(systemID string) ([]Memory, error)
	GetEthernetInterfaces(systemID string) ([]EthernetInterface, error)
	GetNetworkInterfaces(systemID string) ([]NetworkInterface, error)
	GetStorages(systemID string) ([]Storage, error)
	// For chassis info
	GetOemHuaweiBoards(chassisID string) ([]OemHuaweiBoard, error)
	GetPower(chassisID string) (*Power, error)
	GetThermal(chassisID string) (*Thermal, error)
	GetNetworkAdapters(chassisID string) ([]NetworkAdapter, error)
	GetDrives(chassisID string) ([]Drive, error)
	GetPCIeDevices(chassisID string) ([]PCIeDevice, error)
}

// Find the best client for the server.
func FindBestClient(address string, username string, password string) ServerClientInterface {
	var client ServerClientInterface
	client = mock.GetInstance(address)
	if client.Support() {
		return client
	}
	client = redfish.GetInstance(address, username, password, true)
	if client.Support() {
		return client
	}
	beego.Warning("FindBestClient(), can't find a client, server address = ", address)
	return nil
}

// TODO.
func getServerManagementAccount(server *Server) (string, string) {
	if server.OriginUsername != nil && *server.OriginUsername != "" {
		return *server.OriginUsername, *server.OriginPassword
	}
	s := strings.Split(server.Credential, " ")
	if len(s) == 2 {
		return s[0], s[1]
	}
	beego.Warning("getServerManagementAccount(), failed to get username and password.")
	return "", ""
}

func GetServerClient(server *Server) ServerClientInterface {
	switch server.Protocol {
	case RedfishV1:
		username, password := getServerManagementAccount(server)
		return redfish.GetInstance(server.Address, username, password, true)
	case MockProtocol:
		return mock.GetInstance(server.Address)
	default:
		return nil
	}
}
