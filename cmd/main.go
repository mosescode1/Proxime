package main

import (
	"flag"
	"log"
	"log/slog"

	"github.com/mosescode1/Proxime/config"
	"github.com/mosescode1/Proxime/internal/redis"
	"github.com/mosescode1/Proxime/internal/tile38"
	"github.com/mosescode1/Proxime/pkg/logger"
	"github.com/mosescode1/Proxime/server"
)

var configPath = flag.String("config", ".env", "path to config file")

func main() {
	flag.Parse()

	appConfig, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatal("Error loading config: ", err)
	}
	logger.SetUpLogger(appConfig.APP_ENV)

	client, err := redis.NewRedisClient()
	if err != nil {
		slog.Error("Error connecting to Redis: ", "error", err)
		log.Fatal("Error connecting to Redis: ", err)
	}

	defer redis.CloseRedis(client)

	t38, err := tile38.NewT38Client()
	if err != nil {
		log.Fatal("Error connecting to Tile38: ", err)
	}
	defer tile38.T38Close(t38)

	slog.Info("Starting Http Server", slog.String("port", appConfig.APP_PORT))
	server.StartServer(appConfig)
}
