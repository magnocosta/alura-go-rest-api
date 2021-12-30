package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func OpenDBConnection() {
	connStr := "host=localhost user=alura password=alura!12 dbname=alura_db port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
}
