package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"project/routes" // make sure this import is active
)

var app *fiber.App

func New() *fiber.App {
	if app == nil {
		app = fiber.New(fiber.Config{
			ErrorHandler: errorHandler,
			BodyLimit:    10 * 1024 * 1024,
		})

		// Middlewares
		app.Use(recover.New())
		app.Use(logger.New())

		// Register routes here
		routes.Users(app)
		routes.Friendships(app)
		routes.Posts(app)

		// 404 handler
		app.Use(notFoundHandler)
	}

	return app
}
