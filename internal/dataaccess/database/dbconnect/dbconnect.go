package dbconnect

import (
	"database/sql"
	"log"

	"github.com/aq2208/goload/configs"
	"github.com/go-sql-driver/mysql"
)

func NewMySqlConnection() (*sql.DB, error) {
    cfg := mysql.NewConfig()
	cfg.User = configs.GetEnv("DB_USER")
	cfg.Passwd = configs.GetEnv("DB_PWD")
	cfg.Net = "tcp"
	cfg.Addr = configs.GetEnv("DB_HOST")
	cfg.DBName = configs.GetEnv("DB_NAME")
	cfg.ParseTime = true

    log.Default().Printf("Test get config: %s", configs.GetEnv("DB_USER"))
    log.Default().Printf("Database DSN: %s", cfg.FormatDSN())
    
    db, err := sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatalf("Connect MySQL Database error: %v", err)
        return nil, err
    }

    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatalf("Ping MySQL Database error: %v", pingErr)
        return nil, pingErr
    }

    log.Printf("Database connected successful!")

    return db, nil
}