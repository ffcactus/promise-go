package dto

import (
	"promise/server/object/model"
)

type ResourceRef struct {
	Ref string `json:"$ref"`
}

type ResourceResponse struct {
	Name           *string `json:"Name,omitempty"`
	Description    *string `json:"Description,omitempty"`
	State          *string `json:"State,omitempty"`
	Health         *string `json:"Health,omitempty"`
	PhysicalState  *string `json:"PhysicalState,omitempty"`
	PhysicalHealth *string `json:"PhysicalHealth,omitempty"`
}

func (dto *ResourceResponse) LoadResourceResponse(m *model.Resource) {
	dto.Name = m.Name
	dto.Description = m.Description
	dto.State = m.State
	dto.Health = m.Health
	dto.PhysicalState = m.PhysicalState
	dto.PhysicalHealth = m.PhysicalHealth
}

type MemberResponse struct {
	// PageURI            *string  `json:"PageURI,omitempty"`
	MemberID       string  `json:"MemberID"`
	Name           *string `json:"Name,omitempty"`
	Description    *string `json:"Description,omitempty"`
	State          *string `json:"State,omitempty"`
	Health         *string `json:"Health,omitempty"`
	PhysicalState  *string `json:"PhysicalState,omitempty"`
	PhysicalHealth *string `json:"PhysicalHealth,omitempty"`
}

func (dto *MemberResponse) LoadMemberResponse(m *model.Member) {
	// dto.URI = m.URI
	dto.MemberID = m.MemberID
	dto.Name = m.Name
	dto.Description = m.Description
	dto.State = m.State
	dto.Health = m.Health
	dto.PhysicalState = m.PhysicalState
	dto.PhysicalHealth = m.PhysicalHealth
}

// The commom number properties of a resource.
type ProductInfoResponse struct {
	Model           *string `json:"Model,omitempty"`           // The model string for this product.
	Manufacturer    *string `json:"Manufacturer,omitempty"`    // The manufacturer string  of this product.
	SKU             *string `json:"SKU,omitempty"`             // The SKU string of this product.
	SerialNumber    *string `json:"SerialNumber,omitempty"`    // The serial number for this resource.
	PartNumber      *string `json:"PartNumber,omitempty"`      // The part number for this resource.
	SparePartNumber *string `json:"SparePartNumber,omitempty"` // The spare part number for this resource.
	AssetTag        *string `json:"AssetTag,omitempty"`        // The value of this property shall be an identifying string used to track the resource for inventory purposes.
}

func (dto *ProductInfoResponse) LoadProductInfoResponse(m *model.ProductInfo) {
	dto.Model = m.Model
	dto.Manufacturer = m.Manufacturer
	dto.SKU = m.SKU
	dto.SerialNumber = m.SerialNumber
	dto.PartNumber = m.PartNumber
	dto.SparePartNumber = m.SparePartNumber
	dto.AssetTag = m.AssetTag
}

type ThresholdResponse struct {
	UpperThresholdNonCritical *float64 `json:"UpperThresholdNonCritical"` // Above normal range.
	UpperThresholdCritical    *float64 `json:"UpperThresholdCritical"`    // Above normal range but not yet fatal.
	UpperThresholdFatal       *float64 `json:"UpperThresholdFatal"`       // Above normal range and is fatal.
	LowerThresholdNonCritical *float64 `json:"LowerThresholdNonCritical"` // Below normal range.
	LowerThresholdCritical    *float64 `json:"LowerThresholdCritical"`    // Below normal range but not yet fatal.
	LowerThresholdFatal       *float64 `json:"LowerThresholdFatal"`       // Below normal range and is fatal.
}

func (dto *ThresholdResponse) LoadThresholdResponse(m *model.Threshold) {
	dto.UpperThresholdNonCritical = m.UpperThresholdNonCritical
	dto.UpperThresholdCritical = m.UpperThresholdCritical
	dto.UpperThresholdFatal = m.UpperThresholdFatal
	dto.LowerThresholdNonCritical = m.LowerThresholdNonCritical
	dto.LowerThresholdCritical = m.LowerThresholdCritical
	dto.LowerThresholdFatal = m.LowerThresholdFatal
}

// The placement within the addressed location.
type Placement struct {
	Row             *string `json:"Row,omitempty"`             // Name of row.
	Rack            *string `json:"Rack,omitempty"`            // Name of a rack location within a row.
	RackOffsetUnits *string `json:"RackOffsetUnits,omitempty"` // The type of Rack Units in use.
	RackOffset      *int    `json:"RackOffset,omitempty"`      // Vertical location of the item in terms of RackOffsetUnits.
}

// The PostalAddress for a resource.
type PostalAddress struct {
	Country                *string `json:"Country,omitempty"`                // Country.
	Territory              *string `json:"Territory,omitempty"`              // A top-level subdivision within a country.
	District               *string `json:"District,omitempty"`               // A county, parish, gun (JP), or  district (IN).
	City                   *string `json:"City,omitempty"`                   // City, township, or shi (JP).
	Division               *string `json:"Division,omitempty"`               // City division, borough, dity district, ward, chou (JP).
	Neighborhood           *string `json:"Neighborhood,omitempty"`           // Neighborhood or block.
	LeadingStreetDirection *string `json:"LeadingStreetDirection,omitempty"` // A leading street direction.
	Street                 *string `json:"Street,omitempty"`                 // Street name.
	TrailingStreetSuffix   *string `json:"TrailingStreetSuffix,omitempty"`   // A trailing street suffix.
	StreetSuffix           *string `json:"StreetSuffix,omitempty"`           // Avenue, Platz, Street, Circle.
	HouseNumber            *int    `json:"HouseNumber,omitempty"`            // Numeric portion of house number.
	HouseNumberSuffix      *string `json:"HouseNumberSuffix,omitempty"`      // House number suffix.
	Landmark               *string `json:"Landmark,omitempty"`               // Landmark.
	Location               *string `json:"Location,omitempty"`               // Room designation or other additional info.
	Floor                  *string `json:"Floor,omitempty"`                  // Floor.
	Name                   *string `json:"Name,omitempty"`                   // Name.
	PostalCode             *string `json:"PostalCode,omitempty"`             // Postal code (or zip code).
	Building               *string `json:"Building,omitempty"`               // Name of the building.
	Unit                   *string `json:"Unit,omitempty"`                   // Name or number of the unit (apartment, suite).
	Room                   *string `json:"Room,omitempty"`                   // Name or number of the room.
	Seat                   *string `json:"Seat,omitempty"`                   // Seat (desk, cubicle, workstation).
	PlaceType              *string `json:"PlaceType,omitempty"`              // A description of the type of place that is addressed.
	Community              *string `json:"Community,omitempty"`              // Postal community name.
	POBox                  *string `json:"POBox,omitempty"`                  // Post office box (P.O. box).
	AdditionalCode         *string `json:"AdditionalCode,omitempty"`         // Additional code.
	Road                   *string `json:"Road,omitempty"`                   // A primary road or street.
	RoadSection            *string `json:"RoadSection,omitempty"`            // Road Section.
	RoadBranch             *string `json:"RoadBranch,omitempty"`             // Road branch.
	RoadSubBranch          *string `json:"RoadSubBranch,omitempty"`          // Road sub branch.
	RoadPreModifier        *string `json:"RoadPreModifier,omitempty"`        // Road pre-modifier.
	RoadPostModifier       *string `json:"RoadPostModifier,omitempty"`       // Road post-modifier.
	GPSCoords              *string `json:"GPSCoords,omitempty"`              // The GPS coordinates of the part.
}

// The location of a resource.
type Location struct {
	Info          *string        `json:"Info,omitempty"`          // This indicates the location of the resource.
	InfoFormat    *string        `json:"InfoFormat,omitempty"`    // This represents the format of the Info property.
	PostalAddress *PostalAddress `json:"PostalAddress,omitempty"` // Postal address of the addressed resource.
	Placement     *Placement     `json:"Placement,omitempty"`     // A place within the addressed location.
}

func (dto *Location) Load(m *model.Location) {
	dto.Info = m.Info
	dto.InfoFormat = m.InfoFormat
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

		dto.PostalAddress = postalAddress
	}
	if m.Placement != nil {
		placement := new(Placement)
		placement.Row = m.Placement.Row
		placement.Rack = m.Placement.Rack
		placement.RackOffsetUnits = m.Placement.RackOffsetUnits
		placement.RackOffset = m.Placement.RackOffset

		dto.Placement = placement
	}
}
