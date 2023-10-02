package repository_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/AZRV17/goWEB/internal/domain"
	mock_repository "github.com/AZRV17/goWEB/internal/repository/mocks"
	"github.com/golang/mock/gomock"
)

func TestAccount_Create(t *testing.T) {
	type fields struct {
		repo *mock_repository.MockAccounts
	}
	type args struct {
		account domain.Account
	}
	tests := []struct {
		name    string
		prepare func(f *fields)
		args    args
		want    *domain.Account
		wantErr bool
	}{
		{
			name: "success",
			prepare: func(f *fields) {
				f.repo.EXPECT().Create(domain.Account{
					ID:        1,
					Owner:     "Alexa",
					Balance:   1000,
					Currency:  "RUB",
					CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
				}).Return(&domain.Account{
					ID:        1,
					Owner:     "Alexa",
					Balance:   1000,
					Currency:  "RUB",
					CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
				}, nil)
			},
			args: args{
				account: domain.Account{
					ID:        1,
					Owner:     "Alexa",
					Balance:   1000,
					Currency:  "RUB",
					CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
				},
			},
			want: &domain.Account{
				ID:        1,
				Owner:     "Alexa",
				Balance:   1000,
				Currency:  "RUB",
				CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				repo: mock_repository.NewMockAccounts(ctrl),
			}

			if tt.prepare != nil {
				tt.prepare(&f)
			}

			account, err := f.repo.Create(tt.args.account)

			if err != nil {
				t.Errorf("AccountService.CreateAccount() error = %v", err)
			}

			if account.ID != tt.want.ID {
				t.Errorf("AccountService.CreateAccount() = %v, want %v", account, tt.want)
			}
		})
	}
}

func TestAccount_GetAccount(t *testing.T) {
	type fields struct {
		repo *mock_repository.MockAccounts
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		prepare func(f *fields)
		args    args
		want    *domain.Account
		wantErr bool
	}{
		{
			name: "success",
			prepare: func(f *fields) {
				f.repo.EXPECT().GetAccount(int64(1)).Return(&domain.Account{
					ID:        1,
					Owner:     "Alexa",
					Balance:   1000,
					Currency:  "RUB",
					CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
				}, nil)
			},
			args: args{
				id: 1,
			},
			want: &domain.Account{
				ID:        1,
				Owner:     "Alexa",
				Balance:   1000,
				Currency:  "RUB",
				CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				repo: mock_repository.NewMockAccounts(ctrl),
			}

			if tt.prepare != nil {
				tt.prepare(&f)
			}

			account, err := f.repo.GetAccount(tt.args.id)

			if err != nil {
				t.Errorf("AccountService.GetAccount() error = %v", err)
			}

			if account.ID != tt.want.ID {
				t.Errorf("AccountService.GetAccount() = %v, want %v", account, tt.want)
			}
		})
	}
}

func TestAccount_Update(t *testing.T) {
	type fields struct {
		repo *mock_repository.MockAccounts
	}
	type args struct {
		account *domain.Account
	}
	tests := []struct {
		name    string
		prepare func(f *fields)
		args    args
		wantErr bool
	}{
		{
			name: "success",
			prepare: func(f *fields) {
				f.repo.EXPECT().Update(&domain.Account{
					ID:        1,
					Owner:     "Alexa",
					Balance:   1000,
					Currency:  "RUB",
					CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
				}).Return(nil)
			},
			args: args{
				account: &domain.Account{
					ID:        1,
					Owner:     "Alexa",
					Balance:   1000,
					Currency:  "RUB",
					CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				repo: mock_repository.NewMockAccounts(ctrl),
			}

			if tt.prepare != nil {
				tt.prepare(&f)
			}

			err := f.repo.Update(tt.args.account)

			if err != nil {
				t.Errorf("AccountService.UpdateAccount() error = %v", err)
			}
		})
	}
}

func TestAccount_Delete(t *testing.T) {
	type fields struct {
		repo *mock_repository.MockAccounts
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		prepare func(f *fields)
		args    args
		wantErr bool
	}{
		{
			name: "success",
			prepare: func(f *fields) {
				f.repo.EXPECT().Delete(int64(1)).Return(nil)
			},
			args: args{
				id: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				repo: mock_repository.NewMockAccounts(ctrl),
			}

			if tt.prepare != nil {
				tt.prepare(&f)
			}

			err := f.repo.Delete(tt.args.id)

			if err != nil {
				t.Errorf("AccountService.DeleteAccount() error = %v", err)
			}
		})
	}
}

func TestAccount_GetAll(t *testing.T) {
	type fields struct {
		repo *mock_repository.MockAccounts
	}
	tests := []struct {
		name    string
		prepare func(f *fields)
		want    []domain.Account
		wantErr bool
	}{
		{
			name: "success",
			prepare: func(f *fields) {
				f.repo.EXPECT().GetAll().Return([]domain.Account{
					{
						ID:        1,
						Owner:     "Alexa",
						Balance:   1000,
						Currency:  "RUB",
						CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
					},
					{
						ID:        2,
						Owner:     "Bob",
						Balance:   2000,
						Currency:  "RUB",
						CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
					},
				}, nil)
			},
			want: []domain.Account{
				{
					ID:        1,
					Owner:     "Alexa",
					Balance:   1000,
					Currency:  "RUB",
					CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
				},
				{
					ID:        2,
					Owner:     "Bob",
					Balance:   2000,
					Currency:  "RUB",
					CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				repo: mock_repository.NewMockAccounts(ctrl),
			}

			if tt.prepare != nil {
				tt.prepare(&f)
			}

			got, err := f.repo.GetAll()

			if (err != nil) != tt.wantErr {
				t.Errorf("AccountService.GetAll() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccountService.GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
