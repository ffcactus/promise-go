package constValue

import (
	"promise/common/app"
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
func CategoryToURI(category string, id string) string {
	switch category {
	case CategoryAA:
		return app.RootURL + AuthBaseURI + "/" + id
	case CategoryTask:
		return app.RootURL + TaskBaseURI + "/" + id
	case CategoryServer:
		return app.RootURL + ServerBaseURI + "/" + id
	case CategoryServerGroup:
		return app.RootURL + ServerGroupBaseURI + "/" + id
	case CategoryServerServerGroup:
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
