package domain

type SportSubscription struct {
	Model
	UserId  string `json:"userId" gorm:"uniqueIndex:user_to_sport_subscription_index;index;not null"`
	SportId string `json:"sportId" validate:"required,uuid" gorm:"uniqueIndex:user_to_sport_subscription_index;not null"`
}

type EventSubscription struct {
	Model
	UserId  string `json:"userId" gorm:"uniqueIndex:user_to_sport_subscription_index;index;not null"`
	EventId string `json:"eventId" validate:"required,uuid" gorm:"uniqueIndex:user_to_sport_subscription_index;not null"`
}
