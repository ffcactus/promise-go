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
)
