package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/niroopreddym/cityfalcon/pkg/rabbitmq"
)

func main() {
	// Capture Ctrl-C
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	defer func() {
		signal.Stop(c)
	}()

	rabbitConnection, err := rabbitmq.NewConnection()
	if err != nil {
		fmt.Println(err)
	}
	rabbitConnection.PublishMessage([]byte("Hello World"))

	stopChan := make(chan bool)
	errorChan := make(chan error)

	go func(stopChan chan bool, errorChan chan error) {
		rabbitConnection.ConsumeMessage(stopChan, errorChan)
	}(stopChan, errorChan)

	select {
	case err := <-errorChan:
		fmt.Println(err)
		return
	case <-c:
		fmt.Println("cancel operation")
		stopChan <- true
		return
	}
}
