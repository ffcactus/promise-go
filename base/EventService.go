package base

// EventServiceInterface is the interface that event service should have.
type EventServiceInterface interface {
	DispatchCreateEvent(ResponseInterface) ([]Message, error)
	DispatchUpdateEvent(ResponseInterface) ([]Message, error)
	DispatchDeleteEvent(ResponseInterface) ([]Message, error)
	DispatchDeleteCollectionEvent(category string) ([]Message, error)
}

// // EventService is the event service.
// type EventService struct {
// }
