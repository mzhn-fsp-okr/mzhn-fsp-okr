package domain

import (
	"mzhn/subscriptions-service/pb/espb"
	"mzhn/subscriptions-service/pb/sspb"
)

type SubscriptionsStorage interface {
	CreateSport(dto *SportSubscription) (*SportSubscription, error)
	CreateEvent(dto *EventSubscription) (*EventSubscription, error)
	DeleteSport(sportSubscription *SportSubscription) error
	DeleteEvent(eventSubscription *EventSubscription) error
	GetUserEventsId(userId string) ([]string, error)
	GetUserSportsId(userId string) ([]string, error)
	GetUsersSubscribedToEvent(eventId string) ([]string, error)
	GetUsersSubscribedToSport(sportId string) ([]string, error)
	GetUsersFromEventByDaysLeft(eventId string, daysLeft sspb.DaysLeft) ([]string, error)
	NotifyUser(userId string, daysLeft sspb.DaysLeft) error
}

type SubscriptionsService interface {
	SubscribeToSport(dto *SportSubscription) (*SportSubscription, error)
	SubscribeToEvent(dto *EventSubscription) (*EventSubscription, error)
	UnsubscribeFromSport(dto *SportSubscription) error
	UnsubscribeFromEvent(dto *EventSubscription) error
	GetUserEvents(userId string) ([]*espb.EventInfo, error)
	GetUserSports(userId string) ([]*espb.SportTypeWithSubtypes, error)
	GetUsersSubscribedToEvent(eventId string) ([]string, error)
	GetUsersSubscribedToSport(sportId string) ([]string, error)
	GetUsersFromEventByDaysLeft(eventId string, daysLeft sspb.DaysLeft) ([]string, error)
	NotifyUser(userId string, daysLeft sspb.DaysLeft) error
}
