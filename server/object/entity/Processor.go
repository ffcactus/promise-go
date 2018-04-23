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

// Load will load data from model.
func (e *Processor) Load(m *model.Processor) {
	updateResourceEntity(&(*e).EmbeddedResource, &(*m).Resource)
	updateProductInfoEntity(&(*e).ProductInfo, &(*m).ProductInfo)
	e.Socket = m.Socket
	e.ProcessorType = m.ProcessorType
	e.ProcessorArchitecture = m.ProcessorArchitecture
	e.InstructionSet = m.InstructionSet
	e.MaxSpeedMHz = m.MaxSpeedMHz
	e.TotalCores = m.TotalCores
	e.TotalThreads = m.TotalThreads
	e.ProcessorIDVendorID = m.ProcessorID.VendorID
	e.ProcessorIDIdentificationRegisters = m.ProcessorID.IdentificationRegisters
	e.ProcessorIDEffectiveFamily = m.ProcessorID.EffectiveFamily
	e.ProcessorIDEffectiveModel = m.ProcessorID.EffectiveModel
	e.ProcessorIDStep = m.ProcessorID.Step
	e.ProcessorIDMicrocodeInfo = m.ProcessorID.MicrocodeInfo
}
