package db

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func initDB() {

	DBUsername := os.Getenv("DBUsername")
	DBPassword := os.Getenv("DBPassword")
	DBUrl := os.Getenv("DBHost")
	DBPort := os.Getenv("DBPort")
	DBName := os.Getenv("DBName")

	dbCONN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DBUsername, DBPassword, DBUrl, DBPort, DBName)

	var err error

	DB, err := gorm.Open(mysql.Open(dbCONN), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	DB.AutoMigrate()
}
