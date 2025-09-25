package friendships

import (
	"context"
	"encoding/json"
	"fmt"
	"project/internals/cache"
	"project/internals/dto"
	"project/models/friendships"
	"time"
	"github.com/redis/go-redis/v9"
	"github.com/google/uuid"
)
type Friend struct {
	UserID uuid.UUID
	FriendID uuid.UUID
	Friends *dto.Friends
	AllFriends []dto.AllFriends
}


func New() *Friend {
	return &Friend{}
}
func (f *Friend) Create(ctx context.Context){
	m := friendships.New()
	m.Friends = f.Friends
	m.Create(ctx)
	f.Friends.CreatedAt = m.Friends.CreatedAt
}


// func (f *Friend) GetAll(ctx context.Context){
// 	val , err := cache.Client().Get(ctx, f.UserID.String()).Result()
// 	if val != "" && err == nil {
// 		json.Unmarshal([]byte(val), &f.AllFriends)
// 		return
// 	}
// 	m := friendships.New()
// 	m.UserID = f.UserID
// 	m.Get(ctx)

// 	f.AllFriends = m.AllFriends
// 	b,_ := json.Marshal(f.AllFriends)
// 	if err :=cache.Client().Set(ctx, f.UserID.String(), b, 24*time.Hour).Err(); err != nil {
// 		fmt.Println("Error setting up cache for friendships", err)
		
// 	}
// }
func (f *Friend) GetAll(ctx context.Context) {
	val, err := cache.Client().Get(ctx, f.UserID.String()).Result()
	if err == nil {
		// Cache hit
		if err := json.Unmarshal([]byte(val), &f.AllFriends); err != nil {
			fmt.Printf("[CACHE HIT] Failed to unmarshal cache for user %s: %v\n", f.UserID, err)
		} else {
			fmt.Printf("[CACHE HIT] Served friendships for user %s from Redis\n", f.UserID)
		}
		return
	} else if err != redis.Nil {
		// Real Redis error
		fmt.Printf("[CACHE ERROR] Could not fetch cache for user %s: %v\n", f.UserID, err)
	}

	// Cache miss → fetch from DB
	m := friendships.New()
	m.UserID = f.UserID
	m.Get(ctx)

	f.AllFriends = m.AllFriends

	// Store in cache
	b, err := json.Marshal(f.AllFriends)
	if err != nil {
		fmt.Printf("[CACHE SET ERROR] Failed to marshal friendships for user %s: %v\n", f.UserID, err)
		return
	}

	if err := cache.Client().Set(ctx, f.UserID.String(), b, 24*time.Hour).Err(); err != nil {
		fmt.Printf("[CACHE SET ERROR] Could not set cache for user %s: %v\n", f.UserID, err)
		return
	}

	fmt.Printf("[CACHE MISS] Fetched friendships for user %s from DB and cached\n", f.UserID)
}


func (u *Friend) Delete(ctx context.Context) error {
	m := friendships.New()
	m.UserID = u.Friends.UserID
	m.FriendID = u.Friends.FriendID
	if err := m.Delete(ctx); err != nil {
		return err
	}
	return nil
}