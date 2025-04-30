package service

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/aq2208/goload/internal/model"
	"github.com/aq2208/goload/internal/repository"
	"github.com/aq2208/goload/utils"
)

type CreateAccountRequest struct {
	Username string
	Password string
}

type CreateAccountResponse struct{}

type AccountService interface {
	CreateAccount(ctx context.Context, username string, password string, email string) (uint64, error)
	CreateSession(ctx context.Context, username string, password string) (string, error)
}

type accountService struct {
	repo repository.AccountRepository
	hash utils.Hash
	token utils.Token
}

// CreateAccount implements AccountService.
func (a *accountService) CreateAccount(ctx context.Context, username string, password string, email string) (uint64, error) {
	// check username existed by username
	existing, err := a.repo.GetAccountByUsername(ctx, username)
	if err != nil && err != sql.ErrNoRows {
		log.Default().Printf("CreateAccount error: %v", err)
		return 0, err
	}
	if existing != nil {
		return 0, errors.New("Account already existed")
	}

	// create new account and hash pwd
	hashed, err := a.hash.Hash(ctx, password)
	if err != nil {
		return 0, err
	}

	return a.repo.CreateAccount(
		ctx, 
		model.User{
			Email: email,
			Username: username,
			PasswordHash: hashed,
		})
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
