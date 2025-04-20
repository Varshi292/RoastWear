package database

import (
	"github.com/Varshi292/RoastWear/internal/models"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SqliteSessionDatabase struct {
	dsn string
	db  *gorm.DB
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
	s.db = db
	return db, nil
}

func (s *SqliteSessionDatabase) Migrate() error {
	if err := s.db.AutoMigrate(&models.Session{}); err != nil {
		return err
	}
	return nil
}
