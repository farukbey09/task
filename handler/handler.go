package handler

import (
	"github.com/farukbey09/task/model"
	"github.com/farukbey09/task/rabbitmq"
	"github.com/farukbey09/task/redis"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	redisClient    *redis.RedisClient
	rabbitmqClient *rabbitmq.RabbitmqClient
}

func NewHandler(redisClient *redis.RedisClient, rabbitmqClient *rabbitmq.RabbitmqClient) *Handler {
	return &Handler{
		redisClient:    redisClient,
		rabbitmqClient: rabbitmqClient,
	}
}

func (h *Handler) HandleAddMessage(c *gin.Context) {
	dto := model.DTO{}
	if err := c.BindJSON(&dto); err != nil {
		c.Status(400)
		return
	}

	model := model.MessageData{
		Sender:   dto.Sender,
		Receiver: dto.Receiver,
		Message:  dto.Message,
	}

	err := h.rabbitmqClient.AddQueue(model)
	if err != nil {
		c.Status(400)
		return
	}

	err = h.redisClient.AddMessageToRedis(model)
	if err != nil {
		c.Status(400)
		return
	}

	h.redisClient.FetchMessage("model")

}
