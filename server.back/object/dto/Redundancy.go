package dto

import (
	"promise/server/object/model"
)

// Redundancy This is the redundancy definition to be used in other resource schemas.
type Redundancy struct {
	ResourceResponse
	Mode              *string   `json:"Mode,omitempty"`              // This is the redundancy mode of the group.
	MaxNumSupported   *int      `json:"MaxNumSupported,omitempty"`   // This is the maximum number of members allowable for this particular redundancy group.
	MinNumNeeded      *int      `json:"MinNumNeeded,omitempty"`      // This is the minumum number of members needed for this group to be redundant.
	RedundancySet     *[]string `json:"RedundancySet,omitempty"`     // Contains any ids that represent components of this redundancy set.
	RedundancyEnabled *bool     `json:"RedundancyEnabled,omitempty"` // This indicates whether redundancy is enabled.
}

// Load will load data from model.
func (dto *Redundancy) Load(m *model.Redundancy) {
	dto.LoadResourceResponse(&m.Resource)
	dto.Mode = m.Mode
	dto.MaxNumSupported = m.MaxNumSupported
	dto.MinNumNeeded = m.MinNumNeeded
	dto.RedundancyEnabled = m.RedundancyEnabled
	dto.RedundancySet = m.RedundancySet
}
