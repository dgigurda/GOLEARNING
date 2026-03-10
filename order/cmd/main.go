package main

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	orderapi "golearning/order/internal/api/order/v1"
	clientgrpc "golearning/order/internal/client/grpc"
	partrepository "golearning/order/internal/repository/order"
	orderservice "golearning/order/internal/service/order"
	orderv1 "golearning/shared/pkg/openapi/order/v1"
)

const (
	httpPort      = "8080"
	inventoryPort = "50051"
	paymentPort   = "50052"

	readHeaderTimeout = 5 * time.Second
	shutdownTimeout   = 10 * time.Second
)

func main() {
	inventoryConn, err := grpc.NewClient(
		net.JoinHostPort("localhost", inventoryPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("inventory connection error: %v", err)
	}
	defer func() {
		if closeErr := inventoryConn.Close(); closeErr != nil {
			log.Printf("inventory connection close error: %v", closeErr)
		}
	}()

	paymentConn, err := grpc.NewClient(
		net.JoinHostPort("localhost", paymentPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("payment connection error: %v", err)
	}
	defer func() {
		if closeErr := paymentConn.Close(); closeErr != nil {
			log.Printf("payment connection close error: %v", closeErr)
		}
	}()

	repository := partrepository.NewRepository()
	inventoryClient := clientgrpc.NewInventoryClient(inventoryConn)
	paymentClient := clientgrpc.NewPaymentClient(paymentConn)
	service := orderservice.NewService(repository, inventoryClient, paymentClient)
	api := orderapi.NewAPI(service)

	server, err := orderv1.NewServer(api)
	if err != nil {
		log.Fatalf("openapi server creation error: %v", err)
	}

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(10 * time.Second))
	router.Mount("/", server)

	httpServer := &http.Server{
		Addr:              net.JoinHostPort("localhost", httpPort),
		Handler:           router,
		ReadHeaderTimeout: readHeaderTimeout,
	}

	go func() {
		log.Printf("HTTP server listening on %s", httpPort)
		if serveErr := httpServer.ListenAndServe(); serveErr != nil && !errors.Is(serveErr, http.ErrServerClosed) {
			log.Printf("http server error: %v", serveErr)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if shutdownErr := httpServer.Shutdown(ctx); shutdownErr != nil {
		log.Printf("http shutdown error: %v", shutdownErr)
	}
}
