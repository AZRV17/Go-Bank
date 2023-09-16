package repository

import (
	"github.com/AZRV17/goWEB/internal/domain"
	"gorm.io/gorm"
)

type AccountRepo struct {
	db *gorm.DB
}

func NewAccountRepo(db *gorm.DB) *AccountRepo {
	return &AccountRepo{db: db.Model(&domain.Account{})}
}
