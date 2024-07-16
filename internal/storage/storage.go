package storage

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func NewClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:	  "localhost:6379",
		Password: "", // no password set
		DB:		  0,  // use default DB
	})

	return client
}

func Ping(client *redis.Client) (string, error) {
	return client.Ping(ctx).Result()
}

func AddUrl(client *redis.Client, longUrl string) (error) {
	// short := shortener.ShortenUrl(longUrl)
	return client.Set(ctx, longUrl, "test", 0).Err()
}

func GetUrl(client *redis.Client, shortUrl string) (string, error) {
	return client.Get(ctx, shortUrl).Result()
}

func GetAllKeys(client *redis.Client) ([]string, error) {
	return client.Keys(ctx, "*").Result()
}