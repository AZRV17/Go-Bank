package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/AZRV17/goWEB/internal/domain"
	"github.com/AZRV17/goWEB/pkg/redis"
	"gorm.io/gorm"
	"log"
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

	ctx := context.Background()

	val := redis.Rdb.Get(ctx, "account:"+fmt.Sprint(int(id)))

	if val.Val() != "" {
		err := json.Unmarshal([]byte(val.Val()), &acc)
		if err != nil {
			return nil, err
		}

		log.Println("account found in redis")
		return &acc, nil
	}

	tx := repo.db.Begin()

	if err := tx.First(&acc, id).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	data, err := json.Marshal(acc)
	if err != nil {
		return nil, err
	}

	redis.Rdb.Set(ctx, "account:"+fmt.Sprint(int(id)), data, 0)

	return &acc, nil
}

func (repo *Account) Update(acc *domain.Account) error {
	tx := repo.db.Begin()

	if err := tx.Where("id = ?", acc.ID).Save(acc).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("error updating account: %w", err)
	}

	tx.Commit()

	redis.Rdb.Set(context.Background(), "account:"+fmt.Sprint(acc.ID), acc, 0)

	return nil
}

func (repo *Account) Delete(id int64) error {
	tx := repo.db.Begin()

	if err := tx.Delete(&domain.Account{}, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	redis.Rdb.Del(context.Background(), "account:"+fmt.Sprint(id))

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

func (repo *Account) UpdateAccountBalance(accID int64, amount int64) error {
	tx := repo.db.Begin()

	if err := tx.Where("id = ?", accID).Update("balance", gorm.Expr("balance + ?", amount)).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	redis.Rdb.Set(context.Background(), "account:"+fmt.Sprint(accID), accID, 0)

	return nil
}
