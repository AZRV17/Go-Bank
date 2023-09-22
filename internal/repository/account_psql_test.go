package repository_test

import (
	"regexp"
	"testing"
	"time"

	"github.com/AZRV17/goWEB/internal/domain"
	"github.com/AZRV17/goWEB/internal/repository"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func createMockDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock, repository.Account) {
	mockDb, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock db: %v", err)
	}
	dialector := postgres.New(postgres.Config{
		DriverName:           "postgres",
		Conn:                 mockDb,
		PreferSimpleProtocol: true,
	})
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open gorm connection: %v", err)
	}

	r := repository.NewRepositories(db)
	repo := r.Account

	return db, mock, repo
}

func TestCreateAccount(t *testing.T) {
	db, mock, repo := createMockDB(t)
	defer func() {
		Db, _ := db.DB()
		Db.Close()
	}()

	account := domain.Account{
		Owner:     "Alexa",
		Balance:   1000,
		Currency:  "RUB",
		CreatedAt: time.Now(),
	}

	rows := sqlmock.NewRows([]string{"id", "owner", "balance", "currency", "created_at"}).AddRow(1, "Alexa", 1000, "RUB", time.Now())

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "accounts" ("owner","balance","currency","created_at") VALUES ($1,$2,$3,$4) RETURNING "id"`)).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1)).WithArgs(account.Owner, account.Balance, account.Currency, account.CreatedAt)
	mock.ExpectCommit()
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "accounts" ORDER BY "accounts"."id" DESC LIMIT 1`)).
		WillReturnRows(rows)

	_, err := repo.Create(account)
	assert.NoError(t, err)
}

func TestGetAccount(t *testing.T) {
	db, mock, repo := createMockDB(t)
	defer func() {
		Db, _ := db.DB()
		Db.Close()
	}()

	rows := sqlmock.NewRows([]string{"id", "owner", "balance", "currency", "created_at"}).AddRow(1, "Alexa", 1000, "RUB", time.Now())

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "accounts" WHERE "accounts"."id" = $1 ORDER BY "accounts"."id" LIMIT 1`)).
		WillReturnRows(rows).WithArgs(1)

	_, err := repo.GetAccount(1)
	assert.NoError(t, err)
}

func TestGetAllAccounts(t *testing.T) {
	db, mock, repo := createMockDB(t)
	defer func() {
		Db, _ := db.DB()
		Db.Close()
	}()

	rows := sqlmock.NewRows([]string{"id", "owner", "balance", "currency", "created_at"}).AddRow(1, "Alexa", 1000, "RUB", time.Now()).
		AddRow(2, "Bob", 2000, "RUB", time.Now())

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "accounts"`)).WillReturnRows(rows)

	_, err := repo.GetAllAccounts()
	assert.NoError(t, err)
}

func TestDeleteAccount(t *testing.T) {
	db, mock, repo := createMockDB(t)
	defer func() {
		Db, _ := db.DB()
		Db.Close()
	}()

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "accounts" WHERE "accounts"."id" = $1`)).WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := repo.Delete(1)
	assert.NoError(t, err)
}

func TestUpdateAccount(t *testing.T) {
	db, mock, repo := createMockDB(t)
	defer func() {
		Db, _ := db.DB()
		Db.Close()
	}()

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "accounts" SET "id"=$1,"owner"=$2,"balance"=$3,"currency"=$4,"created_at"=$5 WHERE id = $6`)).
		WithArgs(1, "Alexa", 2000, "RUB", time.Now(), 1).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := repo.Update(&domain.Account{
		ID:        1,
		Owner:     "Alexa",
		Balance:   2000,
		Currency:  "RUB",
		CreatedAt: time.Now(),
	})

	assert.NoError(t, err)
}
