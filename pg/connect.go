package pg

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB is a type alias for gorm.DB.
// It represents a connection to the database and provides methods for querying and managing records.
type DB = gorm.DB

// Option is a type alias for gorm.Option.
// It is used to configure the GORM database connection.
type Option = gorm.Option

// Open establishes a connection to the PostgreSQL database using the provided DSN and optional configuration options.
// It panics if the connection cannot be established.
// Returns a pointer to the initialized DB instance.
func Open(dns string, opts ...Option) *DB {
	db, err := gorm.Open(postgres.Open(dns), opts...)
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}
	return db
}
