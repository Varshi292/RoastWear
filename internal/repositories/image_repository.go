package repositories

import (
	"gorm.io/gorm"
)

// ImageRepository ...
//
// Fields:
//   - Db: ...
type ImageRepository struct {
	Db *gorm.DB
}
