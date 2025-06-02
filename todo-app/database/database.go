package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/tahsin005/layered-based-architecture/todo-app/config"
	_ "github.com/lib/pq"
)


func InitDB(cfg config.Config) *sql.DB {
    dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

    db, err := sql.Open("postgres", dsn)
    if err != nil {
        log.Fatal("Failed to open DB:", err)
    }

    if err := db.Ping(); err != nil {
        log.Fatal("Failed to ping DB:", err)
    }

    return db
}