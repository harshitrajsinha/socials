package database

import (
	"fmt"
	"os"
	"time"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func Client() *gorm.DB {
	return DB
}

func loadDataToDatabase(dbClient *gorm.DB, filename string) error {

	// Read file content
	sqlFile, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	// Execute file content (queries)
	result := dbClient.Exec(string(sqlFile))
	if result.Error != nil {
	panic(result.Error)
}
	return nil
}

// Connect opens the DB and waits for it to be ready
func Connect(sqlfile string) {
		_ = godotenv.Load()
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
		err = loadDataToDatabase(db, sqlfile)
	if err != nil {
		panic(err)
	} else {
		log.Println("SQL file executed successfully!")
	}
		return
	}

	// If we reach here, all retries failed
	panic(fmt.Sprintf("Could not connect to database after %d attempts: %v", maxRetries, err))
}
