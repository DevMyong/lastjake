package config

import (
	"LastJake/internal/logger"
	"github.com/joho/godotenv"
	"os"
)

var C Config

func init() {
	logger.L.Info("Start to set config...")
	C.init()
}

type Config struct {
	MongoURL string
	RedisURL string

	Token     string
	secretKey string
}

func (c *Config) init() {
	err := c.loadEnv()
	if err != nil {
		return
	}
}

func (c *Config) loadEnv() error {
	err := godotenv.Load()
	if err != nil {
		logger.L.Fatal(err)
		return err
	}

	c.MongoURL = os.Getenv("MONGO_URL")
	c.RedisURL = os.Getenv("REDIS_URL")

	c.Token = os.Getenv("DISCORD_TOKEN")
	c.secretKey = os.Getenv("SECRET_KEY")

	return nil
}
