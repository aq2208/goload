package utils

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/aq2208/goload/configs"
	"github.com/golang-jwt/jwt/v5"
)

type TokenClaims struct {
	UserId uint64 `json:"user_id"`
	jwt.RegisteredClaims
}

type Token interface {
	GenerateToken(ctx context.Context, accountId uint64) (string, error)
	GetAccountIdAndExpireTime(ctx context.Context, token string) (uint64, error)
	RefreshToken()
}

type token struct {
	JwtSecret []byte
	TokenTtl  time.Duration
}

// RefreshToken implements Token.
func (t *token) RefreshToken() {
	panic("unimplemented")
}

// GenerateToken implements Token.
func (t *token) GenerateToken(ctx context.Context, accountId uint64) (string, error) {
	log.Printf("JWT secret length: %d bytes (%d bits)\n", len(t.JwtSecret), len(t.JwtSecret)*8)

	tokenClaims := &TokenClaims{
		UserId: accountId,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(t.TokenTtl)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)

	signedToken, err := token.SignedString(t.JwtSecret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// GetAccountIdAndExpireTime implements Token.
func (t *token) GetAccountIdAndExpireTime(ctx context.Context, tokenString string) (uint64, error) {
	// if token is expired -> ....

	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		// implement callback function func(token *jwt.Token)
		
		// ensure token was signed with the correct method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		return t.JwtSecret, nil
	})

	if err != nil {
		return 0, err
	}

	// Extract and validate claims
	claims, ok := token.Claims.(*TokenClaims)
	if !ok || !token.Valid {
		return 0, errors.New("invalid token")
	}

	// Optional: check for expiry manually (usually handled by jwt lib)
	if claims.ExpiresAt.Time.Before(time.Now()) {
		return 0, errors.New("token expired")
	}

	return claims.UserId, nil
}

func NewTokenUtil() Token {
	return &token{
		JwtSecret: []byte(configs.GetEnv("JWT_SECRET")),
		TokenTtl: time.Minute * 5,
	}
}
