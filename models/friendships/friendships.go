package friendships

import (
	"context"
	"fmt"
	"project/internals/database"
	"project/internals/dto"
	"project/models/users"

	// ""
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Friendship struct {
	gorm.Model
	UserID     uuid.UUID `gorm:"uniqueIndex:idx_user_friend" json:"user_id"`
	FriendID   uuid.UUID  `gorm:"uniqueIndex:idx_user_friend" json:"friend_id"`

	User users.Users `gorm:"foreignKey:UserID;references:ID" json:"-"`
	Friend users.Users `gorm:"foreignKey:FriendID;references:ID" json:"-"`
	Friends *dto.Friends `gorm:"-"`
	AllFriends []dto.AllFriends `gorm:"-"`
}
func New() *Friendship {
	return &Friendship{}
}

func (f *Friendship) Create(ctx context.Context) error {
    // Check if friendship already exists
    var count int64
    err := database.Client().
        Table("friendships").
        Where("user_id = ? AND friend_id = ?", f.Friends.UserID, f.Friends.FriendID).
        Count(&count).Error
    if err != nil {
        return err
    }

    if count > 0 {
        // Already exists, just return without creating
        return fmt.Errorf("friendship already exists")
    }

    // Create new friendship
    if err := database.Client().Table("friendships").Create(&f.Friends).Error; err != nil {
        fmt.Printf("Unable to Create User: %v\n", err)
        return err
    }
    return nil
}


func (f *Friendship) Get(ctx context.Context) error {
	if err := database.Client().Table("friendships").Where("user_id = ?", f.UserID).Find(&f.AllFriends).Error; err != nil {
		fmt.Printf("Unable to get friends: %v", err)
		return err
	}
	return nil
}

func (f *Friendship) GetAll(ctx context.Context) error {
	var friends []dto.AllFriends
	if err := database.Client().
		Table("friendships").
		Select("friend_id").
		Where("user_id = ?", f.UserID).
		Find(&friends).Error; err != nil {
		return err
	}
	f.AllFriends = friends
	return nil
}



func (f *Friendship) Delete(ctx context.Context) error {
	if err := database.Client().
		Where("user_id = ?", f.UserID).
		Where("friend_id = ?", f.FriendID).
		Unscoped().
		Delete(f).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Printf(("Error Getting User: %v"), err)
			return err
		}
	}
	return nil
}
