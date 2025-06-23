package main

import (
	"github/viniciuscg/go-api/internal/env"
	"log"
)

func main() {
	app := &application{
		config: config{
			addr: env.GetEnvString("ADDR", ":4000"),
		},
	}

	log.Println("Starting server on", app.config.addr)
	mux := app.mount()
	log.Fatal(app.run(mux))
}
