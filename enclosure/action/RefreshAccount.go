package action

import (
	"promise/enclosure/context"
)

type RefreshAccount struct {
	Name                string
	MessageID           string
	Description         string
	ExpectedExecutionMs uint64
}

func (s *RefreshAccount) Execute(c *context.Base) {

}
