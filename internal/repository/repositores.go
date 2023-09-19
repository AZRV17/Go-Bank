package repository

import (
	"gorm.io/gorm"
)

type Repositories struct {
	Account Account
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		Account: *NewAccountRepo(db),
	}
}
