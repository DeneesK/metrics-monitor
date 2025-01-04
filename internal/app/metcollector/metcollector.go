package metcollector

import (
	"context"
	"log/slog"
	"time"

	"github.com/DeneesK/metrics-monitor/internal/pkg/logger"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

func collectMetrics(ctx context.Context, scrapeInterval time.Duration, log *slog.Logger) {
	var (
		cpuUsage = prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "server_cpu_usage_percentage",
			Help: "Текущая загрузка CPU в процентах",
		})
		memoryUsage = prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "server_memory_usage_percentage",
			Help: "Текущая загрузка памяти в процентах",
		})
		networkSent = prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "server_network_bytes_sent",
			Help: "Отправленные байты через сеть",
		})
		networkReceived = prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "server_network_bytes_received",
			Help: "Принятые байты через сеть",
		})
		diskSpace = prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "server_disk_usage_percentage",
			Help: "Текущая загрузка диска в процентах",
		})
	)

	prometheus.MustRegister(cpuUsage, memoryUsage, networkSent, networkReceived, diskSpace)

	for {
		cpuPercent, err := cpu.Percent(time.Second, false)
		if err != nil {
			log.Error("failed to collect cpuPercent", logger.Err(err))
		} else {
			if len(cpuPercent) > 0 {
				cpuUsage.Set(cpuPercent[0])
			}
		}

		vMem, err := mem.VirtualMemory()
		if err != nil {
			log.Error("failed to collect vMem", logger.Err(err))
		} else {
			memoryUsage.Set(vMem.UsedPercent)
		}

		dSpace, err := disk.Usage("/")
		if err != nil {
			log.Error("failed to collect diskSpace", logger.Err(err))
		} else {
			diskSpace.Set(dSpace.UsedPercent)
		}

		netStats, err := net.IOCounters(false)
		if err != nil {
			log.Error("failed to collect netStats", logger.Err(err))
		} else {
			if len(netStats) > 0 {
				networkSent.Set(float64(netStats[0].BytesSent))
				networkReceived.Set(float64(netStats[0].BytesRecv))
			}
		}
		select {
		case <-ctx.Done():
			log.Info("gracefully shutdown metric collector")
			return
		default:
			time.Sleep(scrapeInterval)
		}
	}
}

func StartMetricCollector(ctx context.Context, scrapeInterval time.Duration, log *slog.Logger) {
	log.Info("starting metric collector")
	go collectMetrics(ctx, scrapeInterval, log)
}
