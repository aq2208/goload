package model

import "time"

type AuthProvider string

const (
	AuthProviderLocal  AuthProvider = "local"
	AuthProviderGoogle AuthProvider = "google"
)

type User struct {
	ID           int64        `db:"id"`
	Email        *string      `db:"email"`        // nullable
	Username     *string      `db:"username"`     // nullable
	PasswordHash *string      `db:"password_hash"`// nullable
	AuthProvider AuthProvider `db:"auth_provider"`
	GoogleID     *string      `db:"google_id"`    // nullable
	CreatedAt    time.Time    `db:"created_at"`
	UpdatedAt    time.Time    `db:"updated_at"`
}