package main

import (
	"github.com/farukbey09/task/handler"
	"github.com/farukbey09/task/rabbitmq"
	"github.com/farukbey09/task/redis"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	redisClient := redis.NewRedisClient("localhost:6379")
	rabbitmqClient := rabbitmq.NewRabbitmqClient("amqp://user:password@localhost:7001")
	handler := handler.NewHandler(redisClient, rabbitmqClient)

	r.POST("/message", handler.HandleAddMessage)
	r.GET("/message/list", handler.HandleGetMessages)

	r.Run()
}
