package grpcServer

import (
	"github.com/AZRV17/goWEB/internal/config"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type GrpcServer struct {
	GrpcServer *grpc.Server
	host       string
	port       string
}

func NewGrpcServer(server *grpc.Server, cfg *config.Config) *GrpcServer {
	return &GrpcServer{
		GrpcServer: server,
		host:       cfg.GRPC.Host,
		port:       cfg.GRPC.Port,
	}
}

func (s *GrpcServer) Run() error {
	lis, err := net.Listen("tcp", s.host+":"+s.port)
	if err != nil {
		return err
	}

	return s.GrpcServer.Serve(lis)
}

func (s *GrpcServer) Shutdown() {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	<-sigCh
	log.Println("grpc server shutting down")
	s.GrpcServer.GracefulStop()
}
