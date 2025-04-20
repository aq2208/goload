package repository

import (
	"context"
	"database/sql"
	"github.com/aq2208/goload/internal/model"
)

type AccountRepository interface {
	CreateAccount(ctx context.Context, account model.User) error
	GetAccountById(ctx context.Context, id uint64) (*model.User, error)
	GetAccountByUsername(ctx context.Context) (*model.User, error)
}

type accountRepository struct {
	db *sql.DB
}

// CreateAccount implements AccountRepository.
func (a *accountRepository) CreateAccount(ctx context.Context, account model.User) error {
	panic("unimplemented")
}

// GetAccountById implements AccountRepository.
func (a *accountRepository) GetAccountById(ctx context.Context, id uint64) (*model.User, error) {
	panic("unimplemented")
}

// GetAccountByUsername implements AccountRepository.
func (a *accountRepository) GetAccountByUsername(ctx context.Context) (*model.User, error) {
	panic("unimplemented")
}

func NewMySqlAccountRepository(db *sql.DB) AccountRepository {
	return &accountRepository{db: db}
}
