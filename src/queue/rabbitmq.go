package queue

import (
	"fmt"

	"github.com/streadway/amqp"
)

// RabbitMQ interface
type RabbitMQ interface {
	Sending(queueName string, data string) error
	Receiving(queueName string, callback func(data string) error) error
}

type rabbitMQ struct {
	conn *amqp.Connection
}

// NewRabbitMQ func
func NewRabbitMQ(conn *amqp.Connection) RabbitMQ {
	return &rabbitMQ{conn: conn}
}

// Sending func
func (r *rabbitMQ) Sending(queueName string, data string) error {
	ch, err := r.conn.Channel()
	defer ch.Close()
	if err != nil {
		return err
	}

	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return err
	}

	return ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(data),
		})
}

// Receiving func
func (r *rabbitMQ) Receiving(queueName string, callback func(data string) error) error {
	ch, err := r.conn.Channel()
	defer ch.Close()
	if err != nil {
		return err
	}

	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return err
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		true,   // no-wait
		nil,    // args
	)
	if err != nil {
		return err
	}

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			errCallback := callback(string(d.Body))
			if errCallback != nil {
				fmt.Println(errCallback.Error())
			}
			//time.Sleep(1 * time.Millisecond)
			//d.Ack(true)
		}
	}()
	<-forever

	return nil
}
