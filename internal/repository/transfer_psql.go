package repository

import (
	"github.com/AZRV17/goWEB/internal/domain"
	"gorm.io/gorm"
)

type Transfer struct {
	db *gorm.DB
}

func NewTransferRepo(db *gorm.DB) *Transfer {
	return &Transfer{db: db.Model(domain.Transfer{})}
}

func (repo *Transfer) GetAll() ([]domain.Transfer, error) {
	tx := repo.db.Begin()
	defer tx.Rollback()

	var transfers []domain.Transfer
	err := tx.Find(&transfers).Error
	if err != nil {
		return nil, err
	}

	tx.Commit()
	return transfers, nil
}

func (repo *Transfer) GetTransfer(id int64) (*domain.Transfer, error) {
	tx := repo.db.Begin()

	var transfer domain.Transfer
	if err := tx.First(&transfer, id).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return &transfer, nil
}

func (repo *Transfer) Create(transfer domain.Transfer) (*domain.Transfer, error) {

	tx := repo.db.Begin()

	if err := tx.Create(&transfer).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	//repo.db.Last(transfer)
	return &transfer, nil
}
