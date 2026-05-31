package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/Ozenkol/rbk-go-final/internal/container"
	http_delivery "github.com/Ozenkol/rbk-go-final/internal/delivery/http"
	http_types "github.com/Ozenkol/rbk-go-final/internal/delivery/http/types"
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

	deps := &http_types.Dependencies{
		App: c.App(),
	}

	ctx := context.Background()
	cfg := http_delivery.DefaultConfig()

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
