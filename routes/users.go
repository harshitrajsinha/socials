package routes

import (
	"github.com/gofiber/fiber/v2"
	userController "project/controllers/users"
)

func Users(r fiber.Router) {
	users := r.Group("/users")

	users.Get("/", userController.GetAll)
	users.Post("/", userController.Add)
	users.Get("/:id", userController.Get)
	users.Put("/:id", func(c *fiber.Ctx) error {
		return c.SendString("Update user") // implement real update later
	})
	users.Delete("/:id", userController.Delete)
}
