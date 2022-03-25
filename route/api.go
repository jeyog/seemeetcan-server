package route

import (
	jwtware "github.com/gofiber/jwt/v3"
)

func (r *Route) setupApi() {
	api := r.app.Group("/api")
	v1 := api.Group("/v1")
	{
		auth := v1.Group("/auth")
		auth.Post("/authenticate", r.authController.Authenticate)
	}

	api.Use(jwtware.New(jwtware.Config{
		SigningMethod: "RS256",
		SigningKey:    r.privateKey.Public(),
	}))
}
