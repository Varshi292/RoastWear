package interfaces

import (
	"gorm.io/gorm"
)

type Database interface {
	Connect() (*gorm.DB, error)
	Migrate() error
}
