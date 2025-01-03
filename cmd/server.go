package main

import (
	"fmt"
	"os"

	"github.com/DeneesK/metrics-monitor/internal/app/config"
	"github.com/DeneesK/metrics-monitor/internal/app/server"
	"github.com/DeneesK/metrics-monitor/internal/pkg/logger"
)

func main() {
	conf := config.MustLoad()
	log := logger.NewLogger(conf.LogEnv)

	srv := server.NewServer(conf.Address, conf.Timeout)
	log.Info(fmt.Sprintf("starting server, listen: %s", conf.Address))
	if err := srv.ListenAndServe(); err != nil {
		log.Error("failed to start server", logger.Err(err))
		os.Exit(1)
	}

}
