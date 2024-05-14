package config

import (
	"fmt"
	"log"
	"os"

	"github.com/CRobinDev/Gemastik/entity"
	"github.com/CRobinDev/Gemastik/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func LoadDatabaseConfig() string {
	DataSourceName := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable ",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	return DataSourceName
}

func ConnectToDB() (*gorm.DB, error) {

	db, err := gorm.Open(postgres.Open(LoadDatabaseConfig()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Fatalf("[ayam] Failed to connect to database : %v\n", err)
		return nil, errors.ErrConnectDatabase
	}
	return db, nil
}

func Migrate(db *gorm.DB) error {
	db = db.Debug()
	if err := db.AutoMigrate(
		&entity.User{},
		&entity.Chat{},
		&entity.Leaderboard{},
	); err != nil {
		return errors.ErrMigrateDatabase
	}
	return nil
}
