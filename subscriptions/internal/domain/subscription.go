package domain

import "time"

type SportSubscription struct {
	Model
	UserId  string `json:"userId" gorm:"uniqueIndex:user_to_sport_subscription_index;index;not null"`
	SportId string `json:"sportId" validate:"required,uuid" gorm:"uniqueIndex:user_to_sport_subscription_index;not null"`
}

type EventSubscription struct {
	Model
	UserId              string     `json:"userId" gorm:"uniqueIndex:user_to_event_subscription_index;index;not null"`
	EventId             string     `json:"eventId" validate:"required,uuid" gorm:"uniqueIndex:user_to_event_subscription_index;not null"`
	MothNotifiedAt      *time.Time `json:"mothNotifiedAt"`
	WeekNotifiedAt      *time.Time `json:"weekNotifiedAt"`
	ThreeDaysNotifiedAt *time.Time `json:"threeDaysNotifiedAt"`
}
