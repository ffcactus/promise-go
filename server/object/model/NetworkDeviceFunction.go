package model

// Ethernet This type describes Ethernet capabilities, status, and configuration of a network device function.
type Ethernet struct {
	PermanentMACAddress string // This is the permanent MAC address assigned to this network device function (physical function).
	MACAddress          string // This is the currently configured MAC address of the (logical port) network device function.
	MTUSize             int    // The Maximum Transmission Unit (MTU) configured for this network device function.
}

// ISCSIBoot This type describes iSCSI boot capabilities, status, and configuration of a network device function.
type ISCSIBoot struct {
	IPAddressType              string // The type of IP address (IPv6 or IPv4) being populated in the iSCSIBoot IP address fields.
	InitiatorIPAddress         string // The IPv6 or IPv4 address of the iSCSI initiator.
	InitiatorName              string // The iSCSI initiator name.
	InitiatorDefaultGateway    string // The IPv6 or IPv4 iSCSI boot default gateway.
	InitiatorNetmask           string // The IPv6 or IPv4 netmask of the iSCSI boot initiator.
	TargetInfoViaDHCP          bool   // Whether the iSCSI boot target name, LUN, IP address, and netmask should be obtained from DHCP.
	PrimaryTargetName          string // The name of the iSCSI primary boot target.
	PrimaryTargetIPAddress     string // The IP address (IPv6 or IPv4) for the primary iSCSI boot target.
	PrimaryTargetTCPPort       int    // The TCP port for the primary iSCSI boot target.
	PrimaryLUN                 int    // The logical unit number (LUN) for the primary iSCSI boot target.
	PrimaryVLANEnable          bool   // This indicates if the primary VLAN is enabled.
	PrimaryVLANID              int    // The 802.1q VLAN ID to use for iSCSI boot from the primary target.
	PrimaryDNS                 string // The IPv6 or IPv4 address of the primary DNS server for the iSCSI boot initiator.
	SecondaryTargetName        string // The name of the iSCSI secondary boot target.
	SecondaryTargetIPAddress   string // The IP address (IPv6 or IPv4) for the secondary iSCSI boot target.
	SecondaryTargetTCPPort     int    // The TCP port for the secondary iSCSI boot target.
	SecondaryLUN               int    // The logical unit number (LUN) for the secondary iSCSI boot target.
	SecondaryVLANEnable        bool   // This indicates if the secondary VLAN is enabled.
	SecondaryVLANID            int    // The 802.1q VLAN ID to use for iSCSI boot from the secondary target.
	SecondaryDNS               string // The IPv6 or IPv4 address of the secondary DNS server for the iSCSI boot initiator.
	IPMaskDNSViaDHCP           bool   // Whether the iSCSI boot initiator uses DHCP to obtain the iniator name, IP address, and netmask.
	RouterAdvertisementEnabled bool   // Whether IPv6 router advertisement is enabled for the iSCSI boot target.
	AuthenticationMethod       string // The iSCSI boot authentication method for this network device function.
	CHAPUsername               string // The username for CHAP authentication.
	CHAPSecret                 string // The shared secret for CHAP authentication.
	MutualCHAPUsername         string // The CHAP Username for 2-way CHAP authentication.
	MutualCHAPSecret           string // The CHAP Secret for 2-way CHAP authentication.
}

// BootTargets A Fibre Channel boot target configured for a network device function.
type BootTargets struct {
	WWPN         string // The World-Wide Port Name to boot from.
	LUNID        string // The Logical Unit Number (LUN) ID to boot from on the device referred to by the corresponding WWPN.
	BootPriority int    // The relative priority for this entry in the boot targets array.
}

// FibreChannel This type describes Fibre Channel capabilities, status, and configuration of a network device function.
type FibreChannel struct {
	PermanentWWPN         string      // This is the permanent WWPN address assigned to this network device function (physical function).
	PermanentWWNN         string      // This is the permanent WWNN address assigned to this network device function (physical function).
	WWPN                  string      // This is the currently configured WWPN address of the network device function (physical function).
	WWNN                  string      // This is the currently configured WWNN address of the network device function (physical function).
	WWNSource             string      // The configuration source of the WWNs for this connection (WWPN and WWNN).
	FCoELocalVLANID       int         // The locally configured FCoE VLAN ID.
	AllowFIPVLANDiscovery bool        // Whether the FCoE Initialization Protocol (FIP) is used for populating the FCoE VLAN Id.
	FCoEActiveVLANID      int         // The active FCoE VLAN ID.
	BootTargets           BootTargets // An array of Fibre Channel boot targets configured for this network device function.
}

// NetworkDeviceFunctionLinks Contains references to other resources that are related to this resource.
type NetworkDeviceFunctionLinks struct {
	PCIeFunction string // Contains the members of this collection.
}

// NetworkDeviceFunction A Network Device Function represents a logical interface exposed by the network adapter.
type NetworkDeviceFunction struct {
	Resource
	NetDevFuncType          string                     // The configured capability of this network device function.
	DeviceEnabled           bool                       // Whether the network device function is enabled.
	NetDevFuncCapabilities  []string                   // Capabilities of this network device function.
	Ethernet                Ethernet                   // Ethernet.
	ISCSIBoot               ISCSIBoot                  // iSCSI Boot.
	FibreChannel            FibreChannel               // Fibre Channel.
	BootMode                string                     // The boot mode configured for this network device function.
	VirtualFunctionsEnabled bool                       // Whether Single Root I/O Virtualization (SR-IOV) Virual Functions (VFs) are enabled for this Network Device Function.
	MaxVirtualFunctions     int                        // The number of virtual functions (VFs) that are available for this Network Device Function.
	Links                   NetworkDeviceFunctionLinks // Links for this NetworkDeviceFunction.
	AssignablePhysicalPorts []NetworkPort              // Contains the members of this collection.
	PhysicalPortAssignment  NetworkPort                // Contains the members of this collection.
}
