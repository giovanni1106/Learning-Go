package api

import (
	"github.com/Shopify/sarama"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SetupRoutes configura as rotas do servidor Gin.
func SetupRoutes(router *gin.Engine, producer sarama.SyncProducer) {
	router.POST("/task", func(c *gin.Context) {
		var task Task
		if err := c.ShouldBindJSON(&task); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}
		// Serializar task como JSON e enviar como mensagem Kafka
		msg := &sarama.ProducerMessage{
			Topic: "tasks",
			Value: sarama.StringEncoder(task),
		}
		_, _, err := producer.SendMessage(msg)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send message"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "success"})
	})
}

// Task representa uma tarefa.
type Task struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}
