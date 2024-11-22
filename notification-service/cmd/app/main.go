package main

import (
	"context"
	"flag"
	"fmt"

	"mzhn/notification-service/internal/app"

	"github.com/joho/godotenv"
)

func main() {
	flag.Parse()

	if err := godotenv.Load(); err != nil {
		panic(fmt.Errorf("cannot load env: %w", err))
	}

	ctx := context.Background()

	instance, cleanup, err := app.New()
	if err != nil {
		panic(err)
	}
	defer cleanup()

	instance.Run(ctx)
}
