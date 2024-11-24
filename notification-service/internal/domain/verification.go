package domain

import "time"

type VerificationRequest struct {
	Token string
}

type VerificationCodeRequest struct {
	UserId string
}

type Verification struct {
	UserId   string
	Token    string
	ExpireAt time.Time
}
