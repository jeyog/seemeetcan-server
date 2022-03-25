package controller

import (
	fiber "github.com/gofiber/fiber/v2"
	"github.com/jeyog/seemeetcan-server/service"
)

type AuthController struct {
	authService *service.AuthService
}

type AuthRequest struct {
	Username string
	Password string
}

func (a *AuthController) Authenticate(ctx *fiber.Ctx) error {
	return ctx.SendString("Authenticate")
}
