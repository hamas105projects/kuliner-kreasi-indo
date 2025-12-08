package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	DB        *gorm.DB
	JWTSecret string
	JWTExpire int
	AppPort   string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, reading from environment variables")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSSL := os.Getenv("DB_SSLMODE")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		dbHost, dbUser, dbPassword, dbName, dbPort, dbSSL,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	jwtExpire, _ := strconv.Atoi(os.Getenv("JWT_EXPIRE_HOURS"))
	appPort := os.Getenv("APP_PORT")

	return &Config{
		DB:        db,
		JWTSecret: jwtSecret,
		JWTExpire: jwtExpire,
		AppPort:   appPort,
	}
}
