package infrastructure

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSQLiteDB() *gorm.DB {
	// dbPath := os.Getenv("DB_FILE")
	// db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	db, err := gorm.Open(sqlite.Open("sqlite/quote.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}
	return db
}
