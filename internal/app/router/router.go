package router

import (
	"log/slog"

	"github.com/DeneesK/metrics-monitor/internal/app/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewRouter(log *slog.Logger) *chi.Mux {
	mux := chi.NewMux()
	mux.Handle("/metrics", promhttp.Handler())
	mux.Get("/ping", handlers.PingHandler)
	return mux
}
