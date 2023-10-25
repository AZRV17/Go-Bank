package service

import (
	"time"

	"github.com/AZRV17/goWEB/internal/domain"
	"github.com/AZRV17/goWEB/internal/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type CreateAccountInput struct {
	Owner     string    `json:"owner"`
	Balance   int64     `json:"balance"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"created_at"`
}

type UpdateAccountInput struct {
	ID        int64     `json:"id"`
	Owner     string    `json:"owner"`
	Balance   int64     `json:"balance"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"created_at"`
}

type Account interface {
	CreateAccount(input CreateAccountInput) error
	UpdateAccount(input UpdateAccountInput) error
	DeleteAccount(id int64) error
	GetAccount(id int64) (*domain.Account, error)
	GetAllAccounts() ([]domain.Account, error)
}

type CreateEntryInput struct {
	AccountID int64     `json:"account_id"`
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

type UpdateEntryInput struct {
	ID        int64 `json:"id"`
	AccountID int64 `json:"account_id"`
	Amount    int64 `json:"amount"`
	CreatedAt time.Time
}

type Entry interface {
	GetEntry(id int64) (*domain.Entry, error)
	GetAllEntries() ([]domain.Entry, error)
}

type CreateTransferInput struct {
	FromAccountID int64     `json:"from_account_id"`
	ToAccountID   int64     `json:"to_account_id"`
	Amount        int64     `json:"amount"`
	CreatedAt     time.Time `json:"created_at"`
}

type UpdateTransferInput struct {
	ID            int64     `json:"id"`
	FromAccountID int64     `json:"from_account_id"`
	ToAccountID   int64     `json:"to_account_id"`
	Amount        int64     `json:"amount"`
	CreatedAt     time.Time `json:"created_at"`
}

type Transfer interface {
	CreateTransfer(input CreateTransferInput) error
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
		TransferService: NewTransferService(repo.Transfer, repo.Account, repo.Entry),
	}
}
