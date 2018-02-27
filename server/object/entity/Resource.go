package entity

import (
	"time"
)

// Entity entity object.
type Entity struct {
	ID string `gorm:"primary_key"`
}

// MemberID Member ID.
type MemberID struct {
	MemberID string
}

// Status Status struct.
type Status struct {
	State          *string
	Health         *string
	PhysicalState  *string
	PhysicalHealth *string
}

// Resource Resource means the main resources that is index by UUID.
type Resource struct {
	Entity
	Status
	CreatedAt   time.Time
	UpdatedAt   time.Time
	URI         *string
	OriginID    *string
	Name        *string
	Description *string
}

// EmbeddedResource The EmbeddedResource is the resources that belongs to Resource, for example the memory in server.
// They has the resources properties since they are resource object from BMC perspective, when they are
// embedded in Resource in DB, they are retrieved by uint ID.
type EmbeddedResource struct {
	ID uint `gorm:"primary_key"`
	Status
	URI         *string
	OriginID    *string
	Name        *string
	Description *string
}

// EmbeddedObject The EmbeddedObject is sub structure in Resource or EmbeddedResource.
type EmbeddedObject struct {
	ID uint `gorm:"primary_key"`
}

// Member Member is a object in a collection. Each member should have a member ID.
type Member struct {
	ID uint `gorm:"primary_key"`
	Status
	URI            *string
	OriginMemberID *string
	Name           *string
	Description    *string
}

// ResourceCollection Resource collection.
type ResourceCollection struct {
	Description  *string
	Name         *string
	MembersCount *int
}

// ProductInfo Product info.
type ProductInfo struct {
	Model           *string // The model string for this product.
	Manufacturer    *string
	SKU             *string // The SKU string of this product.
	SerialNumber    *string // The serial number for this resource.
	PartNumber      *string // The part number for this resource.
	SparePartNumber *string // The spare part number for this resource.
	AssetTag        *string // The value of this property shall be an identifying string used to track the resource for inventory purposes.
}

// Threshold The common Threshold information.
type Threshold struct {
	UpperThresholdNonCritical *float64 // Above normal range.
	UpperThresholdCritical    *float64 // Above normal range but not yet fatal.
	UpperThresholdFatal       *float64 // Above normal range and is fatal.
	LowerThresholdNonCritical *float64 // Below normal range.
	LowerThresholdCritical    *float64 // Below normal range but not yet fatal.
	LowerThresholdFatal       *float64 // Below normal range and is fatal.
}

// Placement The placement within the addressed location.
type Placement struct {
	LocationRef uint
	EmbeddedObject
	Row             *string // Name of row.
	Rack            *string // Name of a rack location within a row.
	RackOffsetUnits *string // The type of Rack Units in use.
	RackOffset      *int    // Vertical location of the item in terms of RackOffsetUnits.
}

// PostalAddress The PostalAddress for a resource.
type PostalAddress struct {
	LocationRef uint
	EmbeddedObject
	Country                *string // Country.
	Territory              *string // A top-level subdivision within a country.
	District               *string // A county, parish, gun (JP), or  district (IN).
	City                   *string // City, township, or shi (JP).
	Division               *string // City division, borough, dity district, ward, chou (JP).
	Neighborhood           *string // Neighborhood or block.
	LeadingStreetDirection *string // A leading street direction.
	Street                 *string // Street name.
	TrailingStreetSuffix   *string // A trailing street suffix.
	StreetSuffix           *string // Avenue, Platz, Street, Circle.
	HouseNumber            *int    // Numeric portion of house number.
	HouseNumberSuffix      *string // House number suffix.
	Landmark               *string // Landmark.
	Location               *string // Room designation or other additional info.
	Floor                  *string // Floor.
	Name                   *string // Name.
	PostalCode             *string // Postal code (or zip code).
	Building               *string // Name of the building.
	Unit                   *string // Name or number of the unit (apartment, suite).
	Room                   *string // Name or number of the room.
	Seat                   *string // Seat (desk, cubicle, workstation).
	PlaceType              *string // A description of the type of place that is addressed.
	Community              *string // Postal community name.
	POBox                  *string // Post office box (P.O. box).
	AdditionalCode         *string // Additional code.
	Road                   *string // A primary road or street.
	RoadSection            *string // Road Section.
	RoadBranch             *string // Road branch.
	RoadSubBranch          *string // Road sub branch.
	RoadPreModifier        *string // Road pre-modifier.
	RoadPostModifier       *string // Road post-modifier.
	GPSCoords              *string // The GPS coordinates of the part.
}

// Location The location of a resource.
type Location struct {
	EmbeddedObject
	Ref           uint
	Info          *string        // This indicates the location of the resource.
	InfoFormat    *string        // This represents the format of the Info property.
	PostalAddress *PostalAddress `gorm:"ForeignKey:LocationRef"` // Postal address of the addressed resource.
	Placement     *Placement     `gorm:"ForeignKey:LocationRef"` // A place within the addressed location.
}
