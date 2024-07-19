package database

import (
    "fmt"
    "todo-app/models"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func SetupDatabase() {
    var err error
    dsn := "host=localhost user=doni password=123456 dbname=todoapp port=5432 sslmode=disable"
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    err = DB.AutoMigrate(&models.User{}, &models.Todo{})
    if err != nil {
        fmt.Println("Failed to migrate database: ", err)
    }
}
