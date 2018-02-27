package dto

type VLANID int

// This type describes the attributes of a Virtual LAN.
type VLAN struct {
	VLANEnable *bool // This indicates if this VLAN is enabled.
	VLANID     *int  // This indicates the VLAN identifier for this VLAN.
}

// This resource describes the attributes of a Virtual LAN.
type GetVLANResponse struct {
	Resource
	VLANEnable *bool // This indicates if this VLAN is enabled.
	VLANID     *int  // This indicates the VLAN identifier for this VLAN.
}
