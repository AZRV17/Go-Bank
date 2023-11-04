package redis

import (
	"context"
	"github.com/AZRV17/goWEB/internal/config"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

var Rdb *redis.Client

func Connect(cfg *config.Config) error {
	Rdb = redis.NewClient(&redis.Options{
		Addr:         cfg.Redis.Host + ":" + cfg.Redis.Port,
		Password:     cfg.Redis.Password,
		DB:           cfg.Redis.Db,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
		MinIdleConns: 5,
	})

	if err := Rdb.Ping(context.Background()).Err(); err != nil {
		return err
	}

	return nil
}

func Close() error {
	if err := Rdb.Close(); err != nil {
		return err
	}

	log.Println("Redis connection closed")

	return nil
}
