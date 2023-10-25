package service

import (
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

func (s *EntryService) GetEntry(id int64) (*domain.Entry, error) {
	entry, err := s.repo.GetEntry(id)
	return entry, err
}

func (s *EntryService) GetAllEntries() ([]domain.Entry, error) {
	entries, err := s.repo.GetAll()
	return entries, err
}
