package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(DSN string) *gorm.DB {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: DSN,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
