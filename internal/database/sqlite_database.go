// Package database ...
package database

import (
	"RoastWear/internal/models"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

type SqliteDatabase struct {
	dsn string
}

func NewSqliteDatabase(dsn string) *SqliteDatabase {
	return &SqliteDatabase{dsn: dsn}
}

func (s *SqliteDatabase) Connect() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(s.dsn), &gorm.Config{
		TranslateError: true,
		Logger:         logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}
	log.Println("✅ Connected to database successfully")
	return db, nil
}

// Migrate ...
//
// Parameters:
//   - db: ...
//
// Returns:
//   - error: ...
func (s *SqliteDatabase) Migrate() error {
	db, err := s.Connect()
	if err != nil {
		return err
	}
	if err := db.AutoMigrate(&models.User{}); err != nil {
		return err
	}
	if err := db.Table("archived_users").AutoMigrate(&models.User{}); err != nil {
		return err
	}
	log.Println("✅ Migrated database successfully")
	return nil
}
