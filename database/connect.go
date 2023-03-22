package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: "host=localhost user=postgres password=0511 dbname=gin_gonic port=5432 sslmode=disable TimeZone=Asia/Shanghai",
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}
