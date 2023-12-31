package grpc

import (
	v1 "github.com/AZRV17/goWEB/internal/delivery/grpc/v1"
	"github.com/AZRV17/goWEB/internal/service"
	"google.golang.org/grpc"
)

type Handler struct {
	service service.Service
	server  *grpc.Server
}

func NewHandler(service service.Service, grpcServer *grpc.Server) *Handler {
	return &Handler{
		service: service,
		server:  grpcServer,
	}
}

func (h *Handler) Init() {
	v1 := v1.NewHandler(h.service, h.server)
	v1.Init()
}
