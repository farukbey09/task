package rabbitmq

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/farukbey09/task/model"
	"github.com/streadway/amqp"
)

type RabbitmqClient struct {
	client  *amqp.Connection
	channel *amqp.Channel
}

func NewRabbitmqClient(url string) *RabbitmqClient {
	conn, err := amqp.Dial(url)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("succesfully connected to rabbitmq instance")
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return &RabbitmqClient{client: conn, channel: ch}
}

func (r *RabbitmqClient) AddQueue(message model.MessageData) error {
	_, err := r.channel.QueueDeclare(
		"Message",
		false,
		false,
		false,
		false,
		nil)
	if err != nil {
		fmt.Println(err)
		return err
	}

	model, err := json.Marshal(message)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = r.channel.Publish("", "Message", false, false, amqp.Publishing{
		ContentType: "text/plain", Body: model,
	})

	if err != nil {
		fmt.Println(err)
		return err
	}

	go r.Message()
	return nil
}

func (r *RabbitmqClient) Message() {
	msgs, _ := r.channel.Consume("Message", "", true, false, false, false, nil)

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
