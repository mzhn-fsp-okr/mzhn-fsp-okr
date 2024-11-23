package domain

import "mzhn/subscriptions-service/pb/espb"

type SubscriptionsStorage interface {
	CreateSport(dto *SportSubscription) (*SportSubscription, error)
	CreateEvent(dto *EventSubscription) (*EventSubscription, error)
	DeleteSport(sportSubscription *SportSubscription) error
	DeleteEvent(eventSubscription *EventSubscription) error
	GetUserEventsId(userId string) ([]string, error)
	GetUsersSubscribedToEvent(eventId string) ([]string, error)
	GetUsersSubscribedToSport(sportId string) ([]string, error)
}

type SubscriptionsService interface {
	SubscribeToSport(dto *SportSubscription) (*SportSubscription, error)
	SubscribeToEvent(dto *EventSubscription) (*EventSubscription, error)
	UnsubscribeFromSport(dto *SportSubscription) error
	UnsubscribeFromEvent(dto *EventSubscription) error
	GetUserEvents(userId string) ([]*espb.EventInfo, error)
	GetUsersSubscribedToEvent(eventId string) ([]string, error)
	GetUsersSubscribedToSport(sportId string) ([]string, error)
}
