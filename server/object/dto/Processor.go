package dto

import (
	"promise/server/object/model"
)

// Processor This is the schema definition for the Processor resource.  It represents the properties of a processor attached to a System.
type Processor struct {
	ResourceResponse
	ProductInfoResponse
	Socket                *string `json:"Socket,omitempty"`                // The socket or location of the processor.
	ProcessorType         *string `json:"ProcessorType,omitempty"`         // The type of processor.
	ProcessorArchitecture *string `json:"ProcessorArchitecture,omitempty"` // The architecture of the processor.
	InstructionSet        *string `json:"InstructionSet,omitempty"`        // The instruction set of the processor.
	MaxSpeedMHz           *int    `json:"MaxSpeedMHz,omitempty"`           // The maximum clock speed of the processor.
	TotalCores            *int    `json:"TotalCores,omitempty"`            // The total number of cores contained in this processor.
	TotalThreads          *int    `json:"TotalThreads,omitempty"`          // The total number of execution threads supported by this processor.
}

// Load will load data from model.
func (dto *Processor) Load(m *model.Processor) {
	dto.LoadResourceResponse(&m.Resource)
	dto.LoadProductInfoResponse(&m.ProductInfo)
	dto.Socket = m.Socket
	dto.ProcessorType = m.ProcessorType
	dto.ProcessorArchitecture = m.ProcessorArchitecture
	dto.InstructionSet = m.InstructionSet
	dto.MaxSpeedMHz = m.MaxSpeedMHz
	dto.TotalCores = m.TotalCores
	dto.TotalThreads = m.TotalThreads
}
