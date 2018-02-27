package dto

import (
	"promise/server/object/model"
)

// This is the redundancy definition to be used in other resource schemas.
type Redundancy struct {
	ResourceResponse
	Mode              *string   `json:"Mode,omitempty"`              // This is the redundancy mode of the group.
	MaxNumSupported   *int      `json:"MaxNumSupported,omitempty"`   // This is the maximum number of members allowable for this particular redundancy group.
	MinNumNeeded      *int      `json:"MinNumNeeded,omitempty"`      // This is the minumum number of members needed for this group to be redundant.
	RedundancySet     *[]string `json:"RedundancySet,omitempty"`     // Contains any ids that represent components of this redundancy set.
	RedundancyEnabled *bool     `json:"RedundancyEnabled,omitempty"` // This indicates whether redundancy is enabled.
}

func (this *Redundancy) Load(m *model.Redundancy) {
	this.LoadResourceResponse(&m.Resource)
	this.Mode = m.Mode
	this.MaxNumSupported = m.MaxNumSupported
	this.MinNumNeeded = m.MinNumNeeded
	this.RedundancyEnabled = m.RedundancyEnabled
	this.RedundancySet = m.RedundancySet
}
