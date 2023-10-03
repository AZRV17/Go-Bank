package service

import (
	"time"

	"github.com/AZRV17/goWEB/internal/domain"
	"github.com/AZRV17/goWEB/internal/repository"
)

type TransferService struct {
	repo repository.Transfers
}

func NewTransferService(repo repository.Transfers) *TransferService {
	return &TransferService{
		repo: repo,
	}
}

func (s *TransferService) CreateTransfer(input CreateTransferInput) error {
	transfer := domain.Transfer{
		FromAccountID: input.FromAccountID,
		ToAccountID:   input.ToAccountID,
		Amount:        input.Amount,
		CreatedAt:     time.Now(),
	}
	_, err := s.repo.Create(transfer)
	return err
}

func (s *TransferService) GetTransfer(id int64) (*domain.Transfer, error) {
	transfer, err := s.repo.GetTransfer(id)
	return transfer, err
}

func (s *TransferService) GetAllTransfers() ([]domain.Transfer, error) {
	transfers, err := s.repo.GetAll()
	return transfers, err
}

func (s *TransferService) UpdateTransfer(input UpdateTransferInput) error {
	transfer := domain.Transfer{
		ID:            input.ID,
		FromAccountID: input.FromAccountID,
		ToAccountID:   input.ToAccountID,
		Amount:        input.Amount,
		CreatedAt:     time.Now(),
	}
	err := s.repo.Update(transfer)

	return err
}

func (s *TransferService) DeleteTransfer(id int64) error {
	err := s.repo.Delete(id)
	return err
}
