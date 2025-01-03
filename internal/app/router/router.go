package router

import (
	"log/slog"

	"github.com/go-chi/chi/v5"
)

func NewRouter(log *slog.Logger) *chi.Mux {
	mux := chi.NewMux()

	return mux
}
