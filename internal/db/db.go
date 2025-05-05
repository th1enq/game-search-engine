package db

import (
	"fmt"
	"log"

	"github.com/one-piece-search-engine/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct {
	*gorm.DB
}

func LoadDB(config *config.Config) *DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s name=%s sslmode=disabled",
		config.DB.Host,
		config.DB.Port,
		config.DB.User,
		config.DB.Password,
		config.DB.Name,
	)

	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("failed to loading database: %v", err)
	}
	return &DB{gormDB}
}
