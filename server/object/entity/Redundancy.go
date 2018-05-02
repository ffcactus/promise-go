package entity

import (
	"promise/base"
	"promise/server/object/model"
)

// Redundancy This is the redundancy definition to be used in other resource schemas.
type Redundancy struct {
	EmbeddedResource
	Ref               uint
	Mode              *string // This is the redundancy mode of the group.
	MaxNumSupported   *int    // This is the maximum number of members allowable for this particular redundancy group.
	MinNumNeeded      *int    // This is the minumum number of members needed for this group to be redundant.
	RedundancySet     *string // Contains any ids that represent components of this redundancy set.
	RedundancyEnabled *bool   // This indicates whether redundancy is enabled.
}

// ToModel will create a new model from entity.
func (e *Redundancy) ToModel() *model.Redundancy {
	m := model.Redundancy{}
	createResourceModel(&e.EmbeddedResource, &m.Resource)
	m.Mode = e.Mode
	m.MaxNumSupported = e.MaxNumSupported
	m.MinNumNeeded = e.MinNumNeeded
	m.RedundancyEnabled = e.RedundancyEnabled
	a := []string{}
	base.StringToStruct(*e.RedundancySet, &a)
	m.RedundancySet = &a
	return &m
}
