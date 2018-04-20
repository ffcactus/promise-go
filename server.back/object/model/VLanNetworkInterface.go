package model

// VLAN This type describes the attributes of a Virtual LAN.
type VLAN struct {
	VLANEnable *bool // This indicates if this VLAN is enabled.
	VLANID     *int  // This indicates the VLAN identifier for this VLAN.
}

// VLanNetworkInterface This resource describes the attributes of a Virtual LAN.
type VLanNetworkInterface struct {
	Resource
	VLANEnable *bool // This indicates if this VLAN is enabled.
	VLANID     *int  // This indicates the VLAN identifier for this VLAN.
}
