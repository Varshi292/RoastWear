package bootstrap

import (
	"github.com/Varshi292/RoastWear/internal/interfaces"
	"gorm.io/gorm"
	"log"
)

func initializeDatabase(db interfaces.Database) *gorm.DB {
	gormDB, err := db.Connect()
	if err != nil {
		log.Fatalf("Database connection error: %s", err)
	}
	if err := db.Migrate(); err != nil {
		log.Fatalf("Database migration error: %s", err)
	}
	return gormDB
}
