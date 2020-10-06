package main

import (
	"flag"
	"fmt"
	_ "github.com/DATA-DOG/go-txdb"
	"github.com/brpaz/go-api-sample/test/testutil"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"os/exec"
)

// prepares the test database
func setupDB()  {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := testutil.GetTestDBName()

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable",
		dbHost,
		dbUser,
		dbPassword,
		"postgres",
		dbPort)

	adminDb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}

	if err := testutil.CreateDB(adminDb, dbName); err != nil {
		log.Fatalf("failed to create test db: %v", err)
	}

	dsn = fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable",
		dbHost,
		dbUser,
		dbPassword,
		dbName,
		dbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err := testutil.Migrate(db); err != nil {
		log.Fatalf("failed to run database migrations: %v", err)
	}
}

func runTests() error {
	cmd := exec.Command("go", "test", "-v", "--tags", "integrationdb", "-p", "1", "./...")
	cmd.Env = os.Environ()
	cmd.Stdout = os.Stdout

	if err := cmd.Start(); err != nil {
		return err
	}

	if err := cmd.Wait(); err != nil {
		return err
	}

	return nil
}

// TestMain is the entry point to the application database integration tests.
// It is responsible to setup the database for tests by creating the test database and run the migrations on the new schema and launch the tests.
func main() {
	flag.Parse()

	if os.Getenv("APP_ENV") != "test" {
		log.Fatal("You can only run this command with APP_ENV=test")
	}

	log.Println("Setup DB")

	setupDB()

	log.Println("Running Tests")

	if err := runTests(); err != nil {
		os.Exit(-1)
	}
}
