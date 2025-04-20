package database

import (
	"github.com/Varshi292/RoastWear/internal/models"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SqliteUploadDatabase struct {
	dsn string
	db  *gorm.DB
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
	s.db = db
	return db, nil
}

func (s *SqliteUploadDatabase) Migrate() error {
	if err := s.db.AutoMigrate(&models.UserUpload{}); err != nil {
		return err
	}
	return nil
}
