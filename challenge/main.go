package main

import (
	"github.com/gin-gonic/gin"
	"project-root/internal/api"
	"project-root/internal/config"
	"project-root/internal/db"
	"project-root/internal/kafka"
)

func main() {
	// Inicializar Configurações
	cfg := config.LoadConfig()

	// Conectar ao MongoDB
	mongoClient := db.ConnectMongo(cfg)
	defer db.DisconnectMongo(mongoClient)

	// Inicializar Consumidor Kafka
	kafkaConsumer := kafka.InitConsumer(cfg)
	go kafka.ConsumeMessages(kafkaConsumer, mongoClient)

	// Configurar Gin
	router := gin.Default()
	api.SetupRoutes(router, mongoClient)

	// Iniciar servidor
	router.Run(cfg.ServerAddress)
}
