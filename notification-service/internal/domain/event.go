package domain

import "time"

type ParticipantRequirements struct {
	Gender bool   `json:"gender"`
	MinAge *int32 `json:"minAge"`
	MaxAge *int32 `json:"maxAge"`
}

type DateRange struct {
	From time.Time `json:"from"`
	To   time.Time `json:"to"`
}

type NewSubscriptionsMessage struct {
	EntityId string `json:"entityId"`
	UserId   string `json:"userId"`
	IsEvent  bool   `json:"isEvent"`
}

type UpcomingEventMessage struct {
	UserId   string `json:"userId"`
	EventId  string `json:"eventId"`
	DaysLeft uint32 `json:"daysLeft"`
}

type SportType struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type SportSubtype struct {
	Id     string    `json:"id"`
	Name   string    `json:"name"`
	Parent SportType `json:"sportType"`
}
type SportSubtype2 struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type EventInfo struct {
	Id                      string                    `json:"id"`
	EkpId                   string                    `json:"ekpId"`
	SportSubtype            SportSubtype              `json:"sportSubtype"`
	Name                    string                    `json:"name"`
	Description             string                    `json:"description"`
	Dates                   DateRange                 `json:"dates"`
	Location                string                    `json:"location"`
	Participants            int                       `json:"participants"`
	ParticipantRequirements []ParticipantRequirements `json:"participantRequirements"`
}

type EventType int

const (
	EventTypeNew EventType = iota
	EventTypeUpcoming
	EventNewSub
)

func (e EventType) String() string {
	return [...]string{"new", "upcoming"}[e]
}
