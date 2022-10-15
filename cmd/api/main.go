package main

import (
	"fmt"

	"github.com/farukbey09/task/handler"
	"github.com/farukbey09/task/rabbitmq"
	"github.com/farukbey09/task/redis"
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
)

func main() {
	r := gin.Default()
	redisClient := redis.NewRedisClient("localhost:6379")
	rabbitmqClient := rabbitmq.NewRabbitmqClient("amqp://user:password@localhost:7001")
	handler := handler.NewHandler(redisClient, rabbitmqClient)
	conn, err := amqp.Dial("amqp://user:password@localhost:7001")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer conn.Close()

	fmt.Println("succesfully connected to rabbitmq instance")
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer ch.Close()

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, "worked")
	})

	r.POST("/message", handler.HandleAddMessage)

	r.Run()
}
