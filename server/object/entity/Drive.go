package entity

import (
	"promise/server/object/model"
)

// Drive contains properties describing a single physical disk drive for any system, along with links to associated Volumes.
type Drive struct {
	ServerRef string
	EmbeddedResource
	ProductInfo
	StatusIndicator               *string    // The state of the status indicator, used to communicate status information about this drive.
	IndicatorLED                  *string    // The state of the indicator LED, used to identify the drive.
	Revision                      *string    // The revision of this Drive. This is typically the firmware/hardware version of the drive.
	CapacityBytes                 *int64     // The size in bytes of this Drive.
	FailurePredicted              *bool      // Is this drive currently predicting a failure in the near future.
	Protocol                      *string    // The protocol this drive is using to communicate to the storage controller.
	MediaType                     *string    // The type of media contained in this drive.
	HotspareType                  *string    // The type of hotspare this drive is currently serving as.
	CapableSpeedGbs               *int       // The speed which this drive can communicate to a storage controller in ideal conditions in Gigabits per second.
	NegotiatedSpeedGbs            *int       // The speed which this drive is currently communicating to the storage controller in Gigabits per second.
	PredictedMediaLifeLeftPercent *int       // The percentage of reads and writes that are predicted to still be available for the media.
	Location                      []Location `gorm:"ForeignKey:Ref"` // The Location of the drive.
}

// ToModel will create a new model from entity.
func (e *Drive) ToModel() *model.Drive {
	m := new(model.Drive)
	createResourceModel(&e.EmbeddedResource, &m.Resource)
	createProductInfoModel(&e.ProductInfo, &m.ProductInfo)
	m.StatusIndicator = e.StatusIndicator
	m.IndicatorLED = e.IndicatorLED
	m.Revision = e.Revision
	m.CapacityBytes = e.CapacityBytes
	m.FailurePredicted = e.FailurePredicted
	m.Protocol = e.Protocol
	m.MediaType = e.MediaType
	m.HotspareType = e.HotspareType
	m.CapableSpeedGbs = e.CapableSpeedGbs
	m.NegotiatedSpeedGbs = e.NegotiatedSpeedGbs
	m.PredictedMediaLifeLeftPercent = e.PredictedMediaLifeLeftPercent
	for i := range e.Location {
		locationM := new(model.Location)
		createLocationModel(&e.Location[i], locationM)
		m.Location = append(m.Location, *locationM)
	}
	return m
}

// Load will load data from model.
func (e *Drive) Load(m *model.Drive) {
	updateResourceEntity(&e.EmbeddedResource, &m.Resource)
	updateProductInfoEntity(&e.ProductInfo, &m.ProductInfo)
	e.StatusIndicator = m.StatusIndicator
	e.IndicatorLED = m.IndicatorLED
	e.Revision = m.Revision
	e.CapacityBytes = m.CapacityBytes
	e.FailurePredicted = m.FailurePredicted
	e.Protocol = m.Protocol
	e.MediaType = m.MediaType
	e.HotspareType = m.HotspareType
	e.CapableSpeedGbs = m.CapableSpeedGbs
	e.NegotiatedSpeedGbs = m.NegotiatedSpeedGbs
	e.PredictedMediaLifeLeftPercent = m.PredictedMediaLifeLeftPercent
	for _, v := range m.Location {
		locationE := Location{}
		locationE.Load(&v)
		e.Location = append(e.Location, locationE)
	}
}
