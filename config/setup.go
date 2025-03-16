package configs

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func LoadDBConfig() DBConfig {
	config := DBConfig{
		Host:     GetEnvWithDefault("DB_HOST", "localhost"),
		Port:     GetEnvWithDefault("DB_PORT", "54320"),
		User:     GetEnvWithDefault("DB_USER", "postgres"),
		Password: GetEnvWithDefault("DB_PASSWORD", "postgres"),
		DBName:   GetEnvWithDefault("DB_NAME", "postgres"),
		SSLMode:  GetEnvWithDefault("DB_SSLMODE", "disable"),
	}

	log.Printf("Loaded DB Config: %+v", config)
	return config
}

func ConnectDB() *gorm.DB {
	config := LoadDBConfig()

	log.Printf("DB Config: Host=%s, Port=%s, User=%s, DBName=%s, SSLMode=%s",
		config.Host, config.Port, config.User, config.DBName, config.SSLMode)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.Host,
		config.User,
		config.Password,
		config.DBName,
		config.Port,
		config.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	return db
}

// Global DB instance
var DB *gorm.DB = ConnectDB()

// GetDB returns the database instance
func GetDB() *gorm.DB {
	return DB
}
