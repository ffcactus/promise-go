package entity

import (
	"promise/server/object/model"
)

// NetworkInterface A NetworkInterface contains references linking NetworkAdapter, NetworkPort, and NetworkDeviceFunction resources and represents the functionality available to the containing system.
type NetworkInterface struct {
	ServerRef string
	EmbeddedResource
	NetworkAdapterURI string
}

// ToModel will create a new model from entity.
func (e *NetworkInterface) ToModel() *model.NetworkInterface {
	m := model.NetworkInterface{}
	createResourceModel(&e.EmbeddedResource, &m.Resource)
	m.NetworkAdapterURI = e.NetworkAdapterURI
	return &m
}

// Load will load data from model.
func (e *NetworkInterface) Load(m *model.NetworkInterface) {
	updateResourceEntity(&e.EmbeddedResource, &m.Resource)
	e.NetworkAdapterURI = m.NetworkAdapterURI
}
