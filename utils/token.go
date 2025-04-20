package utils

import (
	"context"
	"time"

	// "github.com/golang-jwt/jwt"
)

type Token interface {
	GenerateToken(ctx context.Context, accountId uint64) (string, time.Time, error)
	GetAccountIdAndExpireTime(ctx context.Context, token string) (uint64, time.Time, error)
}

type token struct {
	expiresIn time.Duration
}

// GenerateToken implements Token.
func (t *token) GenerateToken(ctx context.Context, accountId uint64) (string, time.Time, error) {
	panic("unimplemented")
}

// GetAccountIdAndExpireTime implements Token.
func (t *token) GetAccountIdAndExpireTime(ctx context.Context, token string) (uint64, time.Time, error) {
	panic("unimplemented")
}

func NewTokenUtil() Token {
	return &token{}
}
