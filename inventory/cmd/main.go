package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	inventoryapi "golearning/inventory/internal/api/inventory/v1"
	partrepository "golearning/inventory/internal/repository/part"
	partservice "golearning/inventory/internal/service/part"
	inventoryv1 "golearning/shared/pkg/proto/inventory/v1"
)

const grpcPort = 50051

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer func() {
		if closeErr := listener.Close(); closeErr != nil {
			log.Printf("failed to close listener: %v", closeErr)
		}
	}()

	repository := partrepository.NewRepository()
	service := partservice.NewService(repository)
	api := inventoryapi.NewAPI(service)

	grpcServer := grpc.NewServer()
	inventoryv1.RegisterInventoryServiceServer(grpcServer, api)
	reflection.Register(grpcServer)

	go func() {
		log.Printf("gRPC server listening on %d", grpcPort)
		if serveErr := grpcServer.Serve(listener); serveErr != nil {
			log.Printf("failed to serve: %v", serveErr)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("shutting down gRPC server")
	grpcServer.GracefulStop()
	log.Println("server stopped")
}
