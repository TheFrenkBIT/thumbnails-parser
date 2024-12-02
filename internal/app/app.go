// Package app configures and runs application.
package app

import (
	"fmt"
	parsergrpc "github.com/evrone/go-clean-template/internal/controller/grpc/parser"
	cache "github.com/evrone/go-clean-template/internal/infrastructure/cache"
	"github.com/evrone/go-clean-template/internal/infrastructure/youtube"
	parserService "github.com/evrone/go-clean-template/internal/usecase/services/parser"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/evrone/go-clean-template/config"
	"github.com/evrone/go-clean-template/pkg/logger"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	//Cache
	var redis cache.Interface = cache.New(redis.NewClient(&redis.Options{
		Addr: cfg.Redis.Host + ":" + cfg.Redis.Port,
	}))

	//Youtube client
	var client youtube.Interface = youtube.New(&http.Client{})

	//Use case
	var parserUseCase parserService.Interface = parserService.New(client, redis)

	//gRPC Server
	server := grpc.NewServer()
	parsergrpc.Register(server, parserUseCase)
	n, err := net.Listen("tcp", ":"+cfg.GRPC.Port)
	if err != nil {
		fmt.Printf("app - Run - net.Listen: %s\n", err)
	}

	if err := server.Serve(n); err != nil {
		fmt.Printf("app - Run - server.Serve: %s\n", err)
	}

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	}

	server.GracefulStop()
}
