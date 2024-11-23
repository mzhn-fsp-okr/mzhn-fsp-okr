package main

import (
	"context"
	"flag"

	"mzhn/notification-service/internal/app"

	"github.com/joho/godotenv"
)

var (
	local bool = false
)

func init() {
	flag.BoolVar(&local, "local", false, "run in local mode")
}

func main() {
	flag.Parse()

	if local {
		if err := godotenv.Load(); err != nil {
			panic("cannot load .env")
		}
	}

	ctx := context.Background()

	instance, cleanup, err := app.New()
	if err != nil {
		panic(err)
	}
	defer cleanup()

	instance.Run(ctx)
}
