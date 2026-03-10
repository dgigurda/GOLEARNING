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

	paymentapi "golearning/payment/internal/api/payment/v1"
	paymentrepository "golearning/payment/internal/repository/payment"
	paymentservice "golearning/payment/internal/service/payment"
	paymentv1 "golearning/shared/pkg/proto/payment/v1"
)

const grpcPort = 50052

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

	repository := paymentrepository.NewRepository()
	service := paymentservice.NewService(repository)
	api := paymentapi.NewAPI(service)

	grpcServer := grpc.NewServer()
	paymentv1.RegisterPaymentServiceServer(grpcServer, api)
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
	grpcServer.GracefulStop()
}
