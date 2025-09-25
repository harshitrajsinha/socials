package posts

import (
	"project/internals/dto"
	"project/internals/notifications"
	"project/internals/validator"
	servicePosts "project/services/posts"
	serviceUsers "project/services/users"

	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Add creates a new post
func Add(c *fiber.Ctx) error {
	ctx := c.UserContext()
	var post dto.PostCreate

	// parse user ID
	id := c.Params("id")
	userID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Invalid user ID")
	}

	// parse body
	if err := c.BodyParser(&post); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Incorrect post format")
	}

	// validate payload
	if err := validator.Payload(post); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Incorrect post format")
	}

	// ensure user exists
	us := serviceUsers.New()
	user, err := us.Get(ctx, userID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON("User not found")
		}
		return c.Status(fiber.StatusInternalServerError).JSON("Could not fetch user")
	}

	// create post
	ps := servicePosts.New()
	ps.Post = &dto.Post{
		Content: post.Content,
		UserID:  user.ID,
	}
	if err := ps.Create(ctx); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("Could not create post")
	}
	msg := fmt.Sprintf("Hello, your friend %v has created a new post", user.Name)
    notifications.NotifyUser(ctx, userID, msg)

	return c.Status(fiber.StatusCreated).JSON(ps.Post)
}

// Get retrieves all posts for a user
func Get(c *fiber.Ctx) error {
	ctx := c.UserContext()
	id := c.Params("id")
	userID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid user id"})
	}

	ps := servicePosts.New()
	ps.UserID = userID
	ps.Posts = &[]dto.Post{}
	if err := ps.GetAll(ctx); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(ps.Posts)
}

// Delete removes a specific post for a user
func Delete(c *fiber.Ctx) error {
	ctx := c.UserContext()

	userID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Invalid user ID")
	}

	postID, err := uuid.Parse(c.Params("post_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Invalid post ID")
	}

	// ensure user exists
	us := serviceUsers.New()
	_, err = us.Get(ctx, userID) // ✅ capture both values
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON("User not found")
		}
		return c.Status(fiber.StatusInternalServerError).JSON("Could not fetch user")
	}

	// delete post
	ps := servicePosts.New()
	ps.UserID = userID
	ps.ID = postID
	if err := ps.Delete(ctx); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("Could not delete post")
	}

	return c.Status(fiber.StatusOK).JSON("Post deleted")
}
