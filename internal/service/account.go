package service

import (
	"github.com/AZRV17/goWEB/internal/domain"
	"github.com/AZRV17/goWEB/internal/repository"
	"github.com/AZRV17/goWEB/pkg/db/psql"
)

type AccountService struct {
	repo repository.Account
}

func NewAccountService() *AccountService {
	return &AccountService{
		repo: *repository.NewAccountRepo(psql.DB),
	}
}

func (s *AccountService) CreateAccount(owner string, balance int64, currency string) error {
	err := s.repo.CreateAccount(owner, balance, currency)
	return err
}

func (s *AccountService) GetAccount(id int64) (domain.Account, error) {
	acc, err := s.repo.GetAccount(id)
	return acc, err
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

	err = s.repo.UpdateAccount(acc)

	return err
}

func (s *AccountService) DeleteAccount(id int64) error {
	err := s.repo.DeleteAccount(id)
	return err
}
