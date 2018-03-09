package dto

import (
	"promise/server/object/model"
)

// This type describes the Identification information for a processor.
type ProcessorID struct {
	VendorID                *string `json:"VendorID,omitempty"`        // The Vendor Identification for this processor.
	IdentificationRegisters *string `json:"IdentificationRegisters"`   // The contents of the Identification Registers (CPUID) for this processor.
	EffectiveFamily         *string `json:"EffectiveFamily,omitempty"` // The effective Family for this processor.
	EffectiveModel          *string `json:"EffectiveModel,omitempty"`  // The effective Model for this processor.
	Step                    *string `json:"Step,omitempty"`            // The Step value for this processor.
	MicrocodeInfo           *string `json:"MicrocodeInfo,omitempty"`   // The Microcode Information for this processor.
}

// This is the schema definition for the Processor resource.  It represents the properties of a processor attached to a System.
type Processor struct {
	ResourceResponse
	ProductInfoResponse
	Socket                *int         `json:"Socket,omitempty"`                // The socket or location of the processor.
	ProcessorType         *string      `json:"ProcessorType,omitempty"`         // The type of processor.
	ProcessorArchitecture *string      `json:"ProcessorArchitecture,omitempty"` // The architecture of the processor.
	InstructionSet        *string      `json:"InstructionSet,omitempty"`        // The instruction set of the processor.
	ProcessorID           *ProcessorID `json:"ProcessorID,omitempty"`           // Identification information for this processor.
	MaxSpeedMHz           *int         `json:"MaxSpeedMHz,omitempty"`           // The maximum clock speed of the processor.
	TotalCores            *int         `json:"TotalCores,omitempty"`            // The total number of cores contained in this processor.
	TotalThreads          *int         `json:"TotalThreads,omitempty"`          // The total number of execution threads supported by this processor.
}

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
	if m.ProcessorID != nil {
		dto.ProcessorID = new(ProcessorID)
		dto.ProcessorID.VendorID = m.ProcessorID.VendorID
		dto.ProcessorID.IdentificationRegisters = m.ProcessorID.IdentificationRegisters
		dto.ProcessorID.EffectiveFamily = m.ProcessorID.EffectiveFamily
		dto.ProcessorID.EffectiveModel = m.ProcessorID.EffectiveModel
		dto.ProcessorID.Step = m.ProcessorID.Step
		dto.ProcessorID.MicrocodeInfo = m.ProcessorID.MicrocodeInfo
	}
}
