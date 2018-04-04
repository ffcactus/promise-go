package consterror

import (
	"fmt"
)

var (
	// ErrorAlreadyFree is an error.
	ErrorAlreadyFree = fmt.Errorf("IP already free")
	// ErrorNotInPool is an error.
	ErrorNotInPool = fmt.Errorf("IP not in the pool")
)
