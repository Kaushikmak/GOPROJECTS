package db

import (
	"context"
	"crypto/tls"
	"log"
	"os"
	"strings"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()
var Client *redis.Client

func InitRedis() {
	dbURL := os.Getenv("DB_URL")

	// Fallback for local docker-compose without Upstash
	if dbURL == "" {
		dbURL = "redis://" + os.Getenv("DB_ADDRESS")
	}

	opt, err := redis.ParseURL(dbURL)
	if err != nil {
		log.Fatalf("Failed to parse DB_URL: %v", err)
	}

	// Enforce TLS if using Upstash (rediss://)
	if strings.HasPrefix(dbURL, "rediss://") {
		opt.TLSConfig = &tls.Config{
			MinVersion: tls.VersionTLS12,
		}
	}

	Client = redis.NewClient(opt)

	if err := Client.Ping(Ctx).Err(); err != nil {
		log.Fatalf("Redis connection failed: %v", err)
	}
}
