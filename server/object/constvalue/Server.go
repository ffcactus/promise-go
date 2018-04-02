package constvalue

const (
	// ServerStateAdded means the server is added but not ready.
	ServerStateAdded = "Added"
	// ServerStateReady means the server is ready for operation.
	ServerStateReady = "Ready"
	// ServerStateLocked means the server is locked, operation is not allowed.
	ServerStateLocked = "Locked"
	// ServerStateDeleting means the server is deleting.
	ServerStateDeleting = "Deleting"
)

const (
	// ServerActionRefresh is the action to refresh server.
	ServerActionRefresh = "refresh"
)

// Server health.
const (
	// ServerHealthOK is server health state.
	ServerHealthOK = "OK"
	// ServerHealthWarning is server health state.
	ServerHealthWarning = "Warning"
	// ServerHealthCritical is server health state.
	ServerHealthCritical = "Critical"
	// ServerHealthUnknown is server health state.
	ServerHealthUnknown = "Unknown"
)

const (
	// MockProtocol Protocol enum.
	MockProtocol = "MockProtocol"
	// RedfishV1 Protocol enum.
	RedfishV1 = "RedfishV1"
)

const (
	// MockType Server type enum.
	MockType = "Mock"
	// BladeType Server type enum.
	BladeType = "Blade"
	// RackType Server type enum.
	RackType = "Rack"
	// EnclosureType Server type enum.
	EnclosureType = "Enclosure"
	// SwitchType Server type enum.
	SwitchType = "Switch"
	// UnknownServerType Server type enum.
	UnknownServerType = "Unknown"
)

// ServerLockable will return wether the server can be locked.
func ServerLockable(state string) bool {
	var ret bool
	switch state {
	case ServerStateAdded:
		ret = true
	case ServerStateReady:
		ret = true
	default:
		ret = false
	}
	return ret
}
