package main

import (
	"context"
	"fmt"

	"github.com/ServiceWeaver/weaver"
)

func main() {
	if err := weaver.Run(context.Background(), run); err != nil {
		panic(err)
	}
}

type app struct {
	weaver.Implements[weaver.Main]
}

func run(ctx context.Context, a *app) error {
	fmt.Println("Hello, World!")
	return nil
}