package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	paymentV1 "golearning/shared/pkg/proto/payment/v1"
)

const grpcPort = 50052

type paymentService struct {
	paymentV1.UnimplementedPaymentServiceServer

	mu sync.RWMutex
}

func (s *paymentService) PayOrder(_ context.Context, in *paymentV1.PayOrderRequest) (*paymentV1.PayOrderResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	err := uuid.Validate(in.OrderUuid)
	if err != nil {
		return &paymentV1.PayOrderResponse{}, status.Error(codes.InvalidArgument, "Invalid order UUID")
	}
	err = uuid.Validate(in.UserUuid)
	if err != nil {
		return &paymentV1.PayOrderResponse{}, status.Error(codes.InvalidArgument, "Invalid user UUID")
	}
	if in.PaymentMethod == 0 {
		return &paymentV1.PayOrderResponse{}, status.Error(codes.InvalidArgument, "Invalid payment method")
	}

	newTransactionUUID := uuid.New().String()
	log.Printf("Оплата прошла успешно, transaction_uuid: %s\n", newTransactionUUID)
	return &paymentV1.PayOrderResponse{
		TransactionUuid: newTransactionUUID,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Printf("failed to listen: %v\n", err)
		return
	}
	defer func() {
		if cerr := lis.Close(); cerr != nil {
			log.Printf("failed to close listener: %v\n", cerr)
		}
	}()

	s := grpc.NewServer()

	service := &paymentService{}

	paymentV1.RegisterPaymentServiceServer(s, service)

	reflection.Register(s)

	go func() {
		log.Printf("🚀 gRPC server listening on %d\n", grpcPort)
		err = s.Serve(lis)
		if err != nil {
			log.Printf("failed to serve: %v\n", err)
			return
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("🛑 Shutting down gRPC server...")
	s.GracefulStop()
	log.Println("✅ Server stopped")
}
