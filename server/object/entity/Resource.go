package entity

import (
	"promise/server/object/model"
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
	ID uint64 `gorm:"primary_key"`
	Status
	URI         *string
	OriginID    *string
	Name        *string
	Description *string
}

// EmbeddedObject The EmbeddedObject is sub structure in Resource or EmbeddedResource.
type EmbeddedObject struct {
	ID uint64 `gorm:"primary_key"`
}

// Member Member is a object in a collection. Each member should have a member ID.
type Member struct {
	ID uint64 `gorm:"primary_key"`
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
	LocationRef uint64
	EmbeddedObject
	Row             *string // Name of row.
	Rack            *string // Name of a rack location within a row.
	RackOffsetUnits *string // The type of Rack Units in use.
	RackOffset      *int    // Vertical location of the item in terms of RackOffsetUnits.
}

// PostalAddress The PostalAddress for a resource.
type PostalAddress struct {
	LocationRef uint64
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

// Load will load data from model.
func (e *Location) Load(m *model.Location) {
	e.Info = m.Info
	e.InfoFormat = m.InfoFormat
	if m.PostalAddress != nil {
		postalAddress := new(PostalAddress)
		postalAddress.Country = m.PostalAddress.Country
		postalAddress.Territory = m.PostalAddress.Territory
		postalAddress.District = m.PostalAddress.District
		postalAddress.City = m.PostalAddress.City
		postalAddress.Division = m.PostalAddress.Division
		postalAddress.Neighborhood = m.PostalAddress.Neighborhood
		postalAddress.LeadingStreetDirection = m.PostalAddress.LeadingStreetDirection
		postalAddress.Street = m.PostalAddress.Street
		postalAddress.TrailingStreetSuffix = m.PostalAddress.TrailingStreetSuffix
		postalAddress.StreetSuffix = m.PostalAddress.StreetSuffix
		postalAddress.HouseNumber = m.PostalAddress.HouseNumber
		postalAddress.HouseNumberSuffix = m.PostalAddress.HouseNumberSuffix
		postalAddress.Landmark = m.PostalAddress.Landmark
		postalAddress.Location = m.PostalAddress.Location
		postalAddress.Floor = m.PostalAddress.Floor
		postalAddress.Name = m.PostalAddress.Name
		postalAddress.PostalCode = m.PostalAddress.PostalCode
		postalAddress.Building = m.PostalAddress.Building
		postalAddress.Unit = m.PostalAddress.Unit
		postalAddress.Room = m.PostalAddress.Room
		postalAddress.Seat = m.PostalAddress.Seat
		postalAddress.PlaceType = m.PostalAddress.PlaceType
		postalAddress.Community = m.PostalAddress.Community
		postalAddress.POBox = m.PostalAddress.POBox
		postalAddress.AdditionalCode = m.PostalAddress.AdditionalCode
		postalAddress.Road = m.PostalAddress.Road
		postalAddress.RoadSection = m.PostalAddress.RoadSection
		postalAddress.RoadBranch = m.PostalAddress.RoadBranch
		postalAddress.RoadSubBranch = m.PostalAddress.RoadSubBranch
		postalAddress.RoadPreModifier = m.PostalAddress.RoadPreModifier
		postalAddress.RoadPostModifier = m.PostalAddress.RoadPostModifier
		postalAddress.GPSCoords = m.PostalAddress.GPSCoords

		e.PostalAddress = postalAddress
	}
	if m.Placement != nil {
		placement := new(Placement)
		placement.Row = m.Placement.Row
		placement.Rack = m.Placement.Rack
		placement.RackOffsetUnits = m.Placement.RackOffsetUnits
		placement.RackOffset = m.Placement.RackOffset

		e.Placement = placement
	}
}

func createResourceModel(e *EmbeddedResource, m *model.Resource) {
	m.URI = e.URI
	m.OriginID = e.OriginID
	m.Name = e.Name
	m.Description = e.Description
	m.State = e.State
	m.Health = e.Health
	m.PhysicalState = e.PhysicalState
	m.PhysicalHealth = e.PhysicalHealth
}

func createMemberModel(e *Member, m *model.Member) {
	m.URI = e.URI
	m.OriginMemberID = e.OriginMemberID
	m.MemberID = *e.OriginMemberID
	m.Name = e.Name
	m.Description = e.Description
	m.State = e.State
	m.Health = e.Health
	m.PhysicalState = e.PhysicalState
	m.PhysicalHealth = e.PhysicalHealth
}

func createProductInfoModel(e *ProductInfo, m *model.ProductInfo) {
	m.Model = e.Model
	m.Manufacturer = e.Manufacturer
	m.SKU = e.SKU
	m.SerialNumber = e.SerialNumber
	m.SparePartNumber = e.SparePartNumber
	m.PartNumber = e.PartNumber
	m.AssetTag = e.AssetTag
}

func createThresholdModel(e *Threshold, m *model.Threshold) {
	m.UpperThresholdNonCritical = e.UpperThresholdNonCritical
	m.UpperThresholdCritical = e.UpperThresholdCritical
	m.UpperThresholdFatal = e.UpperThresholdFatal
	m.LowerThresholdNonCritical = e.LowerThresholdNonCritical
	m.LowerThresholdCritical = e.LowerThresholdCritical
	m.LowerThresholdFatal = e.LowerThresholdFatal
}

func createLocationModel(e *Location, m *model.Location) {
	m.Info = e.Info
	m.InfoFormat = e.InfoFormat
	if e.PostalAddress != nil {
		postalAddress := new(model.PostalAddress)
		postalAddress.Country = e.PostalAddress.Country
		postalAddress.Territory = e.PostalAddress.Territory
		postalAddress.District = e.PostalAddress.District
		postalAddress.City = e.PostalAddress.City
		postalAddress.Division = e.PostalAddress.Division
		postalAddress.Neighborhood = e.PostalAddress.Neighborhood
		postalAddress.LeadingStreetDirection = e.PostalAddress.LeadingStreetDirection
		postalAddress.Street = e.PostalAddress.Street
		postalAddress.TrailingStreetSuffix = e.PostalAddress.TrailingStreetSuffix
		postalAddress.StreetSuffix = e.PostalAddress.StreetSuffix
		postalAddress.HouseNumber = e.PostalAddress.HouseNumber
		postalAddress.HouseNumberSuffix = e.PostalAddress.HouseNumberSuffix
		postalAddress.Landmark = e.PostalAddress.Landmark
		postalAddress.Location = e.PostalAddress.Location
		postalAddress.Floor = e.PostalAddress.Floor
		postalAddress.Name = e.PostalAddress.Name
		postalAddress.PostalCode = e.PostalAddress.PostalCode
		postalAddress.Building = e.PostalAddress.Building
		postalAddress.Unit = e.PostalAddress.Unit
		postalAddress.Room = e.PostalAddress.Room
		postalAddress.Seat = e.PostalAddress.Seat
		postalAddress.PlaceType = e.PostalAddress.PlaceType
		postalAddress.Community = e.PostalAddress.Community
		postalAddress.POBox = e.PostalAddress.POBox
		postalAddress.AdditionalCode = e.PostalAddress.AdditionalCode
		postalAddress.Road = e.PostalAddress.Road
		postalAddress.RoadSection = e.PostalAddress.RoadSection
		postalAddress.RoadBranch = e.PostalAddress.RoadBranch
		postalAddress.RoadSubBranch = e.PostalAddress.RoadSubBranch
		postalAddress.RoadPreModifier = e.PostalAddress.RoadPreModifier
		postalAddress.RoadPostModifier = e.PostalAddress.RoadPostModifier
		postalAddress.GPSCoords = e.PostalAddress.GPSCoords

		m.PostalAddress = postalAddress
	}
	if e.Placement != nil {
		placement := new(model.Placement)
		placement.Row = e.Placement.Row
		placement.Rack = e.Placement.Rack
		placement.RackOffsetUnits = e.Placement.RackOffsetUnits
		placement.RackOffset = e.Placement.RackOffset

		m.Placement = placement
	}
}

func updateResourceEntity(e *EmbeddedResource, m *model.Resource) {
	// ID won't be updated.
	e.URI = m.URI
	e.OriginID = m.OriginID
	e.Name = m.Name
	e.Description = m.Description
	e.State = m.State
	e.Health = m.Health
	e.PhysicalState = m.PhysicalState
	e.PhysicalHealth = m.PhysicalHealth
}

func updateMemberEntity(e *Member, m *model.Member) {
	e.URI = m.URI
	e.OriginMemberID = m.OriginMemberID
	e.Name = m.Name
	e.Description = m.Description
	e.State = m.State
	e.Health = m.Health
	e.PhysicalState = m.PhysicalState
	e.PhysicalHealth = m.PhysicalHealth
}

func updateProductInfoEntity(e *ProductInfo, m *model.ProductInfo) {
	e.Model = m.Model
	e.Manufacturer = m.Manufacturer
	e.SerialNumber = m.SerialNumber
	e.PartNumber = m.PartNumber
	e.SparePartNumber = m.SparePartNumber
	e.SKU = m.SKU
	e.AssetTag = m.AssetTag
}

func updateThresholdEntity(e *Threshold, m *model.Threshold) {
	e.UpperThresholdNonCritical = m.UpperThresholdNonCritical
	e.UpperThresholdCritical = m.UpperThresholdCritical
	e.UpperThresholdFatal = m.UpperThresholdFatal
	e.LowerThresholdNonCritical = m.LowerThresholdNonCritical
	e.LowerThresholdCritical = m.LowerThresholdCritical
	e.LowerThresholdFatal = m.LowerThresholdFatal
}
