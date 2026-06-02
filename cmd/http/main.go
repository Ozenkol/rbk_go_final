package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/Ozenkol/rbk-go-final/internal/container"
	http_delivery "github.com/Ozenkol/rbk-go-final/internal/delivery/http"
	http_deps "github.com/Ozenkol/rbk-go-final/internal/delivery/http/deps"
)

func main() {
	slog.SetDefault(
		slog.New(
			slog.NewTextHandler(os.Stdout, nil),
		),
	)

	c := container.New(container.Config{
		DSN: "host=localhost user=postgres password=password dbname=rbk_db port=5432 sslmode=disable", // pull from os.Getenv("DATABASE_DSN") in prod
	})

	ctx := context.Background()
	cfg := http_delivery.DefaultConfig()

	deps := http_deps.NewDependencies(c.App())

	server := http_delivery.NewServer(
		cfg,
		slog.Default(),
		deps,
	)

	if err := server.Start(ctx); err != nil {
		slog.Error("server stopped", "error", err)
		os.Exit(1)
	}
}
