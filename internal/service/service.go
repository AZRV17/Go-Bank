package service

import "github.com/AZRV17/goWEB/internal/repository"

type Service struct {
	AccountService *AccountService
	repo           *repository.Repositories
}

func NewService(repo *repository.Repositories) *Service {
	return &Service{
		repo:           repo,
		AccountService: NewAccountService(repo.Account),
	}
}
