package main

import (
	"context"

	"github.com/FrancoRutiliano/orders-api/aplication"
)

func main() {
	app := aplication.New()

	err := app.Start(context.TODO())

	if err != nil {

	}

}
