package db

import (
	"fmt"
	"github.com/brpaz/go-api-sample/internal/config"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"moul.io/zapgorm2"
)

// GetConnection Returns a database connection
func GetConnection(cfg config.Config, logger *zap.Logger) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable",
		cfg.DB.Host,
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Database,
		cfg.DB.Port)

	l := zapgorm2.New(logger)

	if cfg.Env == config.EnvDev {
		l.LogMode(gormLogger.Info)
	}

	l.SetAsDefault() // optional: configure gorm to use this zapgorm.Logger for callbacks
	return gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: l,
	})
}
