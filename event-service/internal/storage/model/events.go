package model

import "time"

type Pagination struct {
	Limit  uint64
	Offset uint64
}

type EventsFilters struct {
	Pagination
	SportTypesId    []string
	SportSubtypesId []string
	MinAge          *int
	MaxAge          *int
	Sex             *bool
	MinParticipants *int
	MaxParticipants *int
	Location        *string
	StartDate       *time.Time
	EndDate         *time.Time
	Name            *string
}
