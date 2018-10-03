package action

import (
	"promise/enclosure/context"
)

type RefreshBlade struct {
	Name                string
	MessageID           string
	Description         string
	ExpectedExecutionMs uint64
}

func (s *RefreshBlade) Execute(c *context.Base) {

}
