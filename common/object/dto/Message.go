package dto

import (
	m "promise/common/object/model"
	"time"
)

type Argument struct {
	Type  string `json:"Type"`
	Name  string `json:"Name"`
	Value string `json:"Value"`
}

func (this *Argument) Model() *m.Argument {
	ret := new(m.Argument)
	ret.Type = this.Type
	ret.Name = this.Name
	ret.Value = this.Value
	return ret
}

type Support struct {
	ID                string     `json:"ID"`
	Reason            string     `json:"Reason"`
	ReasonArguments   []Argument `json:"ReasonArguments"`
	Solution          string     `json:"Solution"`
	SolutionArguments []Argument `json:"SolutionArguments"`
}

func (this *Support) Model() *m.Support {
	ret := new(m.Support)
	ret.ReasonArguments = make([]m.Argument, 0)
	ret.SolutionArguments = make([]m.Argument, 0)
	ret.ID = this.ID
	ret.Reason = this.Reason
	for i, _ := range this.ReasonArguments {
		ret.ReasonArguments = append(ret.ReasonArguments, *this.ReasonArguments[i].Model())
	}
	for i, _ := range this.SolutionArguments {
		ret.SolutionArguments = append(ret.SolutionArguments, *this.SolutionArguments[i].Model())
	}
	return ret
}

type Message struct {
	ID          string `json:"ID"`
	Severity    string `json:"Severity"`
	Category    string `json:"Category"`
	CreateAt    time.Time
	Description string     `json:"Description"`
	Arguments   []Argument `json:"Arguments"`
	Supports    []Support  `json:"Supports"`
}

func NewArgument(m m.Argument) Argument {
	r := Argument{
		Type:  m.Type,
		Name:  m.Name,
		Value: m.Value,
	}
	return r
}

func NewSupport(m m.Support) Support {
	r := Support{
		Reason:   m.Reason,
		Solution: m.Solution,
	}
	r.ReasonArguments = make([]Argument, 0)
	r.SolutionArguments = make([]Argument, 0)
	for i, _ := range m.ReasonArguments {
		r.ReasonArguments = append(r.ReasonArguments, NewArgument(m.ReasonArguments[i]))
	}
	for i, _ := range m.SolutionArguments {
		r.SolutionArguments = append(r.SolutionArguments, NewArgument(m.SolutionArguments[i]))
	}
	return r
}

func (this *Message) Load(m *m.Message) {
	this.Arguments = make([]Argument, 0)
	this.Supports = make([]Support, 0)
	this.ID = m.ID
	this.Severity = m.Severity
	this.Category = m.Category
	this.Description = m.Description
	this.CreateAt = m.CreatedAt
	for i, _ := range m.Arguments {
		this.Arguments = append(this.Arguments, NewArgument(m.Arguments[i]))
	}
	for i, _ := range m.Supports {
		this.Supports = append(this.Supports, NewSupport(m.Supports[i]))
	}
}

func (this *Message) Model() *m.Message {
	ret := new(m.Message)
	ret.Arguments = make([]m.Argument, 0)
	ret.Supports = make([]m.Support, 0)
	ret.ID = this.ID
	ret.Severity = this.Severity
	ret.Category = this.Category
	ret.Description = this.Description
	ret.CreatedAt = this.CreateAt
	for i, _ := range this.Arguments {
		ret.Arguments = append(ret.Arguments, *this.Arguments[i].Model())
	}
	for i, _ := range this.Supports {
		ret.Supports = append(ret.Supports, *this.Supports[i].Model())
	}
	return ret
}

func MessagesToDto(messages []m.Message) []Message {
	ret := []Message{}
	for i, _ := range messages {
		m := Message{}
		m.Load(&messages[i])
		ret = append(ret, m)
	}
	return ret
}
