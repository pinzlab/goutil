package migrator

import "gorm.io/gorm"

// Migration defines the interface that all migration types must implement
type Migration interface {
	GetCode() string
	GetName() string
	Execute(tx *gorm.DB) error
}
