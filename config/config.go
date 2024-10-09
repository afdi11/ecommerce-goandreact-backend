package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Function untuk load environment dari .env file
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

// Struct untuk menampung variabel environment
type Config struct {
	HostPort         string
	RabbitMQHost     string
	RabbitMQPort     string
	RabbitMQUser     string
	RabbitMQPass     string
	MinioEndpoint    string
	MinioAccessKey   string
	MinioSecretKey   string
	MailpitHost      string
	MailpitPort      string
	PostgresUser     string
	PostgresPassword string
	PostgresDB       string
	PostgresHost     string
	PostgresPort     string
	RedisHost        string
	RedisPort        string
}

// Function untuk membaca environment dan mengisi struct Config
func GetConfig() Config {
	return Config{
		HostPort:         os.Getenv("HOST_PORT"),
		RabbitMQHost:     os.Getenv("RABBITMQ_HOST"),
		RabbitMQPort:     os.Getenv("RABBITMQ_PORT"),
		RabbitMQUser:     os.Getenv("RABBITMQ_USER"),
		RabbitMQPass:     os.Getenv("RABBITMQ_PASS"),
		MinioEndpoint:    os.Getenv("MINIO_ENDPOINT"),
		MinioAccessKey:   os.Getenv("MINIO_ACCESS_KEY"),
		MinioSecretKey:   os.Getenv("MINIO_SECRET_KEY"),
		MailpitHost:      os.Getenv("MAILPIT_HOST"),
		MailpitPort:      os.Getenv("MAILPIT_PORT"),
		PostgresUser:     os.Getenv("POSTGRES_USER"),
		PostgresPassword: os.Getenv("POSTGRES_PASSWORD"),
		PostgresDB:       os.Getenv("POSTGRES_DB"),
		PostgresHost:     os.Getenv("POSTGRES_HOST"),
		PostgresPort:     os.Getenv("POSTGRES_PORT"),
		RedisHost:        os.Getenv("REDIS_HOST"),
		RedisPort:        os.Getenv("REDIS_PORT"),
	}
}
