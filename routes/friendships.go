package routes

import (
	"project/controllers/friendships"

	"github.com/gofiber/fiber/v2"
)

func Friendships(r fiber.Router){
	friendshiproutes := r.Group("/friends")
	friendshiproutes.Post("/",friendships.Add)
	friendshiproutes.Get("/:id", friendships.GetAll)
	friendshiproutes.Get("/:id",friendships.Get)
	friendshiproutes.Delete("/:id",friendships.Delete)
}
