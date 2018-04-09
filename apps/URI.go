package apps

const (
	// ProtocolScheme is the protocol scheme used by client.
	ProtocolScheme = "http://"
	// Host is the host used by client.
	Host = "localhost"
	// RootURL is the root URI for all the service.
	RootURL = "/promise/v1"

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
	// IDPoolBaseURI is pool base URI.
	IDPoolBaseURI = "/id-pool"
	// IDPoolIPv4BaseURI is IPv4 pool base URI.
	IDPoolIPv4BaseURI = "/id-pool/ipv4"
	// StudentBaseURI is student base URI.
	StudentBaseURI = "/student"
)

// CategoryToURI turns ID to to URI depends on category.
func CategoryToURI(c string, id string) string {
	switch c {
	case CategoryAA:
		return RootURL + AuthBaseURI + "/" + id
	case CategoryTask:
		return RootURL + TaskBaseURI + "/" + id
	case CategoryServer:
		return RootURL + ServerBaseURI + "/" + id
	case CategoryServerGroup:
		return RootURL + ServerGroupBaseURI + "/" + id
	case CategoryServerServerGroup:
		return RootURL + ServerServerGroupBaseURI + "/" + id
	case CategoryPoolIPv4:
		return RootURL + IDPoolIPv4BaseURI + "/" + id
	case CategoryStudent:
		return RootURL + StudentBaseURI + "/" + id
	default:
		return ""
	}
}

// ToServerURI convert ID to URI.
func ToServerURI(id string) string {
	return RootURL + ServerBaseURI + "/" + id
}

// ToServerGroupURI convert ID to URI.
func ToServerGroupURI(id string) string {
	return RootURL + ServerGroupBaseURI + "/" + id
}

// ToServerServerGroupURI convert ID to URI.
func ToServerServerGroupURI(id string) string {
	return RootURL + ServerServerGroupBaseURI + "/" + id
}

// ToIDPoolIPv4URI convert ID to URI.
func ToIDPoolIPv4URI(id string) string {
	return RootURL + IDPoolIPv4BaseURI + "/" + id
}

// ToStudentURI convert ID to URI.
func ToStudentURI(id string) string {
	return RootURL + StudentBaseURI + "/" + id
}
