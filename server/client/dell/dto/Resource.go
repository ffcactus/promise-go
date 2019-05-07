package dto

type Status struct {
	State  *string `json:"State"`
	Health *string `json:"Health"`
}

type Resource struct {
	OdataID     *string `json:"@odata.id"`
	Id          *string `json:"Id"`
	MemberId    *string `json:"MemberId"`
	Name        *string `json:"Name"`
	Description *string `json:"Description"`
	Status      *Status `json:"Status"`
}

type Member struct {
	OdataID     *string `json:"@odata.id"`
	MemberId    *string `json:"MemberId"`
	Name        *string `json:"Name"`
	Description *string `json:"Description"`
	Status      *Status `json:"Status"`
}

type UUID struct {
	PhysicalUUID string `json:"UUID"`
}

type ProductInfo struct {
	Model           *string `json:"Model"`
	Manufacturer    *string `json:"Manufacturer"`
	SKU             *string `json:"SKU"`
	SerialNumber    *string `json:"SerialNumber"`
	PartNumber      *string `json:"PartNumber"`
	SparePartNumber *string `json:"SparePartNumber"`
	AssetTag        *string `json:"AssetTag"`
}

type Threshold struct {
	UpperThresholdNonCritical *float64 `json:"UpperThresholdNonCritical"` // Above normal range.
	UpperThresholdCritical    *float64 `json:"UpperThresholdCritical"`    // Above normal range but not yet fatal.
	UpperThresholdFatal       *float64 `json:"UpperThresholdFatal"`       // Above normal range and is fatal.
	LowerThresholdNonCritical *float64 `json:"LowerThresholdNonCritical"` // Below normal range.
	LowerThresholdCritical    *float64 `json:"LowerThresholdCritical"`    // Below normal range but not yet fatal.
	LowerThresholdFatal       *float64 `json:"LowerThresholdFatal"`       // Below normal range and is fatal.
}

type OdataID struct {
	Id *string `json:"@odata.id"`
}

type ResourceRef struct {
	OdataId string `json:"@odata.id"`
}

// The placement within the addressed location.
type Placement struct {
	Row             *string `json:"Row"`             // Name of row.
	Rack            *string `json:"Rack"`            // Name of a rack location within a row.
	RackOffsetUnits *string `json:"RackOffsetUnits"` // The type of Rack Units in use.
	RackOffset      *int    `json:"RackOffset"`      // Vertical location of the item in terms of RackOffsetUnits.
}

// The PostalAddress for a resource.
type PostalAddress struct {
	Country                *string `json:"Country"`                // Country.
	Territory              *string `json:"Territory"`              // A top-level subdivision within a country.
	District               *string `json:"District"`               // A county, parish, gun (JP), or  district (IN).
	City                   *string `json:"City"`                   // City, township, or shi (JP).
	Division               *string `json:"Division"`               // City division, borough, dity district, ward, chou (JP).
	Neighborhood           *string `json:"Neighborhood"`           // Neighborhood or block.
	LeadingStreetDirection *string `json:"LeadingStreetDirection"` // A leading street direction.
	Street                 *string `json:"Street"`                 // Street name.
	TrailingStreetSuffix   *string `json:"TrailingStreetSuffix"`   // A trailing street suffix.
	StreetSuffix           *string `json:"StreetSuffix"`           // Avenue, Platz, Street, Circle.
	HouseNumber            *int    `json:"HouseNumber"`            // Numeric portion of house number.
	HouseNumberSuffix      *string `json:"HouseNumberSuffix"`      // House number suffix.
	Landmark               *string `json:"Landmark"`               // Landmark.
	Location               *string `json:"Location"`               // Room designation or other additional info.
	Floor                  *string `json:"Floor"`                  // Floor.
	Name                   *string `json:"Name"`                   // Name.
	PostalCode             *string `json:"PostalCode"`             // Postal code (or zip code).
	Building               *string `json:"Building"`               // Name of the building.
	Unit                   *string `json:"Unit"`                   // Name or number of the unit (apartment, suite).
	Room                   *string `json:"Room"`                   // Name or number of the room.
	Seat                   *string `json:"Seat"`                   // Seat (desk, cubicle, workstation).
	PlaceType              *string `json:"PlaceType"`              // A description of the type of place that is addressed.
	Community              *string `json:"Community"`              // Postal community name.
	POBox                  *string `json:"POBox"`                  // Post office box (P.O. box).
	AdditionalCode         *string `json:"AdditionalCode"`         // Additional code.
	Road                   *string `json:"Road"`                   // A primary road or street.
	RoadSection            *string `json:"RoadSection"`            // Road Section.
	RoadBranch             *string `json:"RoadBranch"`             // Road branch.
	RoadSubBranch          *string `json:"RoadSubBranch"`          // Road sub branch.
	RoadPreModifier        *string `json:"RoadPreModifier"`        // Road pre-modifier.
	RoadPostModifier       *string `json:"RoadPostModifier"`       // Road post-modifier.
	GPSCoords              *string `json:"GPSCoords"`              // The GPS coordinates of the part.
}

// The location of a resource.
type Location struct {
	Info          *string        `json:"Info"`          // This indicates the location of the resource.
	InfoFormat    *string        `json:"InfoFormat"`    // This represents the format of the Info property.
	PostalAddress *PostalAddress `json:"PostalAddress"` // Postal address of the addressed resource.
	Placement     *Placement     `json:"Placement"`     // A place within the addressed location.
}

// This type describes any additional identifiers for a resource.
type Identifier struct {
	DurableName       *string `json:"DurableName"`       // This indicates the world wide, persistent name of the resource.
	DurableNameFormat *string `json:"DurableNameFormat"` // This represents the format of the DurableName property.
}
