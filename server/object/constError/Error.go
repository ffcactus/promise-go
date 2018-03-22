package constError

import (
	"fmt"
)

var (
	// ErrorDeleteDefaultServerGroup is an error.
	ErrorDeleteDefaultServerGroup = fmt.Errorf("can not delete default servergroup")
	// ErrorDeleteDefaultServerServerGroup is an error.
	ErrorDeleteDefaultServerServerGroup = fmt.Errorf("can not delete default server-servergroup")
)