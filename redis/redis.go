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
		"model",
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

func (r *RedisClient) FetchMessage(key string) (*model.MessageData, error) {

	reply, err := redis.Values(r.client.Do("HGETALL", key))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var model model.MessageData
	err = redis.ScanStruct(reply, &model)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(model)
	return &model, err

}
