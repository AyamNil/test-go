package database

import (
    "log"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "example.com/comments-api/models"
)

var DB *gorm.DB

func Connect() {
    dsn := "root:@tcp(127.0.0.1:3306)/comments_db?charset=utf8mb4&parseTime=True&loc=Local"
    var err error
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to the database:", err)
    }

    // Run migrations
    DB.AutoMigrate(&models.Comment{})
    log.Println("Database connected and migrated successfully.")
}
