package repository

import (
	"github.com/AZRV17/goWEB/internal/domain"
	"gorm.io/gorm"
)

type Account struct {
	db *gorm.DB
}

func NewAccountRepo(db *gorm.DB) *Account {
	return &Account{db: db.Model(&domain.Account{})}
}

func (repo *Account) CreateAccount(account domain.Account) error {
	err := repo.db.Create(&account).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *Account) GetAccount(id int64) (*domain.Account, error) {
	var account domain.Account
	err := repo.db.First(&account, id).Error
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (repo *Account) UpdateAccount(account domain.Account) error {
	err := repo.db.Save(&account).Error
	return err
}

func (repo *Account) DeleteAccount(id int64) error {
	err := repo.db.Delete(&domain.Account{}, id).Error
	return err
}

func (repo *Account) GetAllAccounts() ([]domain.Account, error) {
	var accounts []domain.Account

	err := repo.db.Find(&accounts).Error
	if err != nil {
		return nil, err
	}

	return accounts, nil
}
