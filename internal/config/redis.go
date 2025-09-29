package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

type RedisClient struct {
	Client *redis.Client
	Prefix string
	ctx    context.Context
}

func NewRedisClient(viper *viper.Viper) *RedisClient {
	host := viper.GetString("REDIS_HOST")
	port := viper.GetString("REDIS_PORT")
	db := viper.GetInt("REDIS_DB")
	password := viper.GetString("REDIS_PASSWORD")
	prefix := viper.GetString("REDIS_PREFIX")

	addr := fmt.Sprintf("%s:%s", host, port)
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	ctx := context.Background()

	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatalf("[Redis] ❌ Gagal koneksi ke Redis di %s (DB %d): %v", addr, db, err)
	} else {
		log.Printf("[Redis] ✅ Koneksi berhasil ke Redis di %s (DB %d) dengan prefix '%s'", addr, db, prefix)
	}

	return &RedisClient{
		Client: rdb,
		Prefix: prefix,
		ctx:    context.Background(),
	}
}

func (r *RedisClient) key(k string) string {
	return fmt.Sprintf("%s:%s", r.Prefix, k)
}

func (r *RedisClient) Set(key string, value string, ttl int) error {
	return r.Client.Set(r.ctx, r.key(key), value, time.Duration(ttl)).Err()
}

func (r *RedisClient) Get(key string) (string, error) {
	return r.Client.Get(r.ctx, r.key(key)).Result()
}
