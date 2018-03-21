package dto

import (
	"promise/common/object/message"
	"time"
)

// Argument is DTO.
type Argument struct {
	Type  string `json:"Type"`
	Name  string `json:"Name"`
	Value string `json:"Value"`
}

// Model convert to model.
func (dto *Argument) Model() *message.Argument {
	ret := new(message.Argument)
	ret.Type = dto.Type
	ret.Name = dto.Name
	ret.Value = dto.Value
	return ret
}

// Support is the DTO.
type Support struct {
	ID                string     `json:"ID"`
	Reason            string     `json:"Reason"`
	ReasonArguments   []Argument `json:"ReasonArguments"`
	Solution          string     `json:"Solution"`
	SolutionArguments []Argument `json:"SolutionArguments"`
}

// Model convert to model.
func (dto *Support) Model() *message.Support {
	ret := new(message.Support)
	ret.ReasonArguments = make([]message.Argument, 0)
	ret.SolutionArguments = make([]message.Argument, 0)
	ret.ID = dto.ID
	ret.Reason = dto.Reason
	for i := range dto.ReasonArguments {
		ret.ReasonArguments = append(ret.ReasonArguments, *dto.ReasonArguments[i].Model())
	}
	for i := range dto.SolutionArguments {
		ret.SolutionArguments = append(ret.SolutionArguments, *dto.SolutionArguments[i].Model())
	}
	return ret
}

// Message is the DTO.
type Message struct {
	ID          string `json:"ID"`
	Severity    string `json:"Severity"`
	Category    string `json:"Category"`
	CreateAt    time.Time
	Description string     `json:"Description"`
	Arguments   []Argument `json:"Arguments"`
	Supports    []Support  `json:"Supports"`
}

// NewArgument create a default Argument.
func NewArgument(m message.Argument) Argument {
	r := Argument{
		Type:  m.Type,
		Name:  m.Name,
		Value: m.Value,
	}
	return r
}

// NewSupport create a default Support.
func NewSupport(m message.Support) Support {
	r := Support{
		Reason:   m.Reason,
		Solution: m.Solution,
	}
	r.ReasonArguments = make([]Argument, 0)
	r.SolutionArguments = make([]Argument, 0)
	for i := range m.ReasonArguments {
		r.ReasonArguments = append(r.ReasonArguments, NewArgument(m.ReasonArguments[i]))
	}
	for i := range m.SolutionArguments {
		r.SolutionArguments = append(r.SolutionArguments, NewArgument(m.SolutionArguments[i]))
	}
	return r
}

// Load will load message from model.
func (dto *Message) Load(m *message.Message) {
	dto.Arguments = make([]Argument, 0)
	dto.Supports = make([]Support, 0)
	dto.ID = m.ID
	dto.Severity = m.Severity
	dto.Category = m.Category
	dto.Description = m.Description
	dto.CreateAt = m.CreatedAt
	for i := range m.Arguments {
		dto.Arguments = append(dto.Arguments, NewArgument(m.Arguments[i]))
	}
	for i := range m.Supports {
		dto.Supports = append(dto.Supports, NewSupport(m.Supports[i]))
	}
}

// Model will convert DTO to model.
func (dto *Message) Model() *message.Message {
	ret := new(message.Message)
	ret.Arguments = make([]message.Argument, 0)
	ret.Supports = make([]message.Support, 0)
	ret.ID = dto.ID
	ret.Severity = dto.Severity
	ret.Category = dto.Category
	ret.Description = dto.Description
	ret.CreatedAt = dto.CreateAt
	for i := range dto.Arguments {
		ret.Arguments = append(ret.Arguments, *dto.Arguments[i].Model())
	}
	for i := range dto.Supports {
		ret.Supports = append(ret.Supports, *dto.Supports[i].Model())
	}
	return ret
}

// MessagesToDto convert messages to DTO.
func MessagesToDto(messages []message.Message) []Message {
	ret := []Message{}
	for i := range messages {
		m := Message{}
		m.Load(&messages[i])
		ret = append(ret, m)
	}
	return ret
}
