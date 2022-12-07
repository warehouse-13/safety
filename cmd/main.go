package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	mvmv1 "github.com/weaveworks-liquidmetal/flintlock/api/services/microvm/v1alpha1"
	"google.golang.org/grpc"

	"github.com/warehouse-13/safety"
)

func main() {
	var port string

	flag.StringVar(&port, "port", "9090", "port number to start server on")
	flag.Parse()

	l, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		fmt.Println("Failed to listen on localhost:"+port, err)
		os.Exit(1)
	}

	s := &safety.FakeServer{}
	token := os.Getenv("AUTH_TOKEN")
	grpcServer := grpc.NewServer(safety.WithOpts(token)...)
	mvmv1.RegisterMicroVMServer(grpcServer, s)

	fmt.Println("starting server on localhost:" + port)

	if err := grpcServer.Serve(l); err != nil {
		fmt.Println("Failed to start gRPC server", err)
		os.Exit(1)
	}
}
