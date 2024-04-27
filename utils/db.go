package utils

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)


var client = redis.NewClient(&redis.Options{
    Addr: "localhost:6379",
    Password: "",
    DB: 0,
})

var ctx = context.Background()

func Set(key Key) {
    err := client.Set(ctx, key.Name, key.Key, 0).Err()
    Error(err, fmt.Sprintf("Setting key %s failed", key.Name))
}

func Get(key Key) Key {
    val, err := client.Get(ctx, key.Name).Result()
    Error(err, fmt.Sprintf("Getting key %s failed", key.Name))
    return Key{Name: key.Name, Key: val}
}
