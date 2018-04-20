package entity

import (
	"promise/server/object/model"
)

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

// ToModel will create a new model from entity.
func (e *Processor) ToModel() *model.Processor {
	m := model.Processor{}
	createResourceModel(&e.EmbeddedResource, &m.Resource)
	createProductInfoModel(&e.ProductInfo, &m.ProductInfo)
	m.Socket = e.Socket
	m.ProcessorType = e.ProcessorType
	m.ProcessorArchitecture = e.ProcessorArchitecture
	m.InstructionSet = e.InstructionSet
	m.MaxSpeedMHz = e.MaxSpeedMHz
	m.TotalCores = e.TotalCores
	m.TotalThreads = e.TotalThreads
	m.ProcessorID = new(model.ProcessorID)
	if e.ProcessorIDVendorID != nil ||
		e.ProcessorIDIdentificationRegisters != nil ||
		e.ProcessorIDEffectiveFamily != nil ||
		e.ProcessorIDEffectiveModel != nil ||
		e.ProcessorIDStep != nil ||
		e.ProcessorIDMicrocodeInfo != nil {
		m.ProcessorID = new(model.ProcessorID)
		m.ProcessorID.VendorID = e.ProcessorIDVendorID
		m.ProcessorID.IdentificationRegisters = e.ProcessorIDIdentificationRegisters
		m.ProcessorID.EffectiveFamily = e.ProcessorIDEffectiveFamily
		m.ProcessorID.EffectiveModel = e.ProcessorIDEffectiveModel
		m.ProcessorID.Step = e.ProcessorIDStep
		m.ProcessorID.MicrocodeInfo = e.ProcessorIDMicrocodeInfo
	}
	return &m
}
