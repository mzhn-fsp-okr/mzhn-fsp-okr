package domain

import "errors"

var (
	ErrUnathorized = errors.New("unathorized")

	ErrVerificationNotFound     = errors.New("verification not found")
	ErrVerificationExpired      = errors.New("verification expired")
	ErrVerificationInvalidToken = errors.New("verification invalid token")
)
