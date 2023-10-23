package repository

import (
	"github.com/AZRV17/goWEB/internal/domain"
	"gorm.io/gorm"
)

type Entry struct {
	db *gorm.DB
}

func NewEntryRepo(db *gorm.DB) *Entry {
	return &Entry{db: db.Model(&domain.Entry{})}
}

func (repo *Entry) GetAll() ([]domain.Entry, error) {
	var entries []domain.Entry

	tx := repo.db.Begin()

	if err := tx.Find(&entries).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return entries, nil
}

func (repo *Entry) GetEntry(id int64) (*domain.Entry, error) {
	var entry domain.Entry

	tx := repo.db.Begin()

	if err := tx.First(&entry, id).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return &entry, nil
}

func (repo *Entry) Create(entry domain.Entry) (*domain.Entry, error) {
	tx := repo.db.Begin()

	if err := tx.Create(&entry).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	repo.db.Last(entry)
	return &entry, nil
}

func (repo *Entry) Update(entry domain.Entry) error {
	tx := repo.db.Begin()

	if err := tx.Where("id = ?", entry.ID).Save(&entry).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (repo *Entry) Delete(id int64) error {
	tx := repo.db.Begin()

	if err := tx.Delete(&domain.Entry{}, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
