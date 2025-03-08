package infraestructure

import (
	"bus-driver/src/turnos/domain"
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

type RabbitRepository struct {
    conn *amqp.Connection
}

func NewRabbitRepository(conn *amqp.Connection) *RabbitRepository {
	return &RabbitRepository{conn: conn}
}

func (r *RabbitRepository) PublishMessage(turno domain.Turno) error {
    ch, err := r.conn.Channel()
    if err != nil {
        return err
    }
    defer ch.Close()

    q, err := ch.QueueDeclare(
        "turnos", // nombre de la cola
        true, // durable
        false, // delete when unused
        false, // exclusive
        false, // no-wait
        nil,   // arguments
    )
    if err != nil {
        return err
    }

    body, err := json.Marshal(turno)
    if err != nil {
        return err
    }

    err = ch.Publish(
        "",     // exchange
        q.Name, // routing key
        false,  // mandatory
        false,  // immediate
        amqp.Publishing{
            ContentType: "application/json",
            Body:        body,
        },
    )
    if err != nil {
        return err
    }

    log.Printf("Mensaje enviado a la cola %s: %s", q.Name, string(body))
    return nil
}