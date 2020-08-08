package queue

import (
	"fmt"
	"testing"

	"github.com/streadway/amqp"
)

const (
	URL       string = "amqp://guest:guest@localhost:5672/"
	QueueName string = "MIGRATE_DATA_COIN_SERVICE"
)

func TestRabbitMQ_Sending(t *testing.T) {
	conn, err := amqp.Dial(URL)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer conn.Close()

	rab := NewRabbitMQ(conn)

	for i := 0; i < 10000; i++ {
		err := rab.Sending(QueueName, fmt.Sprintf("Message: %d", i))
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func TestRabbitMQ_Receiving(t *testing.T) {
	conn, err := amqp.Dial(URL)
	if err != nil {
		return
	}
	defer conn.Close()

	rab := NewRabbitMQ(conn)
	rab.Receiving(QueueName, callback)
}

func callback(data string) error {
	fmt.Println(data)
	return nil
}

/*
conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
if err != nil {
	return err
}
defer conn.Close()
rab := queue.NewRabbitMQ(conn)
go rab.Receiving("MIGRATE_DATA_COIN_SERVICE", callback)
*/
