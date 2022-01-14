package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // using postgres sql
)

func SetupModels(pgUser string, pgPassword string, pgDB string) *gorm.DB {

	postgres_conname := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=5432 sslmode=disable", pgUser, pgPassword, pgDB)

	fmt.Println("conname is\t\t", postgres_conname)

	db, err := gorm.Open("postgres", postgres_conname)
	if err != nil {
		panic("Failed to connect to database!")
	}

	return db
}
