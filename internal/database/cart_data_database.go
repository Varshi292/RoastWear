// internal/database/sqlite_cart.go
package database

import (
	"github.com/Varshi292/RoastWear/internal/models"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SqliteCartDatabase struct {
	dsn string
	db  *gorm.DB
}

func NewSqliteCartDatabase(dsn string) *SqliteCartDatabase {
	return &SqliteCartDatabase{dsn: dsn}
}

func (c *SqliteCartDatabase) Connect() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(c.dsn), &gorm.Config{
		TranslateError: true,
		Logger:         logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}
	c.db = db
	return db, nil
}

func (c *SqliteCartDatabase) Migrate() error {
	if err := c.db.AutoMigrate(&models.CartItem{}); err != nil {
		return err
	}
	return nil
}
