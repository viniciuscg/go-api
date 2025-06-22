package main

import "log"

func main() {
	app := &application{
		config: config{
			addr: ":8080",
		},
	}

	log.Println("Starting server on", app.config.addr)

	log.Fatal(app.run())
}
