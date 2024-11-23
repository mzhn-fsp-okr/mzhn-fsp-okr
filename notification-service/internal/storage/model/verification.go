package model

import "time"

type Verification struct {
	UserId    string
	Token     string
	CreatedAt time.Time
}

type NewVerification struct {
	UserId string
	Token  string
}
