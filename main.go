package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/FrancoRutiliano/orders-api/aplication"
)

func main() {
	app := aplication.New()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := app.Start(ctx)

	if err != nil {
		fmt.Println("failed to connect to redis: %w", err)
	}

}
