package config

import (
	"project/internals/database"
	friendships "project/models/friendships"
	"project/models/posts"
	"project/models/users"
)

func AutoMigrate(){
	database.Client().AutoMigrate(&users.Users{}, &friendships.Friendship{}, &posts.Posts{})
}