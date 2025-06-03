package pg

import (
	"github.com/pinzlab/goutil/terminal"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Open establishes a connection to the PostgreSQL database using the provided DSN and optional configuration options.
// It panics if the connection cannot be established.
// Returns a pointer to the initialized DB instance.
func Open(dns string, opts ...gorm.Option) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dns), opts...)
	if err != nil {
		terminal.Panic(err)
	}
	return db
}
