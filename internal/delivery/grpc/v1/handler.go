package v1

import (
	proto "github.com/AZRV17/goWEB/internal/server/grpc/transfer/v1"
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
	transferServer := NewTransferServer(h.service.TransferService)
	proto.RegisterTransferServiceServer(h.server, transferServer)
}
