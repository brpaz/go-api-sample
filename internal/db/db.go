package db

import (
	"fmt"
	"github.com/brpaz/go-api-sample/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// GetConnection Returns a database connection
func GetConnection(config config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable",
		config.DB.Host,
		config.DB.User,
		config.DB.Password,
		config.DB.Driver,
		config.DB.Port)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
