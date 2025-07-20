package cache

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

type Service interface {
	Set(ctx context.Context, key string, value any, expiration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	BatchDelByPrefix(ctx context.Context, prefix string) error
}

type service struct {
	client *redis.Client
}

func New() Service {
	opt, _ := redis.ParseURL(os.Getenv("REDIS_URL"))
	client := redis.NewClient(opt)

	return &service{
		client: client,
	}
}

func (s *service) Set(ctx context.Context, key string, value any, expiration time.Duration) error {
	log.Println("Hello?")
	return s.client.Set(ctx, key, value, expiration).Err()
}

func (s *service) Get(ctx context.Context, key string) (string, error) {
	val, err := s.client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	return val, nil
}

func (s *service) BatchDelByPrefix(ctx context.Context, prefix string) error {
	var cursor uint64
	const batchSize = 100

	for {
		keys, nextCursor, err := s.client.Scan(ctx, cursor, prefix+"*", batchSize).Result()
		if err != nil {
			return fmt.Errorf("scan failed: %w", err)
		}

		if len(keys) > 0 {
			deleted := s.client.Del(ctx, keys...)
			if deleted.Err() != nil {
				return fmt.Errorf("delete failed: %w", deleted.Err())
			}
		}

		cursor = nextCursor
		if cursor == 0 {
			break
		}
	}
	return nil
}
