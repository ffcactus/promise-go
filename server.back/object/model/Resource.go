package model

// ResourceStatus This type describes the status and health of a resource and its children.
type ResourceStatus struct {
	State        *string
	HealthRollup *string // This property shall represent the HealthState of the resource and its dependent resources.
	Health       *string // This property shall represent the HealthState of the resource without consIDering its dependent resources. The values shall conform to those defined in the Redfish specification.
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

// MemberID This is the base type for addressable members of an array.
type MemberID struct {
	MemberID string // This is the IDentifier for the member within the collection.
}

// Resource The whole server information is build up by resource and member.
// Both Resource and Member are referable object, Member is used for array.
type Resource struct {
	URI            *string
	Name           *string
	Description    *string
	State          *string
	Health         *string
	OriginID       *string
	PhysicalState  *string
	PhysicalHealth *string
}

// Member The member object
type Member struct {
	URI            *string
	MemberID       string
	Name           *string
	Description    *string
	State          *string
	Health         *string
	OriginMemberID *string
	PhysicalState  *string
	PhysicalHealth *string
}

//
// type ResourceCollection struct {
// 	Description  string
// 	Name         string
// 	MembersCount int
// }

// ProductInfo The commom number properties of a resource.
type ProductInfo struct {
	Model           *string // The model string for this product.
	Manufacturer    *string // The manufacturer string  of this product.
	SKU             *string // The SKU string of this product.
	SerialNumber    *string // The serial number for this resource.
	PartNumber      *string // The part number for this resource.
	SparePartNumber *string // The spare part number for this resource.
	AssetTag        *string // The value of this property shall be an IDentifying string used to track the resource for inventory purposes.
}

// Placement The placement within the addressed location.
type Placement struct {
	Row             *string // Name of row.
	Rack            *string // Name of a rack location within a row.
	RackOffsetUnits *string // The type of Rack Units in use.
	RackOffset      *int    // Vertical location of the item in terms of RackOffsetUnits.
}

// PostalAddress The PostalAddress for a resource.
type PostalAddress struct {
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
	Info          *string        // This indicates the location of the resource.
	InfoFormat    *string        // This represents the format of the Info property.
	PostalAddress *PostalAddress // Postal address of the addressed resource.
	Placement     *Placement     // A place within the addressed location.
}

// Identifier This type describes any additional identifiers for a resource.
type Identifier struct {
	DurableName       *string // This indicates the world wide, persistent name of the resource.
	DurableNameFormat *string // This represents the format of the DurableName property.
}
