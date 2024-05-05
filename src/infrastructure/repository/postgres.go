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

const (
	// Local DB
	host     = os.Getenv("DB_HOST")
	port     = os.Getenv("DB_PORT")
	user     = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	dbname   = os.Getenv("DB_NAME")
)

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
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=require",
		host, port, user, password, dbname)
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic("Can't connect to database")
	}
	db.AutoMigrate(&model.User{}, &model.Product{}, &model.Category{}, &model.Cart{})

	return db
}
