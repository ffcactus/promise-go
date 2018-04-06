package consterror

import (
	"fmt"
)

var (
	// ErrorNotAllocated is an error.
	ErrorNotAllocated = fmt.Errorf("IP has not been allocated before")
	// ErrorNotInPool is an error.
	ErrorNotInPool = fmt.Errorf("IP not in the pool")
	// ErrorRangeEndAddress is an error.
	ErrorRangeEndAddress = fmt.Errorf("end address should equal or big then start address")
	// ErrorRangeSize is an error.
	ErrorRangeSize = fmt.Errorf("the number addresses in range should not more than 256")
	// ErrorRangeCount is an error.
	ErrorRangeCount = fmt.Errorf("IPv4 pool should contain one range at least")
)
