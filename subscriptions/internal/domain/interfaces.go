package domain

type SubscriptionsStorage interface {
	CreateSport(dto *SportSubscription) (*SportSubscription, error)
	CreateEvent(dto *EventSubscription) (*EventSubscription, error)
}

type SubscriptionsService interface {
	CreateSubscriptionToSport(dto *SportSubscription) (*SportSubscription, error)
	CreateSubscriptionToEvent(dto *EventSubscription) (*EventSubscription, error)
}
