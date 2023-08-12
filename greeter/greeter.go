package main

import (
	"context"
	"fmt"

	"github.com/ServiceWeaver/weaver"
)

type Greeter interface {
	Greet(context.Context, string) (string, error)
}

type greeter struct {
	weaver.Implements[Greeter]
}

func (r *greeter) Greet(ctx context.Context, username string) (string, error) {
	res := fmt.Sprintf("Welcome, %s...!. This is service weaver", username)
	return res, nil
}
