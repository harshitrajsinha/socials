// package posts

// import (
// 	"context"
// 	"fmt"
// 	"project/internals/database"
// 	"project/models/users"
// 	"time"

// 	"github.com/google/uuid"
// 	"gorm.io/gorm"
// )

// type Posts struct {
// 	ID        uuid.UUID   `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
// 	Title     string      `json:"title"`
// 	Content   string      `json:"content"`
// 	UserID    uuid.UUID   `json:"user_id"`
// 	CreatedAt time.Time   `json:"created_at"`
// 	UpdatedAt time.Time   `json:"updated_at"`
// 	User      users.Users `gorm:"foreignKey:UserID;references:ID" json:"user"`
// }

// func New() *Posts {
// 	return &Posts{}
// }

// func (p *Posts) Create(ctx context.Context) error {
// 	if err := database.Client().Table("posts").Create(&p.Post).Error; err != nil {
// 		fmt.Print("Unable to create post:%v", err)
// 		return err
// 	}
// 	return nil
// }


// func (p *Posts) Get(ctx context.Context) error {
// 	if err := database.Client().Table("posts").Where("user_id = ?", p.UserID).Find(&p.Posts.Posts).Error; err != nil {
// 		fmt.Print("Unable to fetch posts:%v", err)
// 		return err
// 	}
// 	return nil
// }


// // func (p *Post) GetAll(ctx context.Context) ([]Post, error) {
// // 	var posts []Post
// // 	err := database.Client().WithContext(ctx).Preload("User").Find(&posts).Error
// // 	return posts, err
// // }

// func (p *Posts) Delete(ctx context.Context) error {
// 	if err := database.Client().Where("user_id = ?", p.UserID).Where("id = ?", p.ID).Delete(p).Error; err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			fmt.Printf("Error getting user: %v", err)
// 		}
// 	}
// 	return nil
// }

package posts

import (
	"context"
	"project/internals/database"
	"project/models/users"
	"time"

	"github.com/google/uuid"
)

type Posts struct {
	ID        uuid.UUID   `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Content   string      `json:"content"`
	UserID    uuid.UUID   `json:"user_id"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	User      users.Users `gorm:"foreignKey:UserID;references:ID" json:"user"`
}

func New() *Posts {
	return &Posts{}
}

func (p *Posts) Create(ctx context.Context) error {
	return database.Client().WithContext(ctx).Create(p).Error
}

func (p *Posts) Get(ctx context.Context, userID uuid.UUID, out *[]Posts) error {
	return database.Client().WithContext(ctx).
		Where("user_id = ?", userID).
		Preload("User").
		Find(out).Error
}

func (p *Posts) Delete(ctx context.Context, userID, postID uuid.UUID) error {
	return database.Client().WithContext(ctx).
		Where("user_id = ? AND id = ?", userID, postID).
		Delete(&Posts{}).Error
}
