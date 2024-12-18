package main

import (
	"context"
	"log"

	"github.com/acronix0/Building-service-Go/internal/app"
)

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf(err.Error())
	}
	a.Run()
}
