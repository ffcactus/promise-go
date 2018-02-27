package util

import ()

// Server state.
var (
	ServerStateAdded    string = "Added"
	ServerStateReady    string = "Ready"
	ServerStateLocked   string = "Locked"
	ServerStateDeleting string = "Deleting"
)

var (
	ServerActionRefresh string = "refresh"
)

// Server health.
var (
	ServerHealthOK       string = "OK"
	ServerHealthWarning  string = "Warning"
	ServerHealthCritical string = "Critical"
	ServerHealthUnknown  string = "Unknown"
)

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
