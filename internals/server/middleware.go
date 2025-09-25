package server
import (
	"github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/logger"
)

func Middleware(app *fiber.App) {
	app.Use(logger.New())
	app.Use(func(c *fiber.Ctx) error {
		// Add your middleware logic here
		return c.Next()
	})
}
