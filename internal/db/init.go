package db

import (
	"LastJake/internal/db/mongo"
	"LastJake/internal/db/redis"
)

func Init() {
	mongo.Connect()
	redis.Connect()
}
