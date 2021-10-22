package main

import (
	"fmt"

	"github.com/elli56/fibo-api/pkg/handler"
	"github.com/elli56/fibo-api/pkg/httpserver"
	"github.com/elli56/fibo-api/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	// HTTP server
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	fmt.Printf("Start HTTP server ':%s'\n", viper.GetString("port.http"))
	// if err := godotenv.Load(); err != nil {
	// 	logrus.Fatalf("error loading env variables: %s", err.Error())
	// }

	// repos := repository.NewRepository(db)
	// services := service.NewService(repos)
	services := service.NewService()
	handlers := handler.NewHandler(services)

	srv := new(httpserver.Server)
	if err := srv.Run(viper.GetString("port.http"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
