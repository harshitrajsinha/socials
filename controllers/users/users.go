package users

import (
	"project/internals/dto"
	"project/internals/notifications"
	"project/internals/validator"
	serviceUsers "project/services/users"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func Add(c *fiber.Ctx) error{
	// ctx := c.Context()
	var user dto.UserCreate

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}
	if err := validator.Users(user);err != nil{
		return c.Status(fiber.StatusBadRequest).JSON("incorrect input body")
	}
	us := serviceUsers.New()
	us.User = &dto.User{}
	us.User.Name = user.Name
	us.User.Email = user.Email
	us.User.Password = user.Password
	us.Create()
	notifications.Register(us.User.ID)
	go notifications.ListenForNotifications(c.UserContext(),us.User.ID)
	return c.Status(fiber.StatusCreated).JSON(us.User)
}

func Get(c *fiber.Ctx) error {
	ctx := c.UserContext()
	id := c.Params("id")
	userID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid user id"})
	}

	us := serviceUsers.New()
	user, err := us.Get(ctx, userID) // <-- get both values
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "user not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}


func Delete(c *fiber.Ctx) error{
	ctx := c.UserContext()
	id := c.Params("id")
	userID, err := uuid.Parse(id)
	if err != nil {
		return  c.Status(fiber.StatusBadRequest).JSON("invalid user id")
	}
	us := serviceUsers.New()
	us.User = &dto.User{}
	us.User.ID = userID
	if err := us.Delete(ctx); err != nil {
		if err == gorm.ErrRecordNotFound{
			return c.Status(fiber.StatusNotFound).JSON("user not found")
		}
		return c.Status(fiber.StatusInternalServerError).JSON("could not delete user")
	}
	return c.Status(fiber.StatusOK).JSON("user deleted")
}
func GetAll(c *fiber.Ctx) error {
	service := serviceUsers.New()
	users, err := service.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "failed to get users",
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"data":   users,
	})
}