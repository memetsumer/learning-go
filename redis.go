package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

func RunAndSetRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	ctx := context.Background()
	begin := time.Now()
	err := client.Set(ctx, "hello", "go", 10*time.Second).Err()
	end := time.Now().Sub(begin)
	fmt.Println(end)
	if err != nil {
		log.Fatal(err)
	}

}
