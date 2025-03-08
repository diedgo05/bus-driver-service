package core

import (
	"fmt"
	"log"
	"os"


	"github.com/streadway/amqp"
	"github.com/joho/godotenv"
)

var (
	instance *amqp.Connection
	
)


func Connect() (*amqp.Connection, error) {
	
		if err := godotenv.Load(); err != nil {
			log.Fatalf("Error al cargar el .env: %v", err)
		}

		user := os.Getenv("R_USER")
		password := os.Getenv("R_PASSWORD")
		host := os.Getenv("R_IP")
		port := os.Getenv("R_PORT")

		rabbitURL := fmt.Sprintf("amqp://%s:%s@%s:%s/", user, password, host, port)

		conn, err := amqp.Dial(rabbitURL)
		if err != nil {
			log.Fatalf("Error al conectar con RabbitMQ: %v", err)
		}

		instance = conn
		log.Println("Conexión a RabbitMQ exitosa")
	

	return instance, nil
}