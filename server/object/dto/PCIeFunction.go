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

func (dto *PCIeFunction) Load(m *model.PCIeFunction, ethernetInterfaces []model.EthernetInterface) {
	dto.LoadResourceResponse(&m.Resource)
	dto.DeviceClass = m.DeviceClass
	dto.DeviceID = m.DeviceID
	dto.VendorID = m.VendorID
	dto.SubsystemID = m.SubsystemID
	dto.SubsystemVendorID = m.SubsystemVendorID
	for i := range m.EthernetInterfaces {
		src := m.EthernetInterfaces[i]
		for j := range ethernetInterfaces {
			target := ethernetInterfaces[j]
			if (target.URI != nil) && (src == *target.URI) {
				ref := ResourceRef{}
				ref.Ref = fmt.Sprintf("#/ComputerSystem/EthernetInterfaces/%d", j)
				dto.EthernetInterfaces = append(dto.EthernetInterfaces, ref)
			}
		}
	}
}
