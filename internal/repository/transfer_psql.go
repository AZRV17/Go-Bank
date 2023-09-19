package repository

import (
	"github.com/AZRV17/goWEB/internal/domain"
	"gorm.io/gorm"
)

type Transfer struct {
	db *gorm.DB
}

func NewTransferRepo(db *gorm.DB) *Transfer {
	return &Transfer{db: db}
}

func (repo *Transfer) GetAll() ([]domain.Transfer, error) {
	var transfers []domain.Transfer
	err := repo.db.Find(&transfers).Error
	if err != nil {
		return nil, err
	}

	return transfers, nil
}

func (repo *Transfer) GetTransfer(id int) (*domain.Transfer, error) {
	var transfer domain.Transfer
	err := repo.db.First(&transfer, id).Error
	if err != nil {
		return nil, err
	}

	return &transfer, nil
}

func (repo *Transfer) Create(transfer domain.Transfer) error {
	err := repo.db.Create(&transfer).Error
	return err
}

func (repo *Transfer) Update(transfer domain.Transfer) error {
	err := repo.db.Save(&transfer).Error
	return err
}

func (repo *Transfer) Delete(id int) error {
	err := repo.db.Delete(&domain.Transfer{}, id).Error
	return err
}
