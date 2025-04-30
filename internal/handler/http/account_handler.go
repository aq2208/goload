package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"regexp"

	"github.com/aq2208/goload/internal/service"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

func IsValidEmailRegex(email string) bool {
	return emailRegex.MatchString(email)
}

type AccountHandler struct {
	service service.AccountService
}

func NewAccountHandler(service service.AccountService) *AccountHandler {
	return &AccountHandler{service: service}
}

type CreateAccountRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
}

type CreateAccountResponse struct {
	Id uint64 `json:"id"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func (h *AccountHandler) CreateAccountHandler(w http.ResponseWriter, r *http.Request) {
	var req CreateAccountRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	log.Default().Printf("CreateAccountRequest: %v", req)

	// validate
	if (req.Username == "" || req.Password == "" || req.Email == "") {
		http.Error(w, "Missing required field(s)", http.StatusBadRequest)
		return
	}

	if !IsValidEmailRegex(req.Email) {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// create account
	accountId, err := h.service.CreateAccount(context.Background(), req.Username, req.Password, req.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := &CreateAccountResponse{Id: accountId}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *AccountHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest

	// parse to request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	log.Default().Printf("LoginRequest: %v", req)

	// validate
	if (req.Username == "" || req.Password == "") {
		http.Error(w, "Invalid username or password", http.StatusBadRequest)
		return
	}

	token, err := h.service.CreateSession(context.Background(), req.Username, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	resp := &LoginResponse{Token: token}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}