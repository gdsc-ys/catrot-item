package database

import (
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"src/models"
	"os"
	"github.com/joho/godotenv"
)


var DB *gorm.DB

func Init() {
	dotEnvErr := godotenv.Load()
	if dotEnvErr != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?parseTime=true"

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	DB.AutoMigrate(&models.Item{})

	if err != nil {
		panic("DB 연결이 실패했습니다!")
	}

	log.Println("연결이 성공적으로 ")
}
