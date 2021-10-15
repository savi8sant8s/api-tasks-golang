package database

import (
	"fmt"
	"os"
	"savi8sant8s/api/entity"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct { }

var db = new(Database).Instance()

func (db *Database) Config() string {
	godotenv.Load()
	HOST := "host=localhost"
	USER := "user=" + os.Getenv("PG_USER")
	PASSWORD := "password=" + os.Getenv("PG_PASSWORD")
	DATABASE := "dbname=" + os.Getenv("PG_DATABASE")
	PORT := "port=5432"
	SSLMODE := "sslmode=disable"
	TIMEZONE := "TimeZone=America/Recife"
  
	return fmt.Sprintf("%s %s %s %s %s %s %s", HOST, USER, PASSWORD, DATABASE, PORT, SSLMODE, TIMEZONE)
  }

func (database *Database) Instance() *gorm.DB {
	databaseInstance, _ := gorm.Open(postgres.Open(database.Config()), &gorm.Config{})
	databaseInstance.AutoMigrate(&entity.User{}, &entity.Task{})
	return databaseInstance
}
