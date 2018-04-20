package model

// Drive contains properties describing a single physical disk drive for any system, along with links to associated Volumes.
type Drive struct {
	Resource
	ProductInfo
	StatusIndicator               *string    // The state of the status indicator, used to communicate status information about this drive.
	IndicatorLED                  *string    // The state of the indicator LED, used to identify the drive.
	Revision                      *string    // The revision of this Drive. This is typically the firmware/hardware version of the drive.
	CapacityBytes                 *int64     // The size in bytes of this Drive.
	FailurePredicted              *bool      // Is this drive currently predicting a failure in the near future.
	Protocol                      *string    // The protocol this drive is using to communicate to the storage controller.
	MediaType                     *string    // The type of media contained in this drive.
	Location                      []Location // The Location of the drive.
	HotspareType                  *string    // The type of hotspare this drive is currently serving as.
	CapableSpeedGbs               *int       // The speed which this drive can communicate to a storage controller in ideal conditions in Gigabits per second.
	NegotiatedSpeedGbs            *int       // The speed which this drive is currently communicating to the storage controller in Gigabits per second.
	PredictedMediaLifeLeftPercent *int       // The percentage of reads and writes that are predicted to still be available for the media.
}
