package repository

import (
	"GolangwithFrame/src/domain/model"
	"fmt"
	"github.com/jinzhu/gorm"
	"os"
)

type Database struct {
	Connection *gorm.DB
}

func NewRepository() Database {

	db := NewDB()
	return Database{
		Connection: db,
	}

}

func (db *Database) CloseDB() {
	err := db.Connection.Close()
	if err != nil {
		panic("Failed to Connect")
	}

}

func NewDB() *gorm.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=require",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))
	fmt.Println(psqlInfo)
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic("Can't connect to database")
	}
	db.AutoMigrate(&model.User{}, &model.Product{}, &model.Category{}, &model.Cart{})
    print("here")
	return db
}
