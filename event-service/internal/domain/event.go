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

type EventLoadInfo struct {
	EkpId        string
	SportType    string
	SportSubtype string
	Name         string
	Description  string
	Dates        DateRange
	Location     string
	Participants int

	ParticipantRequirements []ParticipantRequirements
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
