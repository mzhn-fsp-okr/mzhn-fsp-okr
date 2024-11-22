package model

import "time"

type EventsFilters struct {
	Limit  *uint64
	Offset *uint64

	SportTypeId *int

	StartDate *time.Time
	EndDate   *time.Time
}
