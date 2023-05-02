package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"link_shorter/internal/config"
	"link_shorter/internal/handler"
	"link_shorter/internal/pkg/service"
	"link_shorter/internal/pkg/storage"
	"link_shorter/internal/pkg/storage/inmemory"
	"link_shorter/internal/pkg/storage/postgres"
	"link_shorter/internal/protobuf/link_shorter/protobuf/shorter"

	"github.com/joho/godotenv"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	memory     = "inmemory"
	postgresql = "postgresql"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	cfg := config.InitConfig()

	storeType := os.Getenv("TYPE")
	var store storage.Store
	if storeType == memory {
		store = inmemory.NewInMemory()
	} else if storeType == postgresql {
		store = postgres.NewPostgres(cfg)
	} else {
		log.Fatalf("unknown storage type")
	}

	s := grpc.NewServer()
	linkService := service.NewLinkService(store)
	implementation := handler.NewShorter(linkService)
	shorter.RegisterShorterServiceServer(s, implementation)

	lis, err := net.Listen("tcp", ":"+cfg.AppPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	go func() {
		log.Printf("server listening at %v", lis.Addr())
		if err = s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	s.GracefulStop()
	store.Shutdown(ctx)
}
