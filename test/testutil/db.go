package testutil

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	gormPG "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

// Migrate Runs database migrations on the test db
func Migrate(db *gorm.DB) error {

	log.Println("Running initial migrations")

	dbInstance, err := db.DB()

	if err != nil {
		return err
	}

	driver, err := postgres.WithInstance(dbInstance, &postgres.Config{})

	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://../../../migrations",
		"postgres", driver)

	if err != nil {
		return err
	}

	return m.Up()
}

// CreateTestDB Creates a database for testing purposes
func CreateDB(db *gorm.DB, dbName string) error {

	log.Println(fmt.Sprintf("Creating test database %s", dbName))

	stmt := fmt.Sprintf("DROP DATABASE IF EXISTS %s", dbName)
	if err := db.Exec(stmt).Error; err != nil {
		return err
	}

	stmt = fmt.Sprintf("CREATE DATABASE %s", dbName)
	return db.Exec(stmt).Error
}

// GetConnection Returns the connection to the test db
func GetTestDBConnection() (*gorm.DB, error) {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := fmt.Sprintf("%s_test", os.Getenv("DB_DATABASE"))

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable",
		dbHost,
		dbUser,
		dbPassword,
		dbName,
		dbPort,
	)

	return gorm.Open(gormPG.Open(dsn), &gorm.Config{})
}
