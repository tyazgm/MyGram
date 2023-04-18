package database

import (
	"MyGram/model"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func ConnectDB() {
	var (
		host     = "localhost"
		port     = "5432"
		user     = "postgres"
		password = "postgres"
		dbname   = "mygram"
	)

	config := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("Connecting to Database Fail...", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	err = sqlDB.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println(db)

	db.Debug().AutoMigrate(model.User{}, model.Photo{}, model.SocialMedia{}, model.Comment{})
}

func GetDB() *gorm.DB {
	return db
}
