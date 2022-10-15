package redis

import (
	"fmt"
	"log"

	"github.com/farukbey09/task/model"
	"github.com/gomodule/redigo/redis"
)

type RedisClient struct {
	client redis.Conn
}

func NewRedisClient(url string) *RedisClient {
	conn, err := redis.Dial("tcp", url)
	if err != nil {

	}

	if err != nil {
		log.Fatal(err)
	}

	return &RedisClient{client: conn}
}

func (r *RedisClient) AddMessageToRedis(message model.MessageData) error {

	_, err := r.client.Do(
		"HMSET",
		"model:0",
		"message",
		message.Message,
		"sender",
		message.Sender,
		"receiver",
		message.Receiver,
	)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (r *RedisClient) FetchLastMessage() {

	values, _ := redis.StringMap(r.client.Do("HGETALL", "model:0"))

	for k, v := range values {
		fmt.Println("Key:", k)
		fmt.Println("Value:", v)
	}

}
