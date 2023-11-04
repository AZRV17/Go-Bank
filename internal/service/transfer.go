package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/AZRV17/goWEB/internal/domain"
	"github.com/AZRV17/goWEB/internal/repository"
)

type TransferService struct {
	repo      repository.Transfers
	accRepo   repository.Accounts
	entryRepo repository.Entries
}

func NewTransferService(repo repository.Transfers, accRepo repository.Accounts, entryRepo repository.Entries) *TransferService {
	return &TransferService{
		repo:      repo,
		accRepo:   accRepo,
		entryRepo: entryRepo,
	}
}

func (s *TransferService) CreateTransfer(input CreateTransferInput) error {
	transfer := domain.Transfer{
		FromAccountID: input.FromAccountID,
		ToAccountID:   input.ToAccountID,
		Amount:        input.Amount,
		CreatedAt:     time.Now(),
	}

	if err := s.addMoney(input.FromAccountID, input.ToAccountID, input.Amount); err != nil {
		return err
	}

	if _, err := s.repo.Create(transfer); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (s *TransferService) GetTransfer(id int64) (*domain.Transfer, error) {
	transfer, err := s.repo.GetTransfer(id)
	return transfer, err
}

func (s *TransferService) GetAllTransfers() ([]domain.Transfer, error) {
	transfers, err := s.repo.GetAll()
	return transfers, err
}

func (s *TransferService) addMoney(fromAccountID, toAccountID int64, amount int64) error {
	fromAccount, err := s.accRepo.GetAccount(fromAccountID)
	if err != nil {
		return err
	}

	toAccount, err := s.accRepo.GetAccount(toAccountID)
	if err != nil {
		return err
	}

	if fromAccount.Balance < amount {
		return errors.New("not enough money on account")
	}

	amountTo, err := calculateAmountByCurrency(amount, fromAccount.Currency, toAccount.Currency)
	if err != nil {
		return err
	}

	if err := s.accRepo.UpdateAccountBalance(fromAccountID, -amount); err != nil {
		return err
	}

	if err := s.accRepo.UpdateAccountBalance(toAccountID, amountTo); err != nil {
		return err
	}

	entryFrom := CreateEntryInput{
		AccountID: fromAccountID,
		Amount:    amount * -1,
		CreatedAt: time.Now(),
	}

	entryTo := CreateEntryInput{
		AccountID: toAccountID,
		Amount:    amount,
		CreatedAt: time.Now(),
	}

	if err := s.createEntry(entryFrom); err != nil {
		return err
	}

	if err := s.createEntry(entryTo); err != nil {
		return err
	}

	return nil
}

type Currency struct {
	Rates map[string]float32 `json:"rates"`
}

func calculateAmountByCurrency(amount int64, currencyFrom, currencyTo string) (int64, error) {
	url := fmt.Sprintf(`https://openexchangerates.org/api/latest.json?app_id=e3b53564d7c64c0c8afa76b84d01403d&base=%s&symbols=%s`,
		currencyFrom, currencyTo)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(res.Body)
	body, _ := io.ReadAll(res.Body)

	var rates Currency
	if err := json.Unmarshal(body, &rates); err != nil {
		return 0, err
	}

	if rate, ok := rates.Rates[currencyTo]; ok {
		return amount * int64(rate), nil
	}

	return 0, errors.New("invalid currency")
}

func (s *TransferService) createEntry(input CreateEntryInput) error {
	entry := domain.Entry{
		AccountID: input.AccountID,
		Amount:    input.Amount,
		CreatedAt: time.Now(),
	}
	_, err := s.entryRepo.Create(entry)
	return err
}
