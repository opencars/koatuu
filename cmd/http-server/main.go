package main

import (
	"context"
	"flag"
	"os/signal"
	"strconv"
	"syscall"

	_ "github.com/lib/pq"

	"github.com/opencars/bot/pkg/logger"
	"github.com/opencars/koatuu/pkg/api/http"
	"github.com/opencars/koatuu/pkg/config"
	"github.com/opencars/koatuu/pkg/domain/service"
	"github.com/opencars/koatuu/pkg/store/sqlstore"
	"github.com/opencars/schema/client"
)

func main() {
	cfg := flag.String("config", "config/config.yaml", "Path to the configuration file")
	port := flag.Int("port", 8080, "Port of the server")

	flag.Parse()

	conf, err := config.New(*cfg)
	if err != nil {
		logger.Fatalf("config: %v", err)
	}

	logger.NewLogger(logger.LogLevel(conf.Log.Level), conf.Log.Mode == "dev")

	c, err := client.New(conf.NATS.Address())
	if err != nil {
		logger.Fatalf("nats: %v", err)
	}

	producer, err := c.Producer()
	if err != nil {
		logger.Fatalf("producer: %v", err)
	}

	store, err := sqlstore.New(&conf.DB)
	if err != nil {
		logger.Fatalf("store: %v", err)
	}

	svc := service.NewCustomerService(store, producer)

	addr := ":" + strconv.Itoa(*port)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	logger.Infof("Listening on %s...", addr)
	if err := http.Start(ctx, addr, &conf.Server, svc); err != nil {
		logger.Fatalf("http server failed: %v", err)
	}
}
