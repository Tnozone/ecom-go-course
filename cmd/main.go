package main

import (
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
)

func main() {
	cfg := config{
		addr: ":8080",
		db: dbConfig{
			dsn: "host=localhost user=postgres password=postgres dbname=ecom sslmode=disabled",
		},
	}

	//Database
	conn, err := pgx.Connect(ctx, cfg.db.dsn)
	if err != nil {
		return err
	}
	defer conn.Close(ctx)

	api := application{
		config: cfg,
	}

	// Logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	if err := api.run(api.mount()); err != nil {
		slog.Error("server failed to start", "error", err)
		os.Exit(1)
	}
}
