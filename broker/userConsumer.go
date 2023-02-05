package broker

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/eminoz/rabbitmq/model"
)

func UserConsumer() {
	ch, conn := GetRabbitChannel()
	q, err := ch.QueueDeclare(
		"createdUser", // name
		true,          // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		fmt.Print(err)
	}

	defer ch.Close()
	defer conn.Close()
	mailqueueConsume, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack you have to rerun code when this field true
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatalf("%s: %s", "Failed to register consumer", err)
	}

	forever := make(chan bool)
	go func() {
		for d := range mailqueueConsume {
			user := model.User{}

			json.Unmarshal(d.Body, &user)
			fmt.Println(user.Email, "'e mail yollandı ")
			d.Ack(false)
		}

	}()

	<-forever
}
