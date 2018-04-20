package base

const (
	// CreateEvent is a event type.
	CreateEvent = "Create"
	// UpdateEvent is a event type.
	UpdateEvent = "Update"
	// DeleteEvent is a event type.
	DeleteEvent = "Delete"
	// DeleteCollectionEvent is a event type for resource collection deletion.
	DeleteCollectionEvent = "DeleteCollection"
)

// EventServiceInterface is the interface that event service should have.
type EventServiceInterface interface {
	DispatchCreateEvent(ResponseInterface) ([]Message, error)
	DispatchUpdateEvent(ResponseInterface) ([]Message, error)
	DispatchDeleteEvent(ResponseInterface) ([]Message, error)
	DispatchDeleteCollectionEvent(category string) ([]Message, error)
}

