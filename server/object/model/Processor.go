package model

// Processor This is the schema definition for the Processor resource.  It represents the properties of a processor attached to a System.
type Processor struct {
	Resource
	ProductInfo
	Socket                *string // The socket or location of the processor.
	ProcessorType         *string // The type of processor.
	ProcessorArchitecture *string // The architecture of the processor.
	InstructionSet        *string // The instruction set of the processor.
	MaxSpeedMHz           *int    // The maximum clock speed of the processor.
	TotalCores            *int    // The total number of cores contained in this processor.
	TotalThreads          *int    // The total number of execution threads supported by this processor.
}
