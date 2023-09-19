package repository

import (
	"gorm.io/gorm"
)

type Repositories struct {
	Account  Account
	Entry    Entry
	Transfer Transfer
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		Account:  *NewAccountRepo(db),
		Entry:    *NewEntryRepo(db),
		Transfer: *NewTransferRepo(db),
	}
}
