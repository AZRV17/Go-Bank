package repository

import (
	"github.com/AZRV17/goWEB/internal/domain"
	"gorm.io/gorm"
)

//go:generate mockgen -source=repository.go -destination=mocks/mock.go

type Accounts interface {
	Create(account domain.Account) (*domain.Account, error)
	GetAccount(id int64) (*domain.Account, error)
	Update(account *domain.Account) error
	Delete(id int64) error
	GetAll() ([]domain.Account, error)
}

type Entries interface {
	Create(entry domain.Entry) (*domain.Entry, error)
	GetEntry(id int64) (*domain.Entry, error)
	Update(entry domain.Entry) error
	Delete(id int64) error
	GetAll() ([]domain.Entry, error)
}

type Transfers interface {
	Create(transfer domain.Transfer) (*domain.Transfer, error)
	GetTransfer(id int64) (*domain.Transfer, error)
	Update(transfer domain.Transfer) error
	Delete(id int64) error
	GetAll() ([]domain.Transfer, error)
}

type Repositories struct {
	Account  Accounts
	Entry    Entries
	Transfer Transfers
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		Account:  NewAccountRepo(db),
		Entry:    NewEntryRepo(db),
		Transfer: NewTransferRepo(db),
	}
}
