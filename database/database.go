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
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file.")
	}
	HOST := "host=" + os.Getenv("PG_HOST")
	USER := "user=" + os.Getenv("PG_USER")
	PASSWORD := "password=" + os.Getenv("PG_PASSWORD")
	DBNAME := "dbname=" + os.Getenv("PG_DBNAME")
	PORT := "port=" + os.Getenv("PG_PORT")
	SSLMODE := "sslmode=" + os.Getenv("PG_SSLMODE")
	TIMEZONE := "TimeZone=" + os.Getenv("PG_TIMEZONE")
  
	return fmt.Sprintf("%s %s %s %s %s %s %s", HOST, USER, PASSWORD, DBNAME, PORT, SSLMODE, TIMEZONE)
  }

func (database *Database) Instance() *gorm.DB {
	databaseInstance, err := gorm.Open(postgres.Open(database.Config()), &gorm.Config{})
	if err != nil {
	  panic("Failed to connect database.")
	}
	databaseInstance.AutoMigrate(&entity.User{}, &entity.Session{}, &entity.Task{})
	return databaseInstance
}
