package notifications

import (
	"fmt"
	"log"
	"project/services/friendships"
	serviceUsers "project/services/users"

	// "os/user"
	"context"
	"sync"

	"github.com/google/uuid"
)
var mu sync.Mutex
var Store map[uuid.UUID]chan string

func InitNotificationsSystem(){
	Store = make(map[uuid.UUID]chan string)
}

func Register(userID uuid.UUID){
	mu.Lock()
	defer mu.Unlock()

	if _,ok := Store[userID]; !ok {
		Store[userID] = make(chan string)
	}
}

func ListenForNotifications(ctx context.Context, userID uuid.UUID) {
    mu.Lock()
    channel, ok := Store[userID]
    mu.Unlock()

    if !ok {
        fmt.Printf("No channel found for user %s\n", userID)
        return
    }
	us := serviceUsers.New()
	user, err := us.Get(ctx, userID) // <-- get both values
	if err != nil {
		return
	}

    for {
        select {
        case message := <-channel:
            fmt.Printf("Hey %v you have a notification: %v\n", user.Name, message)

        case <-ctx.Done():
            fmt.Printf("Stopping notification listener for user %s\n", userID)
            return
        }
    }
}


func NotifyUser(ctx context.Context, userID uuid.UUID, msg string){
	fs := friendships.New()
	fs.UserID = userID
	fs.GetAll(ctx)

	mu.Lock()
	defer mu.Unlock()
	for _,f := range fs.AllFriends{
		if ch , ok := Store[f.FriendID]; ok {
			go func(ch chan string){
				ch <- msg
			}(ch)
		}
	}
}

func Hydrate(){

	service := serviceUsers.New()
	users, err := service.GetAll()
	if err != nil {
		log.Fatalf("Failed to hydrate notifications system: %v", err)
	}

	for _, user := range users {
		Register(user.ID)
		go ListenForNotifications(context.Background(), user.ID)
	}
}
