package users

import (
	// "context"
	"context"
	"fmt"
	"project/internals/database"
	"project/internals/dto"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Users struct {
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	User *dto.User `gorm:"-"`
}

func New() *Users{
	return &Users{}
}

func (u *Users) Create() {
	if err := database.Client().Create(u).Error; err != nil {
		fmt.Print("failed to create user: ", err)
	}
}
func (u *Users) Get(ctx context.Context) error {
	if err := database.Client().First(&u, "id = ?", u.ID).Error; err != nil {
		if err == gorm.ErrRecordNotFound{
			fmt.Printf("user with id %s not found", u.ID)
			return err
		}
		fmt.Print("failed to get user: ", err)
	}
	return nil
}


func (u *Users) Delete(ctx context.Context) error {
	if err := database.Client().Delete(&u.User ,u.ID).Error; err != nil {
		if err == gorm.ErrRecordNotFound{
			fmt.Printf("user with id %s not found", u.ID)
			return err
		}
		fmt.Print("failed to delete user: ", err)
	}
	return nil
}
func (u *Users) GetAll() ([]Users, error) {
	var users []Users
	if err := database.Client().Find(&users).Error; err != nil {
		fmt.Print("failed to get users: ", err)
		return nil, err
	}
	return users, nil
}