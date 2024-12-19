package app

import (
	"github.com/gofiber/fiber/v2"
	"karabayyazilim/src/internal/config"
	"karabayyazilim/src/internal/routes"
)

func SetupApp() *fiber.App {
	config.AppConfig()

	app := fiber.New()

	routes.ApiRoute(app)

	err := app.Listen(":" + config.Env().AppPort)

	if err != nil {
		return nil
	}

	return app
}
