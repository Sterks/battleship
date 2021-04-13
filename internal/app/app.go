package app

import (
	"battleship/internal/config"
	"battleship/internal/server"
	"battleship/pkg/logger"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func Run(configPath string){
	cfg, err := config.Init(configPath)
	if err != nil {
		logrus.Error(err)
		return
	}

	//repos := repository.NewRepositories()


	//handlers.Init(cfg.HTTP.Host, cfg.HTTP.Port)
	srv := server.NewServer(cfg, nil)
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
