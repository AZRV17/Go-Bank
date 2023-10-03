package repository

import (
	"log"

	"github.com/AZRV17/goWEB/internal/domain"
	"gorm.io/gorm"
)

type Account struct {
	db *gorm.DB
}

func NewAccountRepo(db *gorm.DB) Accounts {
	return &Account{db: db.Model(&domain.Account{})}
}

func (repo *Account) Create(account domain.Account) (*domain.Account, error) {
	err := repo.db.Create(&account).Error
	if err != nil {
		return nil, err
	}

	var acc domain.Account

	repo.db.Last(&acc)

	return &acc, nil
}

func (repo *Account) GetAccount(id int64) (*domain.Account, error) {
	var account domain.Account
	err := repo.db.First(&account, id).Error
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (repo *Account) Update(account *domain.Account) error {
	result := repo.db.Save(account)
	if result.Error != nil {
		log.Println("Error updating account:", result.Error)
		return result.Error
	}
	return nil
}

func (repo *Account) Delete(id int64) error {
	err := repo.db.Delete(&domain.Account{}, id).Error
	return err
}

func (repo *Account) GetAll() ([]domain.Account, error) {
	var accounts []domain.Account

	err := repo.db.Find(&accounts).Error
	if err != nil {
		return nil, err
	}

	return accounts, nil
}
