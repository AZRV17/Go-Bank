package repository

import (
	"fmt"
	"github.com/AZRV17/goWEB/internal/domain"
	"gorm.io/gorm"
)

type Account struct {
	db *gorm.DB
}

func NewAccountRepo(db *gorm.DB) *Account {
	return &Account{db: db.Model(&domain.Account{})}
}

func (repo *Account) Create(account domain.Account) (*domain.Account, error) {
	tx := repo.db.Begin()

	if err := tx.Create(&account).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return &account, nil
}

func (repo *Account) GetAccount(id int64) (*domain.Account, error) {
	var acc domain.Account
	tx := repo.db.Begin()

	if err := tx.First(&acc, id).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return &acc, nil
}

func (repo *Account) Update(acc *domain.Account) error {
	tx := repo.db.Begin()

	if err := tx.Where("id = ?", acc.ID).Save(acc).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("error updating account: %w", err)
	}

	tx.Commit()

	return nil
}

func (repo *Account) Delete(id int64) error {
	tx := repo.db.Begin()

	if err := tx.Delete(&domain.Account{}, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (repo *Account) GetAll() ([]domain.Account, error) {
	tx := repo.db.Begin()

	var accs []domain.Account
	err := tx.Find(&accs).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return accs, nil
}

func (repo *Account) AddAccountBalance(accID int64, amount int64) error {
	tx := repo.db.Begin()

	if err := tx.Where("id = ?", accID).Update("balance", gorm.Expr("balance + ?", amount)).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
