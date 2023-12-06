package mongo

import (
	"LastJake/config"
	"LastJake/internal/logger"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

var client *mongo.Client

func Connect() {
	logger.L.Info("Connecting MongoDB...")
	clientOptions := options.Client().ApplyURI(config.C.MongoURL)
	clientOptions.SetConnectTimeout(15 * time.Second)
	clientOptions.SetMaxConnIdleTime(10 * time.Second)
	clientOptions.SetMaxPoolSize(300)

	ctx, cancel := context.WithTimeout(context.Background(), ContextTimeout)
	defer cancel()

	var err error
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		logger.L.With("MongoDB").Fatal(err)
		return
	}

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		logger.L.With(err).Fatal("Failed to connect MongoDB")
	}

	logger.L.Info("Connected to MongoDB!")
	initIndexes()
}

func Client() *mongo.Client {
	if client == nil {
		Connect()
	}
	return client
}

func Collection(name collection) *mongo.Collection {
	return Client().Database(DBName).Collection(name.String())
}
