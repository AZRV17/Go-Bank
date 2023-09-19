package service

import (
	"time"

	"github.com/AZRV17/goWEB/internal/domain"
	"github.com/AZRV17/goWEB/internal/repository"
)

type TransferService struct {
	repo repository.Transfer
}

func NewTransferService(repo repository.Transfer) *TransferService {
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
	err := s.repo.Create(transfer)
	return err
}

func (s *TransferService) GetTransfer(id int) (domain.Transfer, error) {
	transfer, err := s.repo.GetTransfer(id)
	return *transfer, err
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

func (s *TransferService) DeleteTransfer(id int) error {
	err := s.repo.Delete(id)
	return err
}
