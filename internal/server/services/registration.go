package services

import (
	"github.com/fatih/color"
	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"main/internal/registration"
	"main/internal/structures"
	"path/filepath"
)

func initEnv() {
	filePrefix, _ := filepath.Abs("/var/www/investments-auto-registration-rambler/configs") // path from the working directory
	err := godotenv.Load(filePrefix + "/.env")
	if err != nil {
		color.New(color.FgRed).Add(color.Underline).Println(errors.Wrap(err, "ENV was not loaded correctly"))
	}
}

func Registration(i structures.AccInfo, rdb *redis.Client) {
	initEnv()

	registration.Registration(i, rdb)
}
