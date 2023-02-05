package main

import (
	"fmt"

	"github.com/eminoz/rabbitmq/broker"
)

func main() {
	// config.SetupConfig()
	// database.SetDatabase()
	fmt.Print("started")
	broker.RabbitConnection()
	broker.Consumers()

}
