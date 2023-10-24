package service

import (
	"log"
	"time"

	"github.com/AZRV17/goWEB/internal/domain"
	"github.com/AZRV17/goWEB/internal/repository"
)

type TransferService struct {
	repo    repository.Transfers
	accRepo repository.Accounts
}

func NewTransferService(repo repository.Transfers, accRepo repository.Accounts) *TransferService {
	return &TransferService{
		repo:    repo,
		accRepo: accRepo,
	}
}

func (s *TransferService) CreateTransfer(input CreateTransferInput) error {
	transfer := domain.Transfer{
		FromAccountID: input.FromAccountID,
		ToAccountID:   input.ToAccountID,
		Amount:        input.Amount,
		CreatedAt:     time.Now(),
	}

	if _, err := s.repo.Create(transfer); err != nil {
		log.Println(err)
		return err
	}

	err := s.addMoney(input.FromAccountID, input.ToAccountID, input.Amount)

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

func (s *TransferService) addMoney(fromAccountID, toAccountID int64, amount int64) error {
	if err := s.accRepo.AddAccountBalance(fromAccountID, -amount); err != nil {
		return err
	}

	if err := s.accRepo.AddAccountBalance(toAccountID, amount); err != nil {
		return err
	}

	return nil
}
