package app

import (
	"fmt"
	"project/internals/database"
	"project/internals/server"
	"project/internals/notifications"
	"project/internals/cache"
)

func Setup() {
	var sqlSchemaFile string = "project/init/01_tables.sql"
	database.Connect(sqlSchemaFile)
	cache.Connect()
	notifications.InitNotificationsSystem()
	notifications.Hydrate()
	app := server.New() // Fiber instance with routes
	if err := app.Listen(":3015"); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
		panic(err)
	}
}
