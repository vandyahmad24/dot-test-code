package repository

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisRepository interface {
	Set(ctx context.Context, key string, value interface{}, duration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Delete(ctx context.Context, key string) error
}

type redisRepository struct {
	client *redis.Client
}

func NewRedisRepository(addr string, password string, db int) *redisRepository {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	return &redisRepository{
		client: rdb,
	}
}

func (r *redisRepository) Set(ctx context.Context, key string, value interface{}, duration time.Duration) error {
	err := r.client.Set(ctx, key, value, 0).Err()
	if err != nil {
		log.Printf("Failed to set key %s: %v", key, err)
	}
	return err
}

// Get retrieves a value from Redis
func (r *redisRepository) Get(ctx context.Context, key string) (string, error) {
	log.Println("Getting value from Redis...")
	val, err := r.client.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			log.Printf("Key %s does not exist", key)
			return "", nil
		}
		log.Printf("Failed to get key %s: %v", key, err)
		return "", err
	}
	log.Println("Value retrieved from Redis:", val)
	return val, nil
}

func (r *redisRepository) Delete(ctx context.Context, key string) error {
	err := r.client.Del(ctx, key).Err()
	if err != nil {
		log.Printf("Failed to delete key %s: %v", key, err)
	}
	return err
}
