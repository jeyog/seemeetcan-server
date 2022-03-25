package route

import (
	"crypto/rsa"

	"github.com/gofiber/fiber/v2"
	"github.com/jeyog/seemeetcan-server/controller"
	"gorm.io/gorm"
)

type Route struct {
	app            *fiber.App
	db             *gorm.DB
	privateKey     *rsa.PrivateKey
	authController *controller.AuthController
}

func New(app *fiber.App, db *gorm.DB, privateKey *rsa.PrivateKey) *Route {
	return &Route{
		app,
		db,
		privateKey,
		&controller.AuthController{},
	}
}

func (r *Route) Setup() {
	r.setupApi()
}
