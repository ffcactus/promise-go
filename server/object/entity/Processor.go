package entity

// Processor This is the schema definition for the Processor resource.  It represents the properties of a processor attached to a System.
type Processor struct {
	ServerRef string
	EmbeddedResource
	ProductInfo
	Socket                             *int    // The socket or location of the processor.
	ProcessorType                      *string // The type of processor.
	ProcessorArchitecture              *string // The architecture of the processor.
	InstructionSet                     *string // The instruction set of the processor.
	MaxSpeedMHz                        *int    // The maximum clock speed of the processor.
	TotalCores                         *int    // The total number of cores contained in this processor.
	TotalThreads                       *int    // The total number of execution threads supported by this processor.
	ProcessorIDVendorID                *string // The Vendor Identification for this processor.
	ProcessorIDIdentificationRegisters *string // The contents of the Identification Registers (CPUID) for this processor.
	ProcessorIDEffectiveFamily         *string // The effective Family for this processor.
	ProcessorIDEffectiveModel          *string // The effective Model for this processor.
	ProcessorIDStep                    *string // The Step value for this processor.
	ProcessorIDMicrocodeInfo           *string // The Microcode Information for this processor.
}
