package constvalue

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
)

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
