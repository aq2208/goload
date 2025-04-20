package model

import "time"

type UserToken struct {
	ID           int64     `db:"id"`
	UserID       int64     `db:"user_id"`
	RefreshToken string    `db:"refresh_token"`
	UserAgent    *string   `db:"user_agent"`  // nullable
	IPAddress    *string   `db:"ip_address"`  // nullable
	ExpiresAt    time.Time `db:"expires_at"`
	CreatedAt    time.Time `db:"created_at"`
}