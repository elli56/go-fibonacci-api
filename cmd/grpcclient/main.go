package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/elli56/fibo-api/pkg/proto"
	"google.golang.org/grpc"
)

var (
	port = flag.String("p", "8080", "first argument X")
	x    = flag.Int64("x", 0, "first argument X")
	y    = flag.Int64("y", 0, "first argument Y")
)

func main() {
	flag.Parse()
	// if flag.NArg() < 2 {
	// 	logrus.Fatal("not enough arguments")
	// }

	// x, err := strconv.Atoi(flag.Arg(0))
	// if err != nil {
	// 	logrus.Fatal(err)
	// }

	// y, err := strconv.Atoi(flag.Arg(1))
	// if err != nil {
	// 	logrus.Fatal(err)
	// }

	connect, err := grpc.Dial(fmt.Sprintf(":%s", *port), grpc.WithInsecure())
	if err != nil {
		logrus.Fatal(err)
	}

	// create a client
	client := proto.NewFibonacciSlicerClient(connect)
	res, err := client.FibonacciSlice(context.Background(), &proto.FiboRequest{X: *x, Y: *y})
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Println(res.GetResult())

}
