package model

// ProcessorID This type describes the Identification information for a processor.
type ProcessorID struct {
	VendorID                *string // The Vendor Identification for this processor.
	IdentificationRegisters *string // The contents of the Identification Registers (CPUID) for this processor.
	EffectiveFamily         *string // The effective Family for this processor.
	EffectiveModel          *string // The effective Model for this processor.
	Step                    *string // The Step value for this processor.
	MicrocodeInfo           *string // The Microcode Information for this processor.
}

// Processor This is the schema definition for the Processor resource.  It represents the properties of a processor attached to a System.
type Processor struct {
	Resource
	ProductInfo
	Socket                *int         // The socket or location of the processor.
	ProcessorType         *string      // The type of processor.
	ProcessorArchitecture *string      // The architecture of the processor.
	InstructionSet        *string      // The instruction set of the processor.
	ProcessorID           *ProcessorID // Identification information for this processor.
	MaxSpeedMHz           *int         // The maximum clock speed of the processor.
	TotalCores            *int         // The total number of cores contained in this processor.
	TotalThreads          *int         // The total number of execution threads supported by this processor.
}
