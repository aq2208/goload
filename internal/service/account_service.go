package service

import (
	"context"
	"errors"

	"github.com/aq2208/goload/internal/repository"
	"github.com/aq2208/goload/utils"
)

type CreateAccountRequest struct {
	Username string
	Password string
}

type CreateAccountResponse struct{}

type AccountService interface {
	CreateAccount(ctx context.Context, req CreateAccountRequest) (CreateAccountResponse, error)
	CreateSession(ctx context.Context, username string, password string) (string, error)
}

type accountService struct {
	repo repository.AccountRepository
	hash utils.Hash
	token utils.Token
}

// CreateAccount implements AccountService.
func (a *accountService) CreateAccount(ctx context.Context, req CreateAccountRequest) (CreateAccountResponse, error) {
	// check username existed, if not, create new user and create hash pwd

	panic("unimplemented")
}

// CreateSession implements AccountService.
func (a *accountService) CreateSession(ctx context.Context, username string, password string) (string, error) {
	// check user exist by username
	user, err := a.repo.GetAccountByUsername(ctx, username)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New("User not found")
	}

	// check password hash matched
	if !a.hash.IsHashEqual(ctx, password, user.PasswordHash)  {
		return "", errors.New("Invalid Credentials")
	}

	// if matched, generate and response jwt
	token, err := a.token.GenerateToken(ctx, user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func NewAccountService(repo repository.AccountRepository, hash utils.Hash, token utils.Token) AccountService {
	return &accountService{repo: repo, hash: hash, token: token}
}
