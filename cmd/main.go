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

	// Kafka producer
	producer, _ := producer.NewKafkaProducer()

	// File Client
	fileClient := CreateFileClient()

	accountService := service.NewAccountService(accountRepo, hash, token)
	downloadTaskService := service.NewDownloadTaskService(
		downloadTaskRepo, token, producer, db, fileClient,
	)
	accountHandler := handler.NewAccountHandler(accountService)
	downloadTaskHandler := handler.NewDownloadTaskHandler(downloadTaskService)

	// Start Kafka consumer
	go consumer.StartKafkaConsumer(downloadTaskService)

	// Handle http requests
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/v1/login", accountHandler.Login)
	mux.HandleFunc("POST /api/v1/users", accountHandler.CreateAccountHandler)
	mux.HandleFunc("POST /api/v1/download-tasks", downloadTaskHandler.CreateDownloadTaskHandler)
	mux.HandleFunc("GET /api/v1/download-tasks", downloadTaskHandler.GetListDownloadTasks)
	mux.HandleFunc("GET /api/v1/download-tasks/{id}", downloadTaskHandler.GetDownloadFile)

	// Start http server
	log.Println("Server running on :8080")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}

	defer db.Close()
}

func CreateFileClient() file.Client {
	storageMode := configs.GetEnv("STORAGE_MODE")
	if storageMode == "S3" {
		return file.NewS3Client("goload-file", "localhost:9000", "aq2208", "quancuanam2003")
	} else if storageMode == "LOCAL" {
		return file.NewLocalClient("../download")
	}

	log.Fatal("Unsupported Storage Mode!")
	return nil
}
