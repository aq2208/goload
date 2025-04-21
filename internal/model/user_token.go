package model

import "time"

type UserToken struct {
	ID           uint64
	UserID       uint64
	RefreshToken string
	UserAgent    *string
	IPAddress    *string
	ExpiresAt    time.Time
	CreatedAt    time.Time
}
