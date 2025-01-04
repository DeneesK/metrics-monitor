package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/DeneesK/metrics-monitor/internal/app/config"
	"github.com/DeneesK/metrics-monitor/internal/app/metcollector"
	"github.com/DeneesK/metrics-monitor/internal/app/router"
	"github.com/DeneesK/metrics-monitor/internal/app/server"
	"github.com/DeneesK/metrics-monitor/internal/pkg/logger"
)

func main() {
	conf := config.MustLoad()

	log := logger.NewLogger(conf.LogEnv)

	router := router.NewRouter(log)

	srv := server.NewServer(conf.Address, conf.Timeout, router)

	log.Info(fmt.Sprintf("starting server, listen %s", conf.Address))

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	metcollector.StartMetricCollector(ctx, conf.ScrapingInterval, log)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error("failed to start server", logger.Err(err))
			os.Exit(1)
		}
	}()

	<-ctx.Done()
	log.Info("shutdown initiated...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), conf.GracefulShutdownTimeout)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Error("Error during shutdown:", logger.Err(err))
		os.Exit(1)
	}
	<-shutdownCtx.Done()
	log.Info("server gracefully stopped")
}
