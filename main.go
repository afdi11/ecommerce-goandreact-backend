package main

import (
	"ecommerce/config"
	"ecommerce/routes"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Get configuration
	cfg := config.GetConfig()

	// Contoh penggunaan variabel environment
	fmt.Println("RabbitMQ Host:", cfg.RabbitMQHost)
	fmt.Println("Minio Endpoint:", cfg.MinioEndpoint)
	fmt.Println("Mailpit Host:", cfg.MailpitHost)

	// Implementasi RabbitMQ connection menggunakan config
	rabbitConnString := fmt.Sprintf("amqp://%s:%s@%s:%s/",
		cfg.RabbitMQUser, cfg.RabbitMQPass, cfg.RabbitMQHost, cfg.RabbitMQPort)

	fmt.Println("RabbitMQ Connection String:", rabbitConnString)

	fmt.Println("Connecting to PostgreSQL...")
	config.ConnectDatabase(cfg)
	fmt.Println("Successfully connected to PostgreSQL!")

	// Implementasi Redis connection menggunakan config
	fmt.Println("Connecting to Redis...")
	rdb := config.ConnectRedis(cfg)
	config.CacheExample(rdb)
	defer rdb.Close()

	// Jalankan server
	r := routes.SetupRouter()
	r.Run(cfg.HostPort)
	// Lakukan setup RabbitMQ, Minio, Mailpit, dll dengan menggunakan cfg sesuai kebutuhan
}
