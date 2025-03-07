package main

import (
	"context"

	"github.com/viebiz/cqrs/sample/internal/controller/auth"
	"github.com/viebiz/lit"
)

func main() {
	// Start Server
	srv := lit.NewHttpServer(
		":8080",
		lit.Handler(context.Background(),
			lit.NewCORSConfig([]string{"*"}),
			func(router lit.Router) {
				initController(router)
			},
		),
	)
	srv.Run()
}

func initController(router lit.Router) {
	auth.New(router)
}
