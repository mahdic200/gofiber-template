package Config

import (
    "gorm.io/driver/postgres"
    "github.com/joho/godotenv"
    "gorm.io/gorm"
    "os"
    "fmt"
    "fiber/Models"
)

var DB *gorm.DB
func Connect() {

    godotenv.Load()
    dbname := os.Getenv("PG_DBNAME")
    dbhost := os.Getenv("PG_HOST")
    dbport := os.Getenv("PG_PORT")
    dbuser := os.Getenv("PG_USER")
    dbpass := os.Getenv("PG_PASSWORD")

    connection := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tehran", dbhost, dbuser, dbpass, dbname, dbport)
    db, err := gorm.Open(postgres.Open(connection), &gorm.Config{
        SkipDefaultTransaction: true,
    })
    
    if err != nil {
        panic("db connection failed")
    }

    DB = db

    db.AutoMigrate(
        &models.User {},
        &models.Post {},
    )
}
