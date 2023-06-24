package rabbitmq

import (
	"fmt"

	"github.com/streadway/amqp"
)

type RabbitEvents struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
	Queue      *amqp.Queue
}

func NewConnection() (*RabbitEvents, error) {
	fmt.Println("RabbitMQ in Golang: Getting started tutorial")

	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected to RabbitMQ instance")

	ch, err := createChannel(connection)
	if err != nil {
		return nil, err
	}

	queue, err := createQueue(ch)
	if err != nil {
		return nil, err
	}

	return &RabbitEvents{
		Connection: connection,
		Channel:    ch,
		Queue:      queue,
	}, nil
}

func createChannel(conn *amqp.Connection) (*amqp.Channel, error) {
	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return channel, nil
}

func createQueue(channel *amqp.Channel) (*amqp.Queue, error) {
	queue, err := channel.QueueDeclare(
		"testing", // name
		false,     // durable
		false,     // auto delete
		false,     // exclusive
		false,     // no wait
		nil,       // args
	)
	if err != nil {
		return nil, err
	}

	return &queue, nil
}

func (r *RabbitEvents) PublishMessage(messageBody []byte) error {
	err := r.Channel.Publish(
		"",           // exchange
		r.Queue.Name, // key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        messageBody,
		},
	)

	if err != nil {
		return err
	}

	fmt.Println("Queue status:", r.Queue)
	fmt.Println("Successfully published message")
	return nil
}

func (r *RabbitEvents) ConsumeMessage(stopChan chan bool, errorChan chan error) {
	// declaring consumer with its properties over channel opened
	msgs, err := r.Channel.Consume(
		r.Queue.Name, // queue
		"",           // consumer
		true,         // auto ack
		false,        // exclusive
		false,        // no local
		false,        // no wait
		nil,          //args
	)

	if err != nil {
		errorChan <- err
	}

	// print consumed messages from queue
	go func() {
		for msg := range msgs {
			fmt.Printf("Received Message: %s\n", msg.Body)
		}
	}()

	fmt.Println("Waiting for messages...")

	<-stopChan
}
