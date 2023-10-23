package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/AZRV17/goWEB/internal/config"
)

type HttpServer struct {
	httpServer *http.Server
}

func NewHttpServer(cfg *config.Config, handler http.Handler) *HttpServer {
	return &HttpServer{
		httpServer: &http.Server{
			Addr:    ":" + cfg.Server.Port,
			Handler: handler,
		},
	}
}

func (s *HttpServer) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *HttpServer) Shutdown(stopped chan struct{}) {
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-sigint
	log.Printf("got interruption signal")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := s.httpServer.Shutdown(ctx); err != nil {
		log.Printf("HTTP Server Shutdown Error: %v", err)
	}
	close(stopped)
}
