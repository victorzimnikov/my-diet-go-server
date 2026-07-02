package database

import (
	"fmt"
	"log"
	"my-diet-server/config"
	"my-diet-server/internal/models"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	dbPort := config.Config("DB_PORT")
	dbUser := config.Config("DB_USER")
	dbPass := config.Config("DB_PASSWORD")
	dbName := config.Config("DB_NAME")
	dbHost := config.Config("DB_HOST")

	port, err := strconv.ParseUint(dbPort, 10, 32)

	if err != nil {
		log.Println("Idiot")
	}

	// Connection URL to connect to Postgres Database
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbHost, port, dbUser, dbPass, dbName)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  false,       // Disable color
		},
	)

	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		panic("failed to connect database")
	}

	DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")

	DB.Exec("DO ' BEGIN IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = ''product_unit_type'') THEN CREATE TYPE product_unit_type AS ENUM (''g'', ''kg'', ''ml'', ''l''); END IF; END ';")
	DB.Exec("DO ' BEGIN IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = ''menu_type'') THEN CREATE TYPE menu_type AS ENUM (''week'', ''two_week'', ''month''); END IF; END ';")
	DB.Exec("DO ' BEGIN IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = ''food_intake_type'') THEN CREATE TYPE food_intake_type AS ENUM (''breakfast'', ''first_snack'', ''launch'', ''afternoon_snack'', ''dinner''); END IF; END ';")
	DB.Exec("DO ' BEGIN IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = ''user_role_type'') THEN CREATE TYPE user_role_type AS ENUM (''user'', ''moderator'', ''admin''); END IF; END ';")
	DB.Exec("DO ' BEGIN IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = ''week_day_type'') THEN CREATE TYPE week_day_type AS ENUM (''mon'',''tue'',''wed'',''thu'',''fri'',''sat'',''sun''); END IF; END ';")

	log.Println("Connection Opened to Database")

	// Migrate the database
	migrateError := DB.AutoMigrate(
		&models.UserModel{},
		&models.ProductModel{},
		&models.RecipeModel{},
		&models.ImageModel{},
		&models.RecipeStepModel{},
		&models.RecipeIngredientModel{},
		&models.NutritionModel{},
	)

	if migrateError != nil {
		log.Println("Failed migrate:",
			migrateError.Error())
	}

	log.Println("Database Migrated")
}
