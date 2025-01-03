package server

import (
	"net/http"
	"time"
)

func NewServer(address string, timeout time.Duration, router http.Handler) *http.Server {
	srv := &http.Server{
		Addr:         address,
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
		Handler:      router,
	}
	return srv
}
