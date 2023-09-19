package service

import (
	"github.com/AZRV17/goWEB/internal/domain"
	"github.com/AZRV17/goWEB/internal/repository"
)

type AccountService struct {
	repo repository.Account
}

func NewAccountService(repo repository.Account) *AccountService {
	return &AccountService{
		repo: repo,
	}
}

func (s *AccountService) CreateAccount(account domain.Account) error {
	err := s.repo.CreateAccount(account)
	return err
}

func (s *AccountService) GetAccount(id int64) (domain.Account, error) {
	acc, err := s.repo.GetAccount(id)
	return *acc, err
}

func (s *AccountService) GetAllAccounts() ([]domain.Account, error) {
	accs, err := s.repo.GetAllAccounts()
	return accs, err
}

func (s *AccountService) UpdateAccount(id int64) error {
	acc, err := s.repo.GetAccount(id)
	if err != nil {
		return err
	}

	err = s.repo.UpdateAccount(*acc)

	return err
}

func (s *AccountService) DeleteAccount(id int64) error {
	err := s.repo.DeleteAccount(id)
	return err
}
