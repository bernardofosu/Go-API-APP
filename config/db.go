package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

func ConnectDB() {
	// Load environment variables if not already loaded
	if AppConfig.MongoURI == "" {
		LoadEnv()
	}

	clientOptions := options.Client().ApplyURI(AppConfig.MongoURI)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("❌ MongoDB connection error: %v", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("❌ Could not ping MongoDB: %v", err)
	}

	fmt.Println("✅ Connected to MongoDB Atlas")
	DB = client
}

func GetCollection(collectionName string) *mongo.Collection {
	return DB.Database(AppConfig.DBName).Collection(collectionName)
}
