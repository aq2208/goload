package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/aq2208/goload/internal/service"
)

type AccountHandler struct {
	service service.AccountService
}

func NewAccountHandler(service service.AccountService) *AccountHandler {
	return &AccountHandler{service: service}
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func CreateAcountHandler(w http.ResponseWriter, r *http.Request) {
}

func (h *AccountHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	log.Default().Printf("LoginRequest: %v", req)

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