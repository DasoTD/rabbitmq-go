package controller

import (
	"net/http"
	"github.com/dasotd/rabbitmq-golang/constants"
	"github.com/dasotd/rabbitmq-golang/lib"

	"github.com/dasotd/rabbitmq-golang/utils"

	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
	// "github.com/streadway/amqp"
)

func Task1Controller(c *gin.Context) {
	ch := lib.RabbitChannel
	err := ch.PublishWithContext(
		c,
		"",              // exchange
		constants.QUEUE, // routing key
		false,           // mandatory
		false,           // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        nil,
			Type:        constants.TASK1,
		})

	utils.FailOnError(err, "Failed to publish a message")

	c.JSON(http.StatusOK, gin.H{
		"message": "Task-1 Received Successfully",
	})
}