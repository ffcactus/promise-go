package model

// ServerHardwareModel describe the hardware info of a server.
// Servers with the same ServerHardwareModel can take the same config.
type ServerHardwareModel struct {
	Type    string
	Version string
	Adapter []Adapter
}

// Adapter is a device on the server, for example the Mezz card.
type Adapter struct {
	Type    string // Generally, this is the slot type.
	Name    string // The public name of this adapter.
	Version string // This value is used to distinguish the adapter with the same name.
	Slot    int
}
