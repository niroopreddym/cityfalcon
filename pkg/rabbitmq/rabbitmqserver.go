package rabbitmq

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/niroopreddym/cityfalcon/pkg/database"
	"github.com/niroopreddym/cityfalcon/pkg/models"
	"github.com/niroopreddym/cityfalcon/pkg/services"
	"github.com/streadway/amqp"
)

type RabbitEvents struct {
	Connection      *amqp.Connection
	Channel         *amqp.Channel
	Queue           *amqp.Queue
	DatabaseService services.ISQLService
	RedisService    services.IRedisService
	errorChan       chan error
}

func NewConnection() (IRMQService, error) {
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
		Connection:      connection,
		Channel:         ch,
		Queue:           queue,
		DatabaseService: services.NewDatabaseServicesInstance(),
		RedisService:    services.NewRedisService(),
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
	r.errorChan = errorChan
	// declaring consumer with its properties over channel opened
	msgs, err := r.Channel.Consume(
		r.Queue.Name, // queue
		"",           // consumer
		false,        // auto ack
		false,        // exclusive
		false,        // no local
		false,        // no wait
		nil,          // args
	)

	if err != nil {
		errorChan <- err
	}

	// print consumed messages from queue
	go func() {
		for msg := range msgs {
			getAccDetailsModal := models.GetAccountDetails{}
			err := json.Unmarshal(msg.Body, &getAccDetailsModal)
			if err != nil {
				r.ackRMQMessage(msg.DeliveryTag)
				errorChan <- err
			}

			accDetails, err := r.DatabaseService.GetAccountDetails(getAccDetailsModal.XCorrelationID)
			time.Sleep(5 * time.Second)

			if err != nil {
				if err == database.NoRowError {
					r.RedisService.AddKey(getAccDetailsModal.XCorrelationID, fmt.Sprintf("bank with bank_uuid %v is not found", getAccDetailsModal.XCorrelationID))
					r.ackRMQMessage(msg.DeliveryTag)
					return
				}

				r.ackRMQMessage(msg.DeliveryTag)
				errorChan <- err
				return
			}

			err = r.RedisService.AddKey(getAccDetailsModal.XCorrelationID, fmt.Sprintf("%v", *accDetails.Balance))
			if err != nil {
				r.ackRMQMessage(msg.DeliveryTag)
				errorChan <- errors.New("Error occured while fetching the bank details")
			}

			//positive ack
			fmt.Println("message consumed: ", string(msg.Body))
			r.ackRMQMessage(msg.DeliveryTag)
		}
	}()

	fmt.Println("Waiting for messages...")

	<-stopChan
}

func (r *RabbitEvents) ackRMQMessage(deliveryTag uint64) {
	err := r.Channel.Ack(deliveryTag, false)
	if err != nil {
		r.errorChan <- errors.New(fmt.Sprintf("Error occured while ack-ing delivery tag: %v", err.Error()))
	}
}
