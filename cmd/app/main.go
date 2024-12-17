package main

import (
	"context"
	"log"
)

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf(err.Error())
	}
	a.Run()
}