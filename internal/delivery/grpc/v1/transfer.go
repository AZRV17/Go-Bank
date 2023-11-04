package v1

import (
	"context"
	proto "github.com/AZRV17/goWEB/internal/server/grpc/transfer/v1"
	"github.com/AZRV17/goWEB/internal/service"
	"log"
	"strconv"
)

type TransferServer struct {
	service service.Transfer
}

func NewTransferServer(service service.Transfer) *TransferServer {
	return &TransferServer{
		service: service,
	}
}

func (s *TransferServer) CreateTransfer(ctx context.Context, request *proto.TransferRequest) (*proto.TransferResponse, error) {
	transfer := &proto.Transfer{
		FromAccountId: request.FromAccountId,
		ToAccountId:   request.ToAccountId,
		Amount:        request.Amount,
	}

	fromId, err := strconv.Atoi(transfer.FromAccountId)
	if err != nil {
		log.Println(transfer)
		return nil, err
	}

	toId, err := strconv.Atoi(transfer.ToAccountId)
	if err != nil {
		return nil, err
	}

	err = s.service.CreateTransfer(service.CreateTransferInput{
		FromAccountID: int64(fromId),
		ToAccountID:   int64(toId),
		Amount:        transfer.Amount,
	})
	if err != nil {
		return nil, err
	}

	return &proto.TransferResponse{
		Transfer: transfer,
	}, nil
}
