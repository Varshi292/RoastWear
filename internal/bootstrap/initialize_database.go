package bootstrap

import (
	"github.com/Varshi292/RoastWear/internal/interfaces"
	"gorm.io/gorm"
)

func InitializeDatabase(db interfaces.Database) *gorm.DB {
	gormDB, err := db.Connect()
	if err != nil {
		panic("Database connection error: " + err.Error())
	}
	if err := db.Migrate(); err != nil {
		panic("Database migration error: " + err.Error())
	}
	return gormDB
}
