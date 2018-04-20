package model

// Redundancy This is the redundancy definition to be used in other resource schemas.
type Redundancy struct {
	Resource
	Mode              *string   // This is the redundancy mode of the group.
	MaxNumSupported   *int      // This is the maximum number of members allowable for this particular redundancy group.
	MinNumNeeded      *int      // This is the minumum number of members needed for this group to be redundant.
	RedundancySet     *[]string // Contains any ids that represent components of this redundancy set.
	RedundancyEnabled *bool     // This indicates whether redundancy is enabled.
}
