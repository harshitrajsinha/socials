package routes

import (
	"github.com/gofiber/fiber/v2"
	"project/controllers/posts"
)

func Posts(r fiber.Router){
	postroutes := r.Group("/users/:id/posts")
	postroutes.Post("",posts.Add)
	postroutes.Get("",posts.Get)
	postroutes.Delete("/:post_id",posts.Delete)
}