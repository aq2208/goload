package utils

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("aq2208") // should be from env var in production
var tokenTtl time.Duration = time.Minute*5

type TokenClaims struct {
	UserId uint64
	jwt.RegisteredClaims
}

type Token interface {
	GenerateToken(ctx context.Context, accountId uint64) (string, error)
	GetAccountIdAndExpireTime(ctx context.Context, token string) (uint64, error)
	RefreshToken()
}

type token struct {}

// RefreshToken implements Token.
func (t *token) RefreshToken() {
	panic("unimplemented")
}

// GenerateToken implements Token.
func (t *token) GenerateToken(ctx context.Context, accountId uint64) (string, error) {
	tokenClaims := &TokenClaims{
		UserId: accountId, 
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTtl)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)

	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// GetAccountIdAndExpireTime implements Token.
func (t *token) GetAccountIdAndExpireTime(ctx context.Context, token string) (uint64, error) {
	// if token is expired -> ....
	panic("unimplemented")
}

func NewTokenUtil() Token {
	return &token{}
}
