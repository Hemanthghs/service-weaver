package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/ServiceWeaver/weaver"
)

type app struct {
	weaver.Implements[weaver.Main]
	reverser weaver.Ref[Greeter]
	greet    weaver.Listener `weaver:"greet"`
}

func serve(ctx context.Context, app *app) error {
	fmt.Printf("Running on port %v\n", app.greet)
	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		username := r.URL.Query().Get("username")
		greet, err := app.reverser.Get().Greet(ctx, username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		fmt.Fprint(w, " ", greet)
	})
	return http.Serve(app.greet, nil)
}

func main() {
	if err := weaver.Run(context.Background(), serve); err != nil {
		log.Fatal(err)
	}
}
