package base

import (
	"fmt"
)

var (
	// ErrorInvalidURLParameter is an error.
	ErrorInvalidURLParameter = fmt.Errorf("invalid URL parameter")
	// ErrorResourceExist is an error.
	ErrorResourceExist = fmt.Errorf("resource exist")
	// ErrorResourceNotExist is an error.
	ErrorResourceNotExist = fmt.Errorf("resource does not exist")
	// ErrorConvertFilter is an error.
	ErrorConvertFilter = fmt.Errorf("can not convert filter")
	// ErrorIDFormat is an error.
	ErrorIDFormat = fmt.Errorf("ID format error")
	// ErrorDataConvert is an error.
	ErrorDataConvert = fmt.Errorf("data convert error")
	// ErrorUnknownFilterName is an error.
	ErrorUnknownFilterName = fmt.Errorf("unknown filter name")
	// ErrorUnknownPropertyValue is an error.
	ErrorUnknownPropertyValue = fmt.Errorf("unknown property value")
	// ErrorTransaction is an error.
	ErrorTransaction = fmt.Errorf("transaction error")
	// ErrorIPv4NotAllocated is an error.
	ErrorIPv4NotAllocated = fmt.Errorf("IP has not been allocated before")
	// ErrorIPv4NotInPool is an error.
	ErrorIPv4NotInPool = fmt.Errorf("IP not in the pool")
	// ErrorIPv4RangeEndAddress is an error.
	ErrorIPv4RangeEndAddress = fmt.Errorf("end address should equal or big then start address")
	// ErrorIPv4RangeSize is an error.
	ErrorIPv4RangeSize = fmt.Errorf("the number addresses in range should not more than 256")
	// ErrorIPv4RangeCount is an error.
	ErrorIPv4RangeCount = fmt.Errorf("IPv4 pool should contain one range at least")
)
