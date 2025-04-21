package model

import "time"

type AuthProvider string

const (
	AuthProviderLocal  AuthProvider = "local"
	AuthProviderGoogle AuthProvider = "google"
)

type User struct {
	ID           uint64
	Email        *string
	Username     *string
	PasswordHash *string
	AuthProvider AuthProvider
	GoogleID     *string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
