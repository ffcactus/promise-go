package dto

import (
	"fmt"
	"promise/server/object/model"
)

// A NetworkInterface contains references linking NetworkAdapter, NetworkPort, and NetworkDeviceFunction resources and represents the functionality available to the containing system.
type NetworkInterface struct {
	ResourceResponse
	NetworkAdapter ResourceRef
}

func (dto *NetworkInterface) Load(m *model.NetworkInterface, networkAdapters []model.NetworkAdapter) {
	dto.LoadResourceResponse(&m.Resource)
	for i := range networkAdapters {
		target := networkAdapters[i]
		if (target.URI != nil) && (m.NetworkAdapterURI == *target.URI) {
			ref := ResourceRef{}
			ref.Ref = fmt.Sprintf("#/Chassis/NetworkAdapters/%d", i)
			dto.NetworkAdapter = ref
		}
	}
}
