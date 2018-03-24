package constError

import (
	"fmt"
)

var (
	// ErrorResourceExist is an error.
	ErrorResourceExist = fmt.Errorf("resource exist")
	// ErrorResourceNotExist is an error.
	ErrorResourceNotExist = fmt.Errorf("resource does not exist")
	// ErrorConvertFilter is an error.
	ErrorConvertFilter = fmt.Errorf("can not convert filter")
	// ErrorIDFormat is an error.
	ErrorIDFormat = fmt.Errorf("ID format error")
	// ErrorDataConvert is an error.
	ErrorDataConvert = fmt.Errorf("data convert error.")
)