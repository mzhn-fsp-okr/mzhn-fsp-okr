package domain

type SubscriptionsStorage interface {
	CreateSport(dto *SportSubscription) (*SportSubscription, error)
	CreateEvent(dto *EventSubscription) (*EventSubscription, error)
	DeleteSport(sportSubscription *SportSubscription) error
	DeleteEvent(eventSubscription *EventSubscription) error
}

type SubscriptionsService interface {
	SubscribeToSport(dto *SportSubscription) (*SportSubscription, error)
	SubscribeToEvent(dto *EventSubscription) (*EventSubscription, error)
	UnsubscribeFromSport(dto *SportSubscription) error
	UnsubscribeFromEvent(dto *EventSubscription) error
}
