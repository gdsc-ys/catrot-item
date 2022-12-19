package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Title    string
	Category string
	Price    int
	Content  string
}

var DB *gorm.DB

func Init() {
	var err error
	dsn := "test:asdf0102!@tcp(localhost:3306)/mydb?parseTime=true"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	DB.AutoMigrate(&Item{})

	if err != nil {
		panic("DB 연결이 실패했습니다!")
	}

	fmt.Println("연결이 성공적으로 ")
}
