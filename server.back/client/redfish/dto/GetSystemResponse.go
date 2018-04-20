package dto

// Contains references to other resources that are related to this resource.
type ComputerSystemLinks struct {
	Chassis        *[]OdataID `json:"Chassis"`        // An array of references to the chassis in which this system is contained.
	ManagedBy      *[]OdataID `json:"ManagedBy"`      // An array of references to the Managers responsible for this system.
	PoweredBy      *[]OdataID `json:"PoweredBy"`      // An array of ID[s] of resources that power this computer system. Normally the ID will be a chassis or a specific set of Power Supplies.
	CooledBy       *[]OdataID `json:"CooledBy"`       // An array of ID[s] of resources that cool this computer system. Normally the ID will be a chassis or a specific set of fans.
	Endpoints      *[]OdataID `json:"Endpoints"`      // An array of references to the endpoints that connect to this system.
	ResourceBlocks *[]OdataID `json:"ResourceBlocks"` // An array of references to the Resource Blocks that are used in this Computer System.
}

type Boot struct {
	BootSourceOverrideTarget     *string `json:"BootSourceOverrideTarget"`     // The current boot source to be used at next boot instead of the normal boot device, if BootSourceOverrideEnabled is true.
	BootSourceOverrideEnabled    *string `json:"BootSourceOverrideEnabled"`    // Describes the state of the Boot Source Override feature.
	UefiTargetBootSourceOverride *string `json:"UefiTargetBootSourceOverride"` // This property is the UEFI Device Path of the device to boot from when BootSourceOverrideSupported is UefiTarget.
	BootSourceOverrideMode       *string `json:"BootSourceOverrideMode"`       // The BIOS Boot Mode (either Legacy or UEFI) to be used when BootSourceOverrideTarget boot source is booted from.
}

type ComputerSystemOemHuawei struct {
	// TODO
}

type ComputerSystemOem struct {
	Huawei *ComputerSystemOemHuawei `json:"Huawei"`
}

type GetSystemResponse struct {
	Resource
	UUID
	Links              *ComputerSystemLinks `json:"Links"`
	SystemType         *string              `json:"SystemType"`
	HostName           *string              `json:"HostName"`
	IndicatorLED       *string              `json:"IndicatorLED"`
	PowerState         *string              `json:"PowerState"`
	Boot               *Boot                `json:"Boot"`
	BiosVersion        *string              `json:"BiosVersion"`
	Processors         *OdataID             `json:"Processors"`
	EthernetInterfaces *OdataID             `json:"EthernetInterfaces"`
	Storage            *OdataID             `json:"Storage"`
	PCIeDevices        *[]OdataID           `json:"PCIeDevices"`
	PCIeFunctions      *[]OdataID           `json:"PCIeFunctions"`
	MemoryDomains      *OdataID             `json:"MemoryDomains"`
	NetworkInterfaces  *OdataID             `json:"NetworkInterfaces"`
	Oem                *ComputerSystemOem   `json:"Oem"`
}
