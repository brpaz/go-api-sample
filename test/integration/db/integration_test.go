// +build integrationdb

package main

import (
	"flag"
	"fmt"
	_ "github.com/DATA-DOG/go-txdb"
	"github.com/brpaz/go-api-sample/test/testutil"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"testing"
)

var logger *zap.Logger

func getConnection(dsn string) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func setupDB() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := fmt.Sprintf("%s_test", os.Getenv("DB_DATABASE"))

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable",
		dbHost,
		dbUser,
		dbPassword,
		"postgres",
		dbPort,
	)

	adminDb, err := getConnection(dsn)

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
		dbPort,
	)

	db, err := getConnection(dsn)
	if err := testutil.Migrate(db); err != nil {
		log.Fatalf("failed to run database migrations: %v", err)
	}
}

// TestMain is the entry point to the application database integration tests.
// It is responsible to setup the database for tests by creating the test database and run the migrations on the new schema and launch the tests.
func TestMain(m *testing.M) {

	flag.Parse()

	// setup database
	setupDB()

	// run tests
	os.Exit(m.Run())
}
