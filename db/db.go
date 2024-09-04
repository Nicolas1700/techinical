package db

import (
	"fmt"
	"log"
	"techinical/config"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

var db *gorm.DB

func InitDb(config config.Config) {
	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s search_path=%s",
		config.Host, config.Port, config.User, config.Pass, config.Database, config.Schema,
	)
	dbCon, err := gorm.Open(postgres.Open(connectionString))
	if err != nil {
		log.Fatal("Error al conectar a la base de datos:", err)
	}
	db = dbCon
}

func ConectionDb() *gorm.DB {
	return db
}
