package rabbitmq

import (
	"context"
	"encoding/json"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQAdapter struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func NewRabbitMQServer(serverurl string) *RabbitMQAdapter {
	conn, err := amqp.Dial(serverurl)
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	//Declare Welcome Queue
	err = ch.ExchangeDeclare(
		"welcome_queue", // name
		"fanout",        // type
		true,            // durable
		false,           // auto-deleted
		false,           // internal
		false,           // no-wait
		nil,             // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	//Declare Ticket Order Queue
	err = ch.ExchangeDeclare(
		"ticket_queue", // name
		"fanout",       // type
		true,           // durable
		false,          // auto-deleted
		false,          // internal
		false,          // no-wait
		nil,            // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	return &RabbitMQAdapter{
		conn: conn,
		ch:   ch,
	}
}

func (q *RabbitMQAdapter) SendWelcomeEmail(email, fullName string) {

	// Create a map with message and phone number
	body := map[string]interface{}{
		"email":     email,
		"full_name": fullName,
	}

	// Encode the body map as JSON
	bodyJSON, err := json.Marshal(body)
	if err != nil {
		log.Fatalf("Failed to encode body map as JSON: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = q.ch.PublishWithContext(ctx,
		"welcome_queue", // exchange
		"",              // routing key
		false,           // mandatory
		false,           // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(bodyJSON),
		})
	failOnError(err, "Failed to publish a message")

	log.Println(" [x][x][x][x][x] Sent Welcome Message to RabbitMQ  [x][x][x][x][x] ")
}

func (q *RabbitMQAdapter) SendOrderConfirmation(email, fullName, ticketPDFUrl string) {

	// Create a map with message and phone number
	body := map[string]interface{}{
		"email":          email,
		"full_name":      fullName,
		"ticket_pdf_url": ticketPDFUrl,
	}

	// Encode the body map as JSON
	bodyJSON, err := json.Marshal(body)
	if err != nil {
		log.Fatalf("Failed to encode body map as JSON: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = q.ch.PublishWithContext(ctx,
		"ticket_queue", // exchange
		"",             // routing key
		false,          // mandatory
		false,          // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(bodyJSON),
		})
	failOnError(err, "Failed to publish a message")

	log.Println(" [x][x][x][x][x] Sent OrderConfirmation Message to RabbitMQ  [x][x][x][x][x] ")
}
