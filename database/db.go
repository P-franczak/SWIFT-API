package database

import (
    "fmt"
    "log"
    "os"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "github.com/P-franczak/swift-api/models"
)

var DB *gorm.DB

func ConnectDatabase() {
    dsn := os.Getenv("DATABASE_URL")
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Nie udało się połączyć z bazą danych:", err)
    }

    db.AutoMigrate(&models.SwiftCode{})
    DB = db
    fmt.Println("Połączono z bazą danych")
}
