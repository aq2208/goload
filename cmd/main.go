package main

import (
	"log"
	"net/http"

	"github.com/aq2208/goload/configs"
	"github.com/aq2208/goload/internal/dataaccess/database/dbconnect"
	handler "github.com/aq2208/goload/internal/handler/http"
	"github.com/aq2208/goload/internal/repository"
	"github.com/aq2208/goload/internal/service"
	"github.com/aq2208/goload/utils"
)

func main() {
	// load env
	configs.LoadEnv()

	// Connect database
	db, err := dbconnect.NewMySqlConnection()
	if err != nil {
		log.Fatalf("Error connecting database: %v", err)
		return
	}

	// Dependency Injection
	accountRepo := repository.NewAccountRepository(db)
	hash := utils.NewHashUtil()
	token := utils.NewTokenUtil()
	accountService := service.NewAccountService(accountRepo, hash, token)
	accountHandler := handler.NewAccountHandler(accountService)

	// Handle http requests
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/v1/login", accountHandler.Login)
	mux.HandleFunc("POST /api/v1/users", accountHandler.CreateAccountHandler)

	log.Println("Server running on :8080")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}

	defer db.Close()
}