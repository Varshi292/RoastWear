package database

import (
	"github.com/Varshi292/RoastWear/internal/models"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SqliteUserDatabase struct {
	dsn string
	db  *gorm.DB
}

func NewSqliteUserDatabase(dsn string) *SqliteUserDatabase {
	return &SqliteUserDatabase{dsn: dsn}
}

func (s *SqliteUserDatabase) Connect() (*gorm.DB, error) {
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

func (s *SqliteUserDatabase) Migrate() error {
	if err := s.db.AutoMigrate(&models.User{}); err != nil {
		return err
	}
	if err := s.db.Table("archived_users").AutoMigrate(&models.User{}); err != nil {
		return err
	}
	return nil
}
