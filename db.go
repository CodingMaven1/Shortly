package main

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

// Client ...
var Client *redis.Client

// ConnectDatabase function to connect to redis ...
func ConnectDatabase() {
	Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	ctx := context.Background()
	_, err := Client.Ping(ctx).Result()
	if err != nil {
		log.Fatal(err)
	}
}
