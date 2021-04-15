package app

import (
	"battleship/internal/config"
	handler "battleship/internal/delivery/http/v1"
	"battleship/internal/server"
	"battleship/internal/service"
	"battleship/pkg/logger"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func Run(configPath string) {
	cfg, err := config.Init(configPath)
	if err != nil {
		logrus.Error(err)
		return
	}

	//repos := repository.NewRepositories()
	services := service.NewService()
	handlers := handler.NewHandler(services)

	//handlers.Init(cfg.HTTP.Host, cfg.HTTP.Port)
	srv := server.NewServer(cfg, handlers.Init())

	go func() {
		if err := srv.Run(); err != nil {
			logrus.Errorf("error occurred while running http server: %s\n", err.Error())
		}
	}()

	logger.Info("Server started")

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit
}
