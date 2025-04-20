package service

import (
	"context"

	"github.com/aq2208/goload/internal/repository"
)

type CreateAccountRequest struct {
	Username string
	Password string
}

type CreateAccountResponse struct {
}

type CreateSessionRequest struct {
}

type CreateSessionResponse struct {
}

type AccountService interface {
	CreateAccount(ctx context.Context, req CreateAccountRequest) (CreateAccountResponse, error)
	CreateSession(ctx context.Context, req CreateSessionRequest) (CreateSessionResponse, error)
}

type accountService struct {
	repo repository.AccountRepository
}

// CreateAccount implements AccountService.
func (a *accountService) CreateAccount(ctx context.Context, req CreateAccountRequest) (CreateAccountResponse, error) {
	// check username existed, if not, create new user and create hash pwd

	panic("unimplemented")
}

// CreateSession implements AccountService.
func (a *accountService) CreateSession(ctx context.Context, req CreateSessionRequest) (CreateSessionResponse, error) {
	// check username and password hashed matched

	// if matched, generate and response jwt
	
	panic("unimplemented")
}

func NewAccountService() AccountService {
	return &accountService{}
}

