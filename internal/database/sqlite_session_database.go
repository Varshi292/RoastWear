package database

import (
	"github.com/Varshi292/RoastWear/internal/models"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SqliteSessionDatabase struct {
	dsn string
}

func NewSqliteSessionDatabase(dsn string) *SqliteSessionDatabase {
	return &SqliteSessionDatabase{dsn: dsn}
}

func (s *SqliteSessionDatabase) Connect() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(s.dsn), &gorm.Config{
		TranslateError: true,
		Logger:         logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (s *SqliteSessionDatabase) Migrate() error {
	db, err := s.Connect()
	if err != nil {
		return err
	}
	if err := db.AutoMigrate(&models.Session{}); err != nil {
		return err
	}
	return nil
}
