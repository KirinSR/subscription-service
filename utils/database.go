package utils

import (
	"fmt"
	"log"

	"example.com/subscription-service/config"
	"github.com/jinzhu/gorm"
)

// InitDB initializes the database
func InitDB(cfg *config.Config) *gorm.DB {
	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBName, cfg.DBPassword, cfg.DBSSLMode)
	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
