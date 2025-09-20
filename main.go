package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/tvoybuket/auth/api"
	"github.com/tvoybuket/auth/config"
	desc "github.com/tvoybuket/auth/pkg/auth_v1"
	"github.com/tvoybuket/tblib/tblogger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	settings := config.MustLoad()
	loggerSetting := config.MustLoadLoggerSettings()
	logger, err := tblogger.New(loggerSetting)
	if err != nil {
		tblogger.Fatal("failed to init logger: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", settings.Port))
	if err != nil {
		tblogger.Fatal("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterAuthServiceServer(s, api.NewServer(logger))

	go func() {
		logger.Info("starting gRPC server on port %v", lis.Addr())
		if err := s.Serve(lis); err != nil {
			logger.Error("gRPC server failed %v", err)
			cancel()
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	select {
	case sig := <-sigChan:
		logger.Info("received shutdown signal %v", sig)
	case <-ctx.Done():
		logger.Info("context cancelled")
	}

	// Graceful shutdown
	logger.Info("shutting down server")

	// Создаем контекст с таймаутом для shutdown
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), time.Duration(settings.ShutdownTimeout))
	defer shutdownCancel()

	// Останавливаем gRPC сервер
	done := make(chan struct{})
	go func() {
		s.GracefulStop()
		close(done)
	}()

	select {
	case <-done:
		logger.Info("server stopped gracefully")
	case <-shutdownCtx.Done():
		logger.Warn("shutdown timeout exceeded, forcing stop")
		s.Stop()
	}

	logger.Info("auth service stopped")
}
