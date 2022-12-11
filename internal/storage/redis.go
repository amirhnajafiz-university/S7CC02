package storage

import (
	"context"
	"time"

	"github.com/go-redis/redis/v9"
)

// Storage manages the database connections.
type Storage struct {
	timeout int
	rdb     *redis.Client
}

// New builds a new storage.
func New(cfg Config) *Storage {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Host,
		Password: cfg.Password, // no password set
		DB:       0,            // use default DB
	})

	return &Storage{
		timeout: cfg.Timeout,
		rdb:     rdb,
	}
}

// Read key from database.
func (s *Storage) Read(key string) (string, error) {
	// create context
	ctx := context.Background()

	return s.rdb.Get(ctx, key).Result()
}

// Write a set into database.
func (s *Storage) Write(key string, value string) error {
	// create context
	ctx := context.Background()

	return s.rdb.Set(ctx, key, value, time.Duration(s.timeout)*time.Minute).Err()
}
