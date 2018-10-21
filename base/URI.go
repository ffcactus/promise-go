package base

const (
	// ProtocolScheme is the protocol scheme used by client.
	ProtocolScheme = "http://"
	// RootURL is the root URI for all the service.
	RootURL = "/promise/v1"

	// DirectorBaseURI is director base URI.
	DirectorBaseURI = "/director"
	// DirecotrNodeBaseURI is node base URI.
	DirecotrNodeBaseURI = "/director/node"
	// ServerBaseURI is server base URI.
	ServerBaseURI = "/server"
	// ServerGroupBaseURI is servergroup base URI.
	ServerGroupBaseURI = "/servergroup"
	// ServerServerGroupBaseURI is server-servergroup base URI.
	ServerServerGroupBaseURI = "/server-servergroup"
	// EnclosureBaseURI is enclosure base URI.
	EnclosureBaseURI = "/enclosure"
	// AdapterConfigBaseURI is adapterconfig base URI.
	AdapterConfigBaseURI = "/adapterconfig"
	// AdapterModelBaseURI is adapterconfig base URI.
	AdapterModelBaseURI = "/adaptermodel"
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
)

// CategoryToURI turns ID to to URI depends on category.
func CategoryToURI(c string, id string) string {
	switch c {
	case CategoryNode:
		return RootURL + DirecotrNodeBaseURI + "/" + id
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
	case CategoryAdapterConfig:
		return RootURL + AdapterConfigBaseURI + "/" + id
	case CategoryAdapterModel:
		return RootURL + AdapterModelBaseURI + "/" + id
	case CategoryEnclosure:
		return RootURL + EnclosureBaseURI + "/" + id
	case CategoryPoolIPv4:
		return RootURL + IDPoolIPv4BaseURI + "/" + id
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

// ToAdapterConfigURI convert ID to URI.
func ToAdapterConfigURI(id string) string {
	return RootURL + AdapterConfigBaseURI + "/" + id
}

// ToAdapterModelURI convert ID to URI.
func ToAdapterModelURI(id string) string {
	return RootURL + AdapterModelBaseURI + "/" + id
}

// ToEnclosureURI convert ID to URI.
func ToEnclosureURI(id string) string {
	return RootURL + EnclosureBaseURI + "/" + id
}

// ToIDPoolIPv4URI convert ID to URI.
func ToIDPoolIPv4URI(id string) string {
	return RootURL + IDPoolIPv4BaseURI + "/" + id
}
