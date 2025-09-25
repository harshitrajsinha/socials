package server

import ("github.com/gofiber/fiber/v2"
	    "project/routes")

func errorHandler(c *fiber.Ctx, err error) error {
	msg := err.Error()
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error":   true,
		"message": msg,
	})
}

var notFoundHandler = func(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"error":   true,
		"message": "Route not found",
	})
}

func addRoutes(app *fiber.App) {
	baseRouter := app.Group("/api")
	routes.Users(baseRouter)
	routes.Posts(baseRouter)
	routes.Friendships(baseRouter)
}