package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"testing"
	"time"
)

func TestSubscribe(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	ctx := context.Background()

	subscriber := rdb.Subscribe(ctx, "my-channel")

	go func() {
		for {
			msg, err := subscriber.ReceiveMessage(ctx)
			if err != nil {
				panic(err)
			}
			fmt.Printf("%v\n", msg)
		}
	}()

	ticker := time.NewTicker(time.Second)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-ticker.C:
				{
					rdb.Publish(ctx, "my-channel", time.Now())
				}
			case <-done:
				return
			}
		}
	}()

	time.Sleep(10 * time.Minute)
}
