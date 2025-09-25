package database

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Client() *gorm.DB {
	return DB
}

// Connect opens the DB and waits for it to be ready
func Connect() {
	dsn := os.Getenv("DATABASE_DSN")
	time.Sleep(2 * time.Second)
	if dsn == "" {
		panic("DATABASE_DSN is not set") // don’t default, let compose handle it
	}

	var db *gorm.DB
	var err error
	maxRetries := 10
	for i := 0; i < maxRetries; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			fmt.Printf("Attempt %d: failed to open DB: %v\n", i+1, err)
			time.Sleep(time.Second * 2)
			continue
		}

		sqlDB, err := db.DB()
		if err != nil {
			fmt.Printf("Attempt %d: failed to get sql.DB: %v\n", i+1, err)
			time.Sleep(time.Second * 2)
			continue
		}

		if err := sqlDB.Ping(); err != nil {
			fmt.Printf("Attempt %d: failed to ping DB: %v\n", i+1, err)
			time.Sleep(time.Second * 2)
			continue
		}

		// Success
		DB = db
		fmt.Println("Successfully Connected to Postgres - Database🧑🏼‍💻")
		return
	}

	// If we reach here, all retries failed
	panic(fmt.Sprintf("Could not connect to database after %d attempts: %v", maxRetries, err))
}
