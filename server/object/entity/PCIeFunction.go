package entity

import (
	"promise/base"
	"promise/server/object/model"
)

// PCIeFunction This is the schema definition for the PCIeFunction resource.  It represents the properties of a PCIeFunction attached to a System.
type PCIeFunction struct {
	PCIeDeviceRef uint
	EmbeddedResource
	DeviceClass        *string // The class for this PCIe Function.
	DeviceID           *string // The Device ID of this PCIe function.
	VendorID           *string // The Vendor ID of this PCIe function.
	SubsystemID        *string // The Subsystem ID of this PCIe function.
	SubsystemVendorID  *string // The Subsystem Vendor ID of this PCIe function.
	EthernetInterfaces string  // An array of references to the ethernet interfaces which the PCIe device produces.
}

// ToModel will create a new model from entity.
func (e *PCIeFunction) ToModel() *model.PCIeFunction {
	m := new(model.PCIeFunction)
	createResourceModel(&e.EmbeddedResource, &m.Resource)
	m.DeviceClass = e.DeviceClass
	m.DeviceID = e.DeviceID
	m.VendorID = e.VendorID
	m.SubsystemID = e.SubsystemID
	m.SubsystemVendorID = e.SubsystemVendorID
	a := []string{}
	base.StringToStruct(e.EthernetInterfaces, &a)
	m.EthernetInterfaces = a
	return m
}

// Load will load data from model.
func (e *PCIeFunction) Load(m *model.PCIeFunction) {
	updateResourceEntity(&e.EmbeddedResource, &m.Resource)
	e.DeviceClass = m.DeviceClass
	e.DeviceID = m.DeviceID
	e.VendorID = m.VendorID
	e.SubsystemID = m.SubsystemID
	e.SubsystemVendorID = m.SubsystemVendorID
	s := base.StructToString(m.EthernetInterfaces)
	e.EthernetInterfaces = s
}
