package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/iskorotkov/rusprofilegrpc/pkg"
)

// https://www.rusprofile.ru/ajax.php?query=5902879646&action=search

var (
	port = flag.Int("port", 8888, "port for gRPC server to listen on")
)

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("couldn't start gRPC server: %v", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	pkg.RegisterCompanyFinderServer(grpcServer, pkg.CompanyFinder{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("gRPC server stopped with error: %v", err)
	}
}
