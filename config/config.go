package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"log"

	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "mysql"
	password = "mysql"
	dbName   = "test"
)

func DatabaseConnection() *gorm.DB {
	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, err := gorm.Open(mysql.Open(sqlInfo), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}
