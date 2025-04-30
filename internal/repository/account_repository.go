package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/aq2208/goload/internal/model"
)

type AccountRepository interface {
	CreateAccount(ctx context.Context, account model.User) (uint64, error)
	GetAccountById(ctx context.Context, id uint64) (*model.User, error)
	GetAccountByUsername(ctx context.Context, username string) (*model.User, error)
}

type accountRepository struct {
	db *sql.DB
}

// CreateAccount implements AccountRepository.
func (a *accountRepository) CreateAccount(ctx context.Context, account model.User) (uint64, error) {
	result, err := a.db.Exec("INSERT INTO user (email, username, password_hash) VALUES(?,?,?)", account.Email, account.Username, account.PasswordHash)
	if err != nil {
		return 0, err
	}

	insertId, _ := result.LastInsertId()
	log.Default().Printf("Inserted user %d successful", insertId)

	return uint64(insertId), nil
}

// GetAccountById implements AccountRepository.
func (a *accountRepository) GetAccountById(ctx context.Context, id uint64) (*model.User, error) {
	var user model.User

	row := a.db.QueryRow("SELECT id, username, password_hash FROM user WHERE id = ? and auth_provider = 'local'", id)

	err := row.Scan(&user.ID, &user.Username, &user.PasswordHash)
	if (err != nil) {
		if (err == sql.ErrNoRows) {
			return nil, fmt.Errorf("accountId %d: no such user", id)
		}

		return nil, err
	}

	return &user, nil
}

// GetAccountByUsername implements AccountRepository.
func (a *accountRepository) GetAccountByUsername(ctx context.Context, username string) (*model.User, error) {
	var user model.User

	row := a.db.QueryRow("SELECT id, username, password_hash FROM user WHERE username = ? and auth_provider = 'local'", username)

	err := row.Scan(&user.ID, &user.Username, &user.PasswordHash)
	if err != nil {
		if err == sql.ErrNoRows {
			// return nil, fmt.Errorf("username %s: no such user", username)
			return nil, err
		}

		return nil, err
	}

	return &user, nil
}

func NewAccountRepository(db *sql.DB) AccountRepository {
	return &accountRepository{db: db}
}
