package entity

import (
	"promise/server/object/model"
)

// Processor This is the schema definition for the Processor resource.  It represents the properties of a processor attached to a System.
type Processor struct {
	ServerRef string
	EmbeddedResource
	ProductInfo
	Socket                *string // The socket or location of the processor.
	ProcessorType         *string // The type of processor.
	ProcessorArchitecture *string // The architecture of the processor.
	InstructionSet        *string // The instruction set of the processor.
	MaxSpeedMHz           *int    // The maximum clock speed of the processor.
	TotalCores            *int    // The total number of cores contained in this processor.
	TotalThreads          *int    // The total number of execution threads supported by this processor.
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
}
