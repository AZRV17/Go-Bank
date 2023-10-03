package repository

import (
	"github.com/AZRV17/goWEB/internal/domain"
	"gorm.io/gorm"
)

type Entry struct {
	db *gorm.DB
}

func NewEntryRepo(db *gorm.DB) Entries {
	return &Entry{db: db.Model(&domain.Entry{})}
}

func (repo *Entry) GetAll() ([]domain.Entry, error) {
	var entries []domain.Entry
	err := repo.db.Find(&entries).Error
	if err != nil {
		return nil, err
	}

	return entries, nil
}

func (repo *Entry) GetEntry(id int64) (*domain.Entry, error) {
	var entry domain.Entry
	err := repo.db.First(&entry, id).Error
	if err != nil {
		return nil, err
	}

	return &entry, nil
}

func (repo *Entry) Create(entry domain.Entry) (*domain.Entry, error) {
	err := repo.db.Create(&entry).Error
	if err != nil {
		return nil, err
	}

	repo.db.Last(entry)
	return &entry, err
}

func (repo *Entry) Update(entry domain.Entry) error {
	err := repo.db.Save(&entry).Error
	return err
}

func (repo *Entry) Delete(id int64) error {
	err := repo.db.Delete(&domain.Entry{}, id).Error
	return err
}
