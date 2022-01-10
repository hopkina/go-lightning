package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // using postgres sql
)

func SetupModels() *gorm.DB {

	postgres_conname := "host=localhost user=park password=life dbname=hawthorn port=5432 sslmode=disable"

	fmt.Println("conname is\t\t", postgres_conname)

	db, err := gorm.Open("postgres", postgres_conname)
	if err != nil {
		panic("Failed to connect to database!")
	}

	return db
}
