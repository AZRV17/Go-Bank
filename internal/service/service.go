package service

import (
	"time"

	"github.com/AZRV17/goWEB/internal/domain"
	"github.com/AZRV17/goWEB/internal/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type CreateAccountInput struct {
	Owner     string
	Balance   int64
	Currency  string
	CreatedAt time.Time
}

type UpdateAccountInput struct {
	ID        int64
	Owner     string
	Balance   int64
	Currency  string
	CreatedAt time.Time
}

type Account interface {
	CreateAccount(input CreateAccountInput) error
	UpdateAccount(input UpdateAccountInput) error
	DeleteAccount(id int64) error
	GetAccount(id int64) (*domain.Account, error)
	GetAllAccounts() ([]domain.Account, error)
}

type CreateEntryInput struct {
	AccountID int64
	Amount    int64
	CreatedAt time.Time
}

type UpdateEntryInput struct {
	ID        int64
	AccountID int64
	Amount    int64
	CreatedAt time.Time
}

type Entry interface {
	CreateEntry(input CreateEntryInput) error
	UpdateEntry(input UpdateEntryInput) error
	DeleteEntry(id int64) error
	GetEntry(id int64) (*domain.Entry, error)
	GetAllEntries() ([]domain.Entry, error)
}

type CreateTransferInput struct {
	FromAccountID int64
	ToAccountID   int64
	Amount        int64
	CreatedAt     time.Time
}

type UpdateTransferInput struct {
	ID            int64
	FromAccountID int64
	ToAccountID   int64
	Amount        int64
	CreatedAt     time.Time
}

type Transfer interface {
	CreateTransfer(input CreateTransferInput) error
	UpdateTransfer(input UpdateTransferInput) error
	DeleteTransfer(id int64) error
	GetTransfer(id int64) (*domain.Transfer, error)
	GetAllTransfers() ([]domain.Transfer, error)
}

type Service struct {
	AccountService  Account
	EntryService    Entry
	TransferService Transfer
	repo            *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		repo:            repo,
		AccountService:  NewAccountService(repo.Account),
		EntryService:    NewEntryService(repo.Entry),
		TransferService: NewTransferService(repo.Transfer),
	}
}
