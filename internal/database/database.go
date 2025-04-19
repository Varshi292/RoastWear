// Package database ...
package database

import (
	"github.com/Varshi292/RoastWear/internal/models"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

// Open ...
//
// Parameters:
//   - dsn: ...
//
// Returns:
//   - *gorm.DB: ...
//   - error: ...
func Open(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		TranslateError: true,
		Logger:         logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}
	log.Println("✅ Opened database successfully")
	return db, nil
}

// Migrate ...
//
// Parameters:
//   - db: ...
//
// Returns:
//   - error: ...
//
// Migrate ...
//
// Parameters:
//   - db: ...
//
// Migrate ...
//
// Parameters:
//   - db: *gorm.DB - the database connection
//
// Returns:
//   - error: error if any migration fails
func Migrate(db *gorm.DB) error {
	// Migrate the 'users' table
	if err := db.AutoMigrate(&models.User{}); err != nil {
		return err
	}

	// Migrate the 'deleted_users' table (manually set name)
	if err := db.Table("deleted_users").AutoMigrate(&models.User{}); err != nil {
		return err
	}

	// Migrate the 'images' table
	if err := db.AutoMigrate(&models.Image{}); err != nil {
		return err
	}

	// Migrate the 'user_uploads' table
	if err := db.AutoMigrate(&models.UserUpload{}); err != nil {
		return err
	}

	// ✅ Migrate the 'sessions' table
	if err := db.AutoMigrate(&models.Session{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&models.Purchase{}); err != nil {
		return err
	}

	log.Println("✅ Migrated database successfully")
	return nil
}
