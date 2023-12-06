package mongo

import (
	"LastJake/internal/logger"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func initIndexes() {
	indexUser()
}

func indexUser() {
	ctx, cancel := context.WithTimeout(context.Background(), ContextTimeout)
	defer cancel()

	idxCreatedAt := mongo.IndexModel{
		Keys:    bson.D{{"createdAt", -1}},
		Options: options.Index().SetName("created_desc"),
	}
	_, err := Collection(User).Indexes().CreateMany(ctx, []mongo.IndexModel{
		idxCreatedAt,
	})
	if err != nil {
		logger.L.With(err).Warn("mongo index error")
	}
}
