package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"slices"
	"sync"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	invetoryV1 "golearning/shared/pkg/proto/inventory/v1"
)

const grpcPort = 50051

type inventoryService struct {
	invetoryV1.UnimplementedInventoryServiceServer

	Parts map[string]*invetoryV1.Part
	mu    sync.RWMutex
}

func (s *inventoryService) GetPart(_ context.Context, in *invetoryV1.GetPartRequest) (*invetoryV1.GetPartResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	part, ok := s.Parts[in.Uuid]
	if !ok {
		return &invetoryV1.GetPartResponse{}, status.Errorf(codes.NotFound, "In DB no parts /w that UUID")
	}

	return &invetoryV1.GetPartResponse{
		Part: part,
	}, nil
}

func (s *inventoryService) ListParts(_ context.Context, in *invetoryV1.ListPartsRequest) (*invetoryV1.ListPartsResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	parts := make([]*invetoryV1.Part, 0)
	filters := in.Filter
	for _, uuid := range filters.Uuids {

		part, ok := s.Parts[uuid]
		if !ok {
			continue
		}
		parts = append(parts, part)
	}
	for i, part := range parts {
		var ok bool
		for _, name := range filters.Names {
			if part.Name == name {
				ok = true
			}
		}
		if !ok {
			parts = slices.Delete(parts, i, i+1)
		}
	}
	for i, part := range parts {
		var ok bool
		for _, category := range filters.Categorys {
			if part.Category == category {
				ok = true
			}
		}
		if !ok {
			parts = slices.Delete(parts, i, i+1)
		}
	}
	for i, part := range parts {
		var ok bool
		for _, country := range filters.ManufacturerCountries {
			if part.Manufacturer.Country == country {
				ok = true
			}
		}
		if !ok {
			parts = slices.Delete(parts, i, i+1)
		}
	}
	for i, part := range parts {
		var ok bool
		for _, tag := range filters.Tags {
			if slices.Contains(part.Tags, tag) {
				ok = true
			}
		}
		if !ok {
			parts = slices.Delete(parts, i, i+1)
		}
	}
	return &invetoryV1.ListPartsResponse{
		Parts: parts,
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

	service := &inventoryService{}

	invetoryV1.RegisterInventoryServiceServer(s, service)

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
