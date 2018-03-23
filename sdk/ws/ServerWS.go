package ws

import (
	"net/http"
	"promise/server/object/model"
)

var (
	// WsServerRoot The root of the service.
	WsServerRoot = app.ProtocolScheme + app.Host + app.RootURL + constValue.WSBaskURI
)


// DispatchServerCreate Dispatch server created.
func DispatchServerCreate(server *model.Server) ([]commonDto.Message, error) {
	messages, err := rest.Do(
		http.MethodPost,
		WsServerRoot,
		*server,
		nil,
		[]int{http.StatusCreated})
	)
	return messages, err
}

// DispatchServerUpdate Dispatch server updated.
func DispatchServerUpdate(server *model.Server) {

}

// DispatchServerDelete Dispatch server deleted.
func DispatchServerDelete(URI string) {

}
