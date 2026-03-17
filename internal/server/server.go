package server

import (
	"context"
	"log/slog"
	"net"
	"net/http"
	"time"

	"github.com/DrxwDev/rest-api/internal/config"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func NewHTTPServer(lc fx.Lifecycle, router *gin.Engine, app config.AppConfig, srv config.ServerConfig, logger *slog.Logger) *http.Server {
	server := &http.Server{
		Addr:              net.JoinHostPort(app.HOST, app.PORT),
		Handler:           router,
		ReadTimeout:       srv.ReadTimeout,
		ReadHeaderTimeout: srv.ReadHeaderTimeout,
		WriteTimeout:      srv.WriteTimeout,
		IdleTimeout:       srv.IdleTimeout,
		MaxHeaderBytes:    srv.MaxHeaderBytes,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info("Server running", "addr", server.Addr)
			go func() {
				if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					logger.Error("Failed to start server", "err", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Shutting down server")
			shutDownCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
			defer cancel()

			if err := server.Shutdown(shutDownCtx); err != nil {
				logger.Error("Server shutdown failed", "err", err)
				return err
			}

			logger.Info("Server shutdown complete")
			return nil
		},
	})

	return server
}
