package service

import "context"

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