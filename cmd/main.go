package main

import (
	"github.com/joho/godotenv"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"link_shorter/internal/config"
	"link_shorter/internal/handler"
	"link_shorter/internal/pkg/service"
	"link_shorter/internal/pkg/storage"
	"link_shorter/internal/pkg/storage/inmemory"
	"link_shorter/internal/pkg/storage/postgres"
	"link_shorter/internal/protobuf/link_shorter/protobuf/shorter"
	"log"
	"net"
	"os"
	"os/signal"
	"time"
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
	if storeType == "inmemory" {
		store = inmemory.NewInMemory()
	} else if storeType == "postgresql" {
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

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, os.Kill)
	<-quit

	s.GracefulStop()
	store.Shutdown(ctx)
}
