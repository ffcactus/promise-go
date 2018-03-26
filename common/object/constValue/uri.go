package constValue

import (
	"promise/common/app"
	"promise/common/category"
)

const (
	// ServerBaseURI is server base URI.
	ServerBaseURI = "/server"
	// ServerGroupBaseURI is servergroup base URI.
	ServerGroupBaseURI = "/servergroup"
	// ServerServerGroupBaseURI is server-servergroup base URI.
	ServerServerGroupBaseURI = "/server-servergroup"
	// AuthBaseURI is auth base URI.
	AuthBaseURI = "/auth"
	// WSBaseURI is websocket base URI.
	WSBaseURI = "/ws"
	// WSSenderBaseURI is websocket sender base URI.
	WSSenderBaseURI = "/ws-sender"
	// TaskBaseURI is task base URI.
	TaskBaseURI = "/task"
)

// CategoryToURI turns ID to to URI depends on category.
func CategoryToURI(c string, id string) string {
	switch c {
	case category.AA:
		return app.RootURL + AuthBaseURI + "/" + id
	case category.Task:
		return app.RootURL + TaskBaseURI + "/" + id
	case category.Server:
		return app.RootURL + ServerBaseURI + "/" + id
	case category.ServerGroup:
		return app.RootURL + ServerGroupBaseURI + "/" + id
	case category.ServerServerGroup:
		return app.RootURL + ServerServerGroupBaseURI + "/" + id
	default:
		return ""
	}
}

// ToServerURI convert id to URI
func ToServerURI(id string) string {
	return app.RootURL + ServerBaseURI + "/" + id
}

// ToServerGroupURI convert id to URI
func ToServerGroupURI(id string) string {
	return app.RootURL + ServerGroupBaseURI + "/" + id
}

// ToServerServerGroupURI convert id to URI
func ToServerServerGroupURI(id string) string {
	return app.RootURL + ServerServerGroupBaseURI + "/" + id
}
