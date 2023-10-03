package service

import (
	"time"

	"github.com/AZRV17/goWEB/internal/domain"
	"github.com/AZRV17/goWEB/internal/repository"
)

type EntryService struct {
	repo repository.Entries
}

func NewEntryService(repo repository.Entries) *EntryService {
	return &EntryService{
		repo: repo,
	}
}

func (s *EntryService) CreateEntry(input CreateEntryInput) error {
	entry := domain.Entry{
		AccountID: input.AccountID,
		Amount:    input.Amount,
		CreatedAt: time.Now(),
	}
	_, err := s.repo.Create(entry)
	return err
}

func (s *EntryService) GetEntry(id int64) (*domain.Entry, error) {
	entry, err := s.repo.GetEntry(id)
	return entry, err
}

func (s *EntryService) GetAllEntries() ([]domain.Entry, error) {
	entries, err := s.repo.GetAll()
	return entries, err
}

func (s *EntryService) UpdateEntry(input UpdateEntryInput) error {
	entry := domain.Entry{
		ID:        input.ID,
		AccountID: input.AccountID,
		Amount:    input.Amount,
		CreatedAt: time.Now(),
	}
	err := s.repo.Update(entry)

	return err
}

func (s *EntryService) DeleteEntry(id int64) error {
	err := s.repo.Delete(id)
	return err
}
