package main

import (
	"fmt"
	"net"

	"github.com/elli56/fibo-api/pkg/grpcserver"
	"github.com/elli56/fibo-api/pkg/proto"
	"github.com/elli56/fibo-api/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	logrus.SetLevel(logrus.DebugLevel)

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	// if err := godotenv.Load(); err != nil {
	// 	logrus.Fatalf("error loading env variables: %s", err.Error())
	// }

	// GRPC server
	fmt.Printf("Start GRPC server ':%s'\n", viper.GetString("port.grpc"))
	// create
	server := grpc.NewServer()

	// initialise Services
	services := service.NewService()
	// Объявим структуру которая реализует интерфейс нашего сервера
	GRPCServers := grpcserver.NewGRPCServer(services)
	//  registration
	proto.RegisterFibonacciSlicerServer(server, GRPCServers)

	l, err := net.Listen("tcp", fmt.Sprintf(":%s", viper.GetString("port.grpc")))
	if err != nil {
		logrus.Fatalf("Listen error occured: %s", err)
	}

	if err = server.Serve(l); err != nil {
		logrus.Fatalf("Serve error occured: %s", err)
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
