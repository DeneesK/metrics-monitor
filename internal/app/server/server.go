package server

import (
	"net/http"
	"time"
)

func NewServer(address string, timeout time.Duration) *http.Server {
	srv := &http.Server{
		Addr:         address,
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
	}
	return srv
}
