package dto

// GetProcessorResponse definition for the Processor resource.  It represents the properties of a processor attached to a System.
type GetProcessorResponse struct {
	Resource
	ProductInfo
	Socket                *string `json:"Socket"`                // The socket or location of the processor.
	ProcessorType         *string `json:"ProcessorType"`         // The type of processor.
	ProcessorArchitecture *string `json:"ProcessorArchitecture"` // The architecture of the processor.
	InstructionSet        *string `json:"InstructionSet"`        // The instruction set of the processor.
	MaxSpeedMHz           *int    `json:"MaxSpeedMHz"`           // The maximum clock speed of the processor.
	TotalCores            *int    `json:"TotalCores"`            // The total number of cores contained in this processor.
	TotalThreads          *int    `json:"TotalThreads"`          // The total number of execution threads supported by this processor.
}
