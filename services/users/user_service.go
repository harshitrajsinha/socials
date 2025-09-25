package users

import (
	"context"
	"project/internals/dto"
	"project/models/users"
	"project/internals/database"
	"fmt"
	"github.com/google/uuid"
)

type User struct {
	User *dto.User
}

func New() *User {
	return &User{}
}

func (u *User) Create() {
	m := users.New()
	m.Name = u.User.Name
	m.Email = u.User.Email
	m.Password = u.User.Password

	if err := database.Client().Create(m).Error; err != nil {
		fmt.Print("failed to create user: ", err)
		return
	}

	// Assign generated ID and CreatedAt back
	u.User.ID = m.ID
	u.User.CreatedAt = &m.CreatedAt
}

func (u *User) Get(ctx context.Context, id uuid.UUID) (*dto.User, error) {
    m := users.New()
    m.ID = id
    if err := m.Get(ctx); err != nil {
        return nil, err
    }

    return &dto.User{
        ID:        m.ID,
        Name:      m.Name,
        Email:     m.Email,
        CreatedAt: &m.CreatedAt,
        UpdatedAt: &m.UpdatedAt,
    }, nil
}

func (u *User) Delete(ctx context.Context) error {
	m := users.New()
	m.ID = u.User.ID
	m.User = u.User
	if err := m.Delete(ctx); err != nil {
		return err
	}
	return nil
}
func (u *User) GetAll() ([]*dto.User, error) {
	var usersModel users.Users
	usersList, err := usersModel.GetAll() // calls your model's GetAll
	if err != nil {
		return nil, err
	}

	var dtoList []*dto.User
	for _, uModel := range usersList {
		dtoList = append(dtoList, &dto.User{
			ID:        uModel.ID,
			Name:      uModel.Name,
			Email:     uModel.Email,
			CreatedAt: &uModel.CreatedAt,
			UpdatedAt: &uModel.UpdatedAt,
		})
	}
	return dtoList, nil
}
