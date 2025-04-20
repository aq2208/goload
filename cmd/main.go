package main

import (
	"log"

	"github.com/aq2208/goload/internal/dataaccess/database/dbconnect"
)

func main() {
	// Connect database
	db, err := dbconnect.NewMySqlConnection()
	if err != nil {
		log.Fatalf("Error connecting database: %v", err)
		return
	}

	defer db.Close()
}