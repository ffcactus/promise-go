package base

// EventServiceInterface is the interface that event service should have.
type EventServiceInterface interface {
	DispatchCreateEvent(ResponseInterface) ([]MessageInterface, error)
	DispatchUpdateEvent(ResponseInterface) ([]MessageInterface, error)
	DispatchDeleteEvent(ResponseInterface) ([]MessageInterface, error)
}

// // EventService is the event service.
// type EventService struct {
// }
