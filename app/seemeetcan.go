package app

import (
	"crypto/rand"
	"crypto/rsa"
	"log"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/jeyog/seemeetcan-server/model"
	"github.com/jeyog/seemeetcan-server/route"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	instance *SeeMeetCan
	once     sync.Once
)

type SeeMeetCan struct {
	app        *fiber.App
	db         *gorm.DB
	privateKey *rsa.PrivateKey
}

func GetInstance() *SeeMeetCan {
	once.Do(func() {
		instance = newApp()
	})
	return instance
}

func newApp() *SeeMeetCan {
	app := fiber.New()

	dsn := "root:1024@tcp(127.0.0.1:3306)/seemeetcan?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(&model.User{})

	rng := rand.Reader
	privateKey, err := rsa.GenerateKey(rng, 2048)
	if err != nil {
		log.Fatalf("rsa.GenerateKey: %v", err)
	}

	return &SeeMeetCan{app, db, privateKey}
}

func (s *SeeMeetCan) Run() {
	route := route.New(s.app, s.db, s.privateKey)
	route.Setup()

	s.app.Listen(":80")
}

func (s *SeeMeetCan) GetApp() *fiber.App {
	return s.app
}

func (s *SeeMeetCan) GetDb() *gorm.DB {
	return s.db
}

func (s *SeeMeetCan) GetPrivateKey() *rsa.PrivateKey {
	return s.privateKey
}
