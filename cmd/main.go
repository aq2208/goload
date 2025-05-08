package main

import (
	"log"
	"net/http"

	"github.com/aq2208/goload/configs"
	"github.com/aq2208/goload/internal/dataaccess/database/dbconnect"
	"github.com/aq2208/goload/internal/dataaccess/file"
	"github.com/aq2208/goload/internal/dataaccess/mq/consumer"
	"github.com/aq2208/goload/internal/dataaccess/mq/producer"
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
	downloadTaskRepo := repository.NewDownloadTaskRepository(db)
	hash := utils.NewHashUtil()
	token := utils.NewTokenUtil()

	// Start Kafka producer and consumer
	producer, _ := producer.NewKafkaProducer()
	go consumer.StartKafkaConsumer()

	// File Client
	LocalFileClient := file.NewLocalFileClient("./download")

	accountService := service.NewAccountService(accountRepo, hash, token)
	downloadTaskService := service.NewDownloadTaskService(downloadTaskRepo, token, producer, db, LocalFileClient)
	accountHandler := handler.NewAccountHandler(accountService)
	downloadTaskHandler := handler.NewDownloadTaskHandler(downloadTaskService)

	// Handle http requests
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/v1/login", accountHandler.Login)
	mux.HandleFunc("POST /api/v1/users", accountHandler.CreateAccountHandler)
	mux.HandleFunc("POST /api/v1/download-tasks", downloadTaskHandler.CreateDownloadTaskHandler)

	// Start http server
	log.Println("Server running on :8080")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}

	defer db.Close()
}