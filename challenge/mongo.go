package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectMongo cria e retorna uma nova conexão ao MongoDB.
func ConnectMongo(uri string) *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("Erro ao criar cliente MongoDB:", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("Erro ao conectar com MongoDB:", err)
	}
	return client
}

// DisconnectMongo encerra a conexão com o MongoDB.
func DisconnectMongo(client *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := client.Disconnect(ctx); err != nil {
		log.Fatal("Erro ao desconectar do MongoDB:", err)
	}
}
