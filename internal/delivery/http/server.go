package http_delivery

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	_ "github.com/Ozenkol/rbk-go-final/api" // runtime doc registration, NOT in route.go

	http_types "github.com/Ozenkol/rbk-go-final/internal/delivery/http/types"
	"github.com/gin-gonic/gin"
)

type Mode string
const (
	Development Mode = "dev"
	Production Mode = "prod"
)

type ServerConfig struct {
	Host string
	Port int
	Mode Mode
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	IdleTimeout     time.Duration
	ShutdownTimeout time.Duration
}

type Server struct {
	httpServer *http.Server
	engine 		*gin.Engine
	cfg 	  ServerConfig	
	logger   *slog.Logger
}

func DefaultConfig() ServerConfig {
	return ServerConfig{
		Host: "localhost",
		Port: 8081,
		Mode: Development,
		ReadTimeout:     30 * time.Second,
		WriteTimeout:    30 * time.Second,
		IdleTimeout:     60 * time.Second,
		ShutdownTimeout: 30 * time.Second,
	}
}

func NewServer(cfg ServerConfig, logger *slog.Logger, deps *http_types.Dependencies) *Server {
	if logger == nil {
		logger = slog.Default()
	}

	engine := gin.New()

	registerRoutes(engine, logger, *deps)

	s := &Server{
		engine: engine,
		logger: logger,
		cfg:    cfg,
	}

	s.httpServer = &http.Server{
		Addr:         cfg.Host + ":" + strconv.Itoa(cfg.Port),
		Handler:      engine,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	}
	return s
}

func (s *Server) Start(ctx context.Context) error {
	serverErr := make(chan error, 1)
	go func() {
		s.logger.Info("Starting server", "host", s.cfg.Host, "port", s.cfg.Port)
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			serverErr <- err
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit,syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-serverErr:
		return fmt.Errorf("server error: %w", err)
	case sig := <-quit:
		s.logger.Info("Received shutdown signal", "signal", sig)
	case <-ctx.Done():
		s.logger.Info("Context cancelled, shutting down server")
	}

	return s.Shutdown()
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.cfg.ShutdownTimeout)
	defer cancel()

	s.logger.Info("Shutting down server")
	if err := s.httpServer.Shutdown(ctx); err != nil {
		return fmt.Errorf("server shutdown failed: %w", err)
	}

	s.logger.Info("Server gracefully stopped")
	return nil
}