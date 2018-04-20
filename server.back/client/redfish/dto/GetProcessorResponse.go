package dto

// This type describes the Identification information for a processor.
type ProcessorID struct {
	VendorID                *string `json:"VendorID"`                // The Vendor Identification for this processor.
	IdentificationRegisters *string `json:"IdentificationRegisters"` // The contents of the Identification Registers (CPUID) for this processor.
	EffectiveFamily         *string `json:"EffectiveFamily"`         // The effective Family for this processor.
	EffectiveModel          *string `json:"EffectiveModel"`          // The effective Model for this processor.
	Step                    *string `json:"Step"`                    // The Step value for this processor.
	MicrocodeInfo           *string `json:"MicrocodeInfo"`           // The Microcode Information for this processor.
}

// This is the schema definition for the Processor resource.  It represents the properties of a processor attached to a System.
type GetProcessorResponse struct {
	Resource
	ProductInfo
	Socket                *int         `json:"Socket"`                // The socket or location of the processor.
	ProcessorType         *string      `json:"ProcessorType"`         // The type of processor.
	ProcessorArchitecture *string      `json:"ProcessorArchitecture"` // The architecture of the processor.
	InstructionSet        *string      `json:"InstructionSet"`        // The instruction set of the processor.
	MaxSpeedMHz           *int         `json:"MaxSpeedMHz"`           // The maximum clock speed of the processor.
	TotalCores            *int         `json:"TotalCores"`            // The total number of cores contained in this processor.
	TotalThreads          *int         `json:"TotalThreads"`          // The total number of execution threads supported by this processor.
	ProcessorID           *ProcessorID `json:"ProcessorID"`           // Identification information for this processor.
}
