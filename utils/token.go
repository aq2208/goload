package utils

import (
	"context"
	"time"

	"github.com/aq2208/goload/configs"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(configs.GetEnv("JWT_SECRET")) // should be from env var in production
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

	// token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
	// 	// Ensure token was signed with the correct method
	// 	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
	// 		return nil, errors.New("unexpected signing method")
	// 	}
	// 	return jwtSecret, nil
	// })

	// if err != nil {
	// 	return nil, err
	// }

	// // Extract and validate claims
	// claims, ok := token.Claims.(*CustomClaims)
	// if !ok || !token.Valid {
	// 	return nil, errors.New("invalid token")
	// }

	// // Optional: check for expiry manually (usually handled by jwt lib)
	// if claims.ExpiresAt.Time.Before(time.Now()) {
	// 	return nil, errors.New("token expired")
	// }

	// return claims, nil

	panic("unimplemented")
}

func NewTokenUtil() Token {
	return &token{}
}
