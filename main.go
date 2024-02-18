package main

import (
	"context"
	"fmt"

	"github.com/mihirkurdekar/lets-go/application"
)

func main() {
	app := application.New()

	err := app.Start(context.TODO())

	if err != nil {
		fmt.Printf("failed to start app:", err)
	}
}
