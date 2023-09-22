package service

import (
	"time"

	"github.com/AZRV17/goWEB/internal/repository"
)

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
	CreateAccount(input CreateAccountInput) (int64, error)
	UpdateAccount(input UpdateAccountInput) error
	DeleteAccount(id int64) error
	GetAccount(id int64) (repository.Account, error)
	GetAllAccounts() ([]repository.Account, error)
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
	CreateEntry(input CreateEntryInput) (int64, error)
	UpdateEntry(input UpdateEntryInput) error
	DeleteEntry(id int64) error
	GetEntry(id int64) (repository.Entry, error)
	GetAllEntries() ([]repository.Entry, error)
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

type TransferServiceInterface interface {
	CreateTransfer(input CreateTransferInput) (int64, error)
	UpdateTransfer(input UpdateTransferInput) error
	DeleteTransfer(id int64) error
	GetTransfer(id int64) (repository.Transfer, error)
	GetAllTransfers() ([]repository.Transfer, error)
}

type Service struct {
	AccountService  *AccountService
	EntryService    *EntryService
	TransferService *TransferService
	repo            *repository.Repositories
}

func NewService(repo *repository.Repositories) *Service {
	return &Service{
		repo:            repo,
		AccountService:  NewAccountService(repo.Account),
		EntryService:    NewEntryService(repo.Entry),
		TransferService: NewTransferService(repo.Transfer),
	}
}
