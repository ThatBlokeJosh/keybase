package utils

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)


var client = redis.NewClient(&redis.Options{
    Addr: "keydb:6379",
    Password: "",
    DB: 0,
})

var ctx = context.Background()

func Register(user User) {
    err := client.HSet(ctx, user.Username, "username", user.Username, "password", user.Password).Err();
    Error(err, fmt.Sprintf("Registering user %s failed", user.Username))
}

func Remove(user User) {
    err := client.HDel(ctx, user.Username).Err();
    Error(err, fmt.Sprintf("Removing user %s failed", user.Username))
}

func Set(key Key) {
    err := client.HSet(ctx, key.Username, key.Name, key.Key).Err();
    Error(err, fmt.Sprintf("Setting key %s failed", key.Name))
}

func Get(key Key) Key {
    val, err := client.HMGet(ctx, key.Username, key.Name).Result();
    Error(err, fmt.Sprintf("Getting key %s failed", key.Name))
    return Key{Username: key.Username, Name: key.Name, Key: val[0].(string)}
}

func Delete(key Key){
    err := client.HDel(ctx, key.Username, key.Name).Err();
    Error(err, fmt.Sprintf("Deleting key %s failed", key.Name))
}

