// internal/database/sqlite_purchase.go
package database

import (
	"github.com/Varshi292/RoastWear/internal/models"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SqlitePurchaseDatabase struct {
	dsn string
	db  *gorm.DB
}

func NewSqlitePurchaseDatabase(dsn string) *SqlitePurchaseDatabase {
	return &SqlitePurchaseDatabase{dsn: dsn}
}

func (p *SqlitePurchaseDatabase) Connect() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(p.dsn), &gorm.Config{
		TranslateError: true,
		Logger:         logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}
	p.db = db
	return db, nil
}

func (p *SqlitePurchaseDatabase) Migrate() error {
	if err := p.db.AutoMigrate(&models.Purchase{}); err != nil {
		return err
	}
	return nil
}
