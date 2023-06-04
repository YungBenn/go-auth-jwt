package database

import (
	"go-auth-jwt/internal/entity"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(connection string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(connection), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(
		&entity.User{},
	)

	if err != nil {
		log.Fatal(err)
	}

	return db
}
