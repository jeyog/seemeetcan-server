package database

import (
	"log"
	"sync"

	"github.com/jeyog/seemeetcan-server/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	instance *Database
	once     sync.Once
)

type Database struct {
	dbConn *gorm.DB
}

func GetInstance() *Database {
	once.Do(func() {
		instance = newDatabase()
	})
	return instance
}

func newDatabase() *Database {
	dsn := "root:1024@tcp(127.0.0.1:3306)/seemeetcan?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(&model.User{})

	return &Database{db}
}

func (db *Database) GetDb() *gorm.DB {
	return db.dbConn
}
