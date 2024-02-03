package main

import (
	"flag"
	"log/slog"
	"os"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *slog.Logger
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	// TODO: Configure logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	app := &application{
		config: cfg,
		logger: logger,
	}

	app.serve()
}
