package main

import (
	"log"
	"socialApp/internal/env"
)

func main() {
	cfg := config{
		addr: env.GetEnv("ADDR", ":8080"),
	}
	app := &application{
		config: cfg,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
