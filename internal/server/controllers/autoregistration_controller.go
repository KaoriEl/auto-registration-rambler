package controllers

import (
	"context"
	"fmt"
	"github.com/gammazero/workerpool"
	"github.com/go-redis/redis"
	"main/internal/server/services"
	"main/internal/structures"
)

var rdb = RedisCli()

func Index(v []structures.AccInfo) {
	wp := workerpool.New(10)
	for _, i := range v {
		r := i
		wp.Submit(func() {
			services.Registration(r, rdb)
		})
	}
	wp.StopWait()
}

func Verify(v []structures.AccInfo, c chan string) {
	wp := workerpool.New(1)
	for _, i := range v {
		r := i
		wp.Submit(func() {
			c <- services.Verify(r, rdb)
		})
	}
	wp.StopWait()
}

var ctx = context.Background()

func RedisCli() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, _ := rdb.Ping().Result()
	fmt.Println(pong + " - Redis Work Normally")

	return rdb

}
