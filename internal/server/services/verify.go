package services

import (
	"github.com/go-redis/redis"
	"main/internal/structures"
	"main/internal/verify"
)

func Verify(i structures.AccInfo, rdb *redis.Client) string {
	initEnv()

	return verify.Verify(i, rdb)
}
