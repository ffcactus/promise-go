package dto

import (
	"fmt"
	"promise/server/object/model"
)

// This is the schema definition for the PCIeFunction resource.  It represents the properties of a PCIeFunction attached to a System.
type PCIeFunction struct {
	ResourceResponse
	DeviceClass        *string       // The class for this PCIe Function.
	DeviceID           *string       // The Device ID of this PCIe function.
	VendorID           *string       // The Vendor ID of this PCIe function.
	SubsystemID        *string       // The Subsystem ID of this PCIe function.
	SubsystemVendorID  *string       // The Subsystem Vendor ID of this PCIe function.
	EthernetInterfaces []ResourceRef // An array of references to the ethernet interfaces which the PCIe device produces.
}

func (this *PCIeFunction) Load(m *model.PCIeFunction, ethernetInterfaces []model.EthernetInterface) {
	this.LoadResourceResponse(&m.Resource)
	this.DeviceClass = m.DeviceClass
	this.DeviceID = m.DeviceID
	this.VendorID = m.VendorID
	this.SubsystemID = m.SubsystemID
	this.SubsystemVendorID = m.SubsystemVendorID
	for i, _ := range m.EthernetInterfaces {
		src := m.EthernetInterfaces[i]
		for j, _ := range ethernetInterfaces {
			target := ethernetInterfaces[j]
			if (target.URI != nil) && (src == *target.URI) {
				ref := ResourceRef{}
				ref.Ref = fmt.Sprintf("#/ComputerSystem/EthernetInterfaces/%d", j)
				this.EthernetInterfaces = append(this.EthernetInterfaces, ref)
			}
		}
	}
}
