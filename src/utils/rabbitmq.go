package utils

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

var conn *amqp.Connection

// QueueInfo struct
type QueueInfo struct {
	Name             string
	Durable          bool
	DeleteWhenUnused bool
	Exclusive        bool
	NoWait           bool
	Arguments        amqp.Table
}

// PublishInfo struct
type PublishInfo struct {
	Exchange  string
	Mandatory bool
	Immediate bool
}

// ConsumeInfo struct
type ConsumeInfo struct {
	Consumer  string
	AutoAck   bool
	exclusive bool
	NoLocal   bool
	NoWait    bool
	Arguments amqp.Table
}

// RabbitMQ interface
type RabbitMQ interface {
	Connect() (*amqp.Connection, error)
	HealthCheck(callback func(mess string))
	Publish(queueInfo QueueInfo, publishInfo PublishInfo, data string) error
	Consume(queueInfo QueueInfo, consumeInfo ConsumeInfo, callback func(data string) error) error
}

type rabbitMQ struct {
	url   string
	err   error
	delay int
}

// NewRabbitMQ func
func NewRabbitMQ(url string) RabbitMQ {
	return &rabbitMQ{url: url, delay: 0}
}

// Connect func
func (r *rabbitMQ) Connect() (*amqp.Connection, error) {
	log.Printf("RabbitMQ: connect to %s\n", r.url)
	c, err := amqp.Dial(r.url)
	if r.err = err; err != nil {
		return nil, err
	}
	conn, r.delay = c, 0
	log.Printf("RabbitMQ: connect to %s %s\n", r.url, "success.")
	return conn, err
}

// HealthCheck func
func (r *rabbitMQ) HealthCheck(callback func(mess string)) {
	var Err *amqp.Error
	for {
		connError := make(chan *amqp.Error)
		conn.NotifyClose(connError)
		Err = <-connError
		if Err != nil {
			r.RetryConnect(callback)
		}
	}
}

// RetryConnect func
func (r *rabbitMQ) RetryConnect(callback func(mess string)) {
	r.err = errors.New("RabbitMQ: re-connect to server")
	for {
		callback("RabbitMQ: re-connect to server...")
		time.Sleep(time.Duration(r.delay) * time.Second)
		r.delay += r.delay + 5
		_, err := r.Connect()
		if err == nil {
			callback("RabbitMQ: re-connect to server success.")
			break
		}
		if r.delay > 60 {
			callback("RabbitMQ: Lost connection.")
			callback(fmt.Sprintf("RabbitMQ: %s", err.Error()))
			return
		}
	}
}

// Warning: Need to close channel
func (r *rabbitMQ) makeChannelAndQueue(queueInfo QueueInfo) (*amqp.Channel, amqp.Queue, error) {
	if r.err != nil {
		return nil, amqp.Queue{}, r.err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, amqp.Queue{}, r.err
	}

	q, err := ch.QueueDeclare(
		queueInfo.Name,             // name
		queueInfo.Durable,          // durable
		queueInfo.DeleteWhenUnused, // delete when unused
		queueInfo.Exclusive,        // exclusive
		queueInfo.NoWait,           // no-wait
		queueInfo.Arguments,        // arguments
	)
	return ch, q, err
}

// Publish func
func (r *rabbitMQ) Publish(queueInfo QueueInfo, publishInfo PublishInfo, data string) error {
	ch, q, err := r.makeChannelAndQueue(queueInfo)
	defer ch.Close()
	if err != nil {
		return err
	}

	return ch.Publish(
		publishInfo.Exchange,  // exchange
		q.Name,                // routing key
		publishInfo.Mandatory, // mandatory
		publishInfo.Immediate, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(data),
		})
}

// Consume func
func (r *rabbitMQ) Consume(queueInfo QueueInfo, consumeInfo ConsumeInfo, callback func(data string) error) error {
	ch, q, err := r.makeChannelAndQueue(queueInfo)
	defer ch.Close()
	if err != nil {
		return err
	}
	msgs, err := ch.Consume(
		q.Name,                // queue
		consumeInfo.Consumer,  // consumer
		consumeInfo.AutoAck,   // auto-ack
		consumeInfo.exclusive, // exclusive
		consumeInfo.NoLocal,   // no-local
		consumeInfo.NoWait,    // no-wait
		consumeInfo.Arguments, // args
	)
	if err != nil {
		return err
	}
	r.releaseItem(msgs, consumeInfo.AutoAck, callback)
	return nil
}

func (r *rabbitMQ) releaseItem(msgs <-chan amqp.Delivery, autoAck bool, callback func(data string) error) {
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			err := callback(string(d.Body))
			if err != nil && !autoAck {
				d.Ack(true)
			}
		}
	}()
	<-forever
}
