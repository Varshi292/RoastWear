package database

import (
	"errors"
	"github.com/Varshi292/RoastWear/internal/models"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SqliteSessionDatabase struct {
	Dsn string
	Db  *gorm.DB
}

func NewSqliteSessionDatabase(dsn string) *SqliteSessionDatabase {
	return &SqliteSessionDatabase{Dsn: dsn}
}

func (s *SqliteSessionDatabase) Connect() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(s.Dsn), &gorm.Config{
		TranslateError: true,
		Logger:         logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}
	s.Db = db
	return db, nil
}

func (s *SqliteSessionDatabase) Migrate() error {
	if s.Db == nil {
		return errors.New("database not initialized")
	}
	if err := s.Db.AutoMigrate(&models.Session{}); err != nil {
		return err
	}
	return nil
}
