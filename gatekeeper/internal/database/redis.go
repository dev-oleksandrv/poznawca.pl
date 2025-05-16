package database

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log/slog"
)

type RedisDatabase struct {
	*redis.Client
}

type NewRedisDatabaseConfig struct {
	Addr     string
	Password string
	DB       int
}

func NewRedisDatabase(config NewRedisDatabaseConfig) (*RedisDatabase, error) {
	redisOptions := &redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.DB,
	}

	client := redis.NewClient(redisOptions)

	if err := client.Ping(context.Background()).Err(); err != nil {
		slog.Error("failed to connect to Redis", "err", err)
		return nil, err
	}

	return &RedisDatabase{Client: client}, nil
}
