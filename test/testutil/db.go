package testutil

import (
	"fmt"
	"github.com/brpaz/go-api-sample/internal/util"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	gormPG "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"sync"
)

var (
	dbConn *gorm.DB
	once   sync.Once
)

// SetupDB setup test database for integration and acceptance tests
func SetupDB() (*gorm.DB, error) {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := GetTestDBName()

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable",
		dbHost,
		dbUser,
		dbPassword,
		"postgres",
		dbPort)

	adminDb, err := gorm.Open(gormPG.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}

	if err := CreateDB(adminDb, dbName); err != nil {
		return nil, fmt.Errorf("failed to create test db: %w", err)
	}

	dsn = fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable",
		dbHost,
		dbUser,
		dbPassword,
		dbName,
		dbPort)

	db, err := gorm.Open(gormPG.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("cannot open database connection: %w", err)
	}

	if err := Migrate(db); err != nil {
		return nil, fmt.Errorf("error running database migrations: %w", err)
	}

	return db, nil
}

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
		fmt.Sprintf("file://%s/migrations", util.GetRootDir()),
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

// GetTestDBName returns the name of the test DB.
func GetTestDBName() string {
	return fmt.Sprintf("%s_test", os.Getenv("DB_DATABASE"))
}

// GetConnection Returns the connection to the test db
func GetTestDBConnection() *gorm.DB {

	once.Do(func() {
		dbHost := os.Getenv("DB_HOST")
		dbPort := os.Getenv("DB_PORT")
		dbUser := os.Getenv("DB_USER")
		dbPassword := os.Getenv("DB_PASSWORD")
		dbName := GetTestDBName()

		dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable",
			dbHost,
			dbUser,
			dbPassword,
			dbName,
			dbPort,
		)

		db, err := gorm.Open(gormPG.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}

		dbConn = db
	})

	return dbConn
}
