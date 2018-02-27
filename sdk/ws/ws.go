package ws

import (
	"promise/common/app"
)

var (
	// WsServerRoot The root of the service.
	WsServerRoot = app.ProtocolScheme + app.Host + app.RootURL + "/ws/"
)