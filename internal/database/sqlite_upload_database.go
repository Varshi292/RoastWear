package database

import (
	"github.com/Varshi292/RoastWear/internal/models"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SqliteUploadDatabase struct {
	dsn string
}

func NewSqliteUploadDatabase(dsn string) *SqliteUploadDatabase {
	return &SqliteUploadDatabase{dsn: dsn}
}

func (s *SqliteUploadDatabase) Connect() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(s.dsn), &gorm.Config{
		TranslateError: true,
		Logger:         logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (s *SqliteUploadDatabase) Migrate() error {
	db, err := s.Connect()
	if err != nil {
		return err
	}
	if err := db.AutoMigrate(&models.UserUpload{}); err != nil {
		return err
	}
	return nil
}
