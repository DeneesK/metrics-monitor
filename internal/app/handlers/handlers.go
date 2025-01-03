package handlers

import (
	"log/slog"
	"net/http"
)

func NewMetricsHandler(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
