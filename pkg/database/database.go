package database

import (
	"log"
	"os"

	"github.com/alexnavarrova/myfiberapp/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
    var err error
    dsn := os.Getenv("DATABASE_URL")
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Error connecting to database: ", err)
    }
}

func AutoMigrate() {
    err := DB.AutoMigrate(&model.User{}, &model.Book{}, &model.Library{})
    if err != nil {
        log.Fatal("Error migrating database: ", err)
    }
}