// package posts

// import (
// 	"context"
// 	"project/models/posts"
// 	// "project/controllers/posts"
// 	// "project/internals/dto"
// )

// type Posts struct {

// }

// func New() *Posts {
// 	return &Posts{}
// }

// func (p *Posts) Create(ctx context.Context) {
// 	m := posts.New()
// 	m.Post = p.Post
// 	m.Create(ctx)
// 	p.Post.UpdatedAt = nil
// }

// func (p *Posts) GetAll(ctx context.Context){
// 	m := posts.New()
// 	m.UserID = p.UserID
// 	m.Posts =p.posts
// 	m.Get(ctx)
// }
// func (p *Posts) Delete(ctx context.Context)error{
// 	m := posts.New()
// 	m.UserID = p.UserID
// 	m.ID = p.ID
// 	if err := m.Delete(ctx); err != nil {
// 		return err
// 	}
// 	return nil
// }
package posts

import (
	"context"
	"project/internals/dto"
	"project/models/posts"

	"github.com/google/uuid"
)

type Posts struct {
	Post   *dto.Post
	Posts  *[]dto.Post
	ID     uuid.UUID
	UserID uuid.UUID
}

func New() *Posts {
	return &Posts{}
}

func (p *Posts) Create(ctx context.Context) error {
	m := posts.New()
	m.Content = p.Post.Content
	m.UserID = p.Post.UserID

	if err := m.Create(ctx); err != nil {
		return err
	}

	p.Post.ID = m.ID
	p.Post.CreatedAt = &m.CreatedAt
	p.Post.UpdatedAt = &m.UpdatedAt
	return nil
}

func (p *Posts) GetAll(ctx context.Context) error {
	m := posts.New()
	var all []posts.Posts
	if err := m.Get(ctx, p.UserID, &all); err != nil {
		return err
	}

	var result []dto.Post
	for _, item := range all {
		result = append(result, dto.Post{
			ID:        item.ID,
			Content:   item.Content,
			UserID:    item.UserID,
			CreatedAt: &item.CreatedAt,
			UpdatedAt: &item.UpdatedAt,
		})
	}
	p.Posts = &result
	return nil
}

func (p *Posts) Delete(ctx context.Context) error {
	m := posts.New()
	return m.Delete(ctx, p.UserID, p.ID)
}
