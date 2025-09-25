package friendships

import (
	"project/internals/cache"
	"project/internals/dto"
	"project/models/friendships"
	"project/services/users"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Add a friend
func Add(c *fiber.Ctx) error {
	ctx := c.UserContext()
	var body struct {
		UserID   string `json:"user_id"`
		FriendID string `json:"friend_id"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	userID, err := uuid.Parse(body.UserID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid user_id"})
	}
	friendID, err := uuid.Parse(body.FriendID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid friend_id"})
	}

	// Check if both users exist
	us := users.New()
	user, err := us.Get(ctx, userID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "user not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	friendUser, err := us.Get(ctx, friendID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "friend not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Create friendship
	fs := friendships.New()
	fs.Friends = &dto.Friends{
		UserID:   user.ID,
		FriendID: friendUser.ID,
	}

	if err := fs.Create(ctx); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if err := cache.Client().Del(ctx, user.ID.String()).Err(); err != nil {
		// Log the error but don't fail the request
		fmt.Printf("Error clearing cache for user %s: %v\n", user.ID, err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Friend added successfully",
		"data":    fs.Friends,
	})
}

// Get all friends of a user
func GetAll(c *fiber.Ctx) error {
	ctx := c.UserContext()
	userID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid user id"})
	}

	fs := friendships.New()
	fs.UserID = userID
	if err := fs.GetAll(ctx); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fs.AllFriends)
}




// Get a specific user
func Get(c *fiber.Ctx) error {
	ctx := c.UserContext()
	id := c.Params("id")
	userID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid user id"})
	}

	us := users.New()
	user, err := us.Get(ctx, userID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "user not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

// Delete a friendship
func Delete(c *fiber.Ctx) error {
	ctx := c.UserContext()
	userID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid user id"})
	}

	friendID, err := uuid.Parse(c.Query("f_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid friend id"})
	}

	us := users.New()
	if _, err := us.Get(ctx, userID); err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "user not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not get user"})
	}

	if _, err := us.Get(ctx, friendID); err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "friend not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not get friend"})
	}

	fs := friendships.New()
	fs.UserID = userID
	fs.FriendID = friendID
	if err := fs.Delete(ctx); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if err := cache.Client().Del(ctx, fs.UserID.String()).Err(); err != nil {
		fmt.Printf("Error clearing cache for user %s: %v\n", fs.UserID, err)
	}
	return c.SendStatus(fiber.StatusNoContent)
}
