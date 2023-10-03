package service

import (
	"time"

	"github.com/AZRV17/goWEB/internal/domain"
	"github.com/AZRV17/goWEB/internal/repository"
)

type AccountService struct {
	repo repository.Accounts
}

func NewAccountService(repo repository.Accounts) Account {
	return &AccountService{
		repo: repo,
	}
}

func (s *AccountService) CreateAccount(input CreateAccountInput) error {
	account := domain.Account{
		Owner:     input.Owner,
		Balance:   input.Balance,
		Currency:  input.Currency,
		CreatedAt: time.Now(),
	}
	_, err := s.repo.Create(account)
	return err
}

func (s *AccountService) GetAccount(id int64) (*domain.Account, error) {
	acc, err := s.repo.GetAccount(id)
	if err != nil {
		return nil, err
	}
	return acc, nil
}

func (s *AccountService) GetAllAccounts() ([]domain.Account, error) {
	accs, err := s.repo.GetAll()
	return accs, err
}

func (s *AccountService) UpdateAccount(input UpdateAccountInput) error {
	account := domain.Account{
		ID:        input.ID,
		Owner:     input.Owner,
		Balance:   input.Balance,
		Currency:  input.Currency,
		CreatedAt: time.Now(),
	}
	err := s.repo.Update(&account)

	return err
}

func (s *AccountService) DeleteAccount(id int64) error {
	err := s.repo.Delete(id)
	return err
}
