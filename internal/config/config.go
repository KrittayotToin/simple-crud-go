package config

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

var MongoClient *mongo.Client

// ConnectMongoDB connects to the MongoDB instance
func ConnectMongoDB() error {
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		return fmt.Errorf("MONGO_URI is not set")
	}

	fmt.Println("Connecting to MongoDB...")

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongoURI).SetServerAPIOptions(serverAPI)

	// ใช้ context.TODO() เพื่อให้แน่ใจว่าใช้ context ที่ถูกต้อง
	client, err := mongo.Connect(opts)
	if err != nil {
		return err
	}

	MongoClient = client

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		return fmt.Errorf("MongoDB ping failed: %v", err)
	}

	fmt.Println("Successfully connected to MongoDB!")
	return nil
}

// DisconnectMongoDB disconnects from the MongoDB instance
func DisconnectMongoDB() error {
	if MongoClient != nil {
		return MongoClient.Disconnect(context.TODO())
	}
	return nil
}
