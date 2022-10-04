package api

import (
	"fmt"
	"os"

	"github.com/bagasalim/simas/model"
	// "github.com/glebarez/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "simaS123"
	dbname   = "simascontact"
)

func SetupDb() (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	dbUrl := os.Getenv("DATABASE_URL")

	if os.Getenv("ENVIRONMENT") == "PROD" {
		db, err = gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	} else {
		config := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

		db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	}

	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database: %w", err)
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	if os.Getenv("AUTO_MIGRATE") == "Y" {
		if err := db.AutoMigrate(model.User{}); err != nil {
			return nil, fmt.Errorf("failed to migrate database: %w", err)
		}

		admin := model.User{
			Username: "admin",
			Password: "$2a$10$BQHCjmHmEsFGJXCGWm7et.2lvVPecg0ibhFd/tgOCCCncTu5ieiA.",
			Name:     "Administrator",
			Role:     1,
		}

		res := db.Create(&admin)
		if res == nil {
			return nil, fmt.Errorf("failed to seeding admin database: %w", res.Error)
		}

		cs := model.User{
			Username: "CS01",
			Password: "2a$10$BQHCjmHmEsFGJXCGWm7et.2lvVPecg0ibhFd/tgOCCCncTu5ieiA.",
			Name:     "Customer Service",
			Role:     2,
		}

		res1 := db.Create(&cs)
		if res1 == nil {
			return nil, fmt.Errorf("failed to seeding cs database: %w", res1.Error)
		}
	}

	return db, err
}
