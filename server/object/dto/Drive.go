package dto

import (
	"promise/server/object/model"
)

// Drive contains properties describing a single physical disk drive for any system, along with links to associated Volumes.
type Drive struct {
	ResourceResponse
	ProductInfoResponse
	StatusIndicator               *string    `json:"StatusIndicator"`               // The state of the status indicator, used to communicate status information about this drive.
	IndicatorLED                  *string    `json:"IndicatorLED"`                  // The state of the indicator LED, used to identify the drive.
	Revision                      *string    `json:"Revision"`                      // The revision of this Drive. This is typically the firmware/hardware version of the drive.
	CapacityBytes                 *int64     `json:"CapacityBytes"`                 // The size in bytes of this Drive.
	FailurePredicted              *bool      `json:"FailurePredicted"`              // Is this drive currently predicting a failure in the near future.
	Protocol                      *string    `json:"Protocol"`                      // The protocol this drive is using to communicate to the storage controller.
	MediaType                     *string    `json:"MediaType"`                     // The type of media contained in this drive.
	HotspareType                  *string    `json:"HotspareType"`                  // The type of hotspare this drive is currently serving as.
	CapableSpeedGbs               *int       `json:"CapableSpeedGbs"`               // The speed which this drive can communicate to a storage controller in ideal conditions in Gigabits per second.
	NegotiatedSpeedGbs            *int       `json:"NegotiatedSpeedGbs"`            // The speed which this drive is currently communicating to the storage controller in Gigabits per second.
	PredictedMediaLifeLeftPercent *int       `json:"PredictedMediaLifeLeftPercent"` // The percentage of reads and writes that are predicted to still be available for the media.
	Location                      []Location `json:"Location"`                      // The Location of the drive.
}

func (this *Drive) Load(m *model.Drive) {
	this.LoadResourceResponse(&m.Resource)
	this.LoadProductInfoResponse(&m.ProductInfo)
	this.StatusIndicator = m.StatusIndicator
	this.IndicatorLED = m.IndicatorLED
	this.Revision = m.Revision
	this.CapacityBytes = m.CapacityBytes
	this.FailurePredicted = m.FailurePredicted
	this.Protocol = m.Protocol
	this.MediaType = m.MediaType
	this.HotspareType = m.HotspareType
	this.CapableSpeedGbs = m.CapableSpeedGbs
	this.NegotiatedSpeedGbs = m.NegotiatedSpeedGbs
	this.PredictedMediaLifeLeftPercent = m.PredictedMediaLifeLeftPercent
	for i, _ := range m.Location {
		each := new(Location)
		each.Load(&m.Location[i])
		this.Location = append(this.Location, *each)
	}
}
