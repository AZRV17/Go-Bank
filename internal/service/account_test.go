package service_test

import (
	"testing"
	"time"

	"github.com/AZRV17/goWEB/internal/service"
	mock_service "github.com/AZRV17/goWEB/internal/service/mocks"
	"github.com/golang/mock/gomock"
)

func TestAccountService_CreateAccount(t *testing.T) {
	type fields struct {
		service *mock_service.MockAccount
	}
	type args struct {
		input service.CreateAccountInput
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
				f.service.EXPECT().CreateAccount(service.CreateAccountInput{
					Owner:     "Alexa",
					Balance:   1000,
					Currency:  "RUB",
					CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
				}).Return(int64(1), nil)
			},
			args: args{
				input: service.CreateAccountInput{
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
				service: mock_service.NewMockAccount(ctrl),
			}

			if tt.prepare != nil {
				tt.prepare(&f)
			}

			id, err := f.service.CreateAccount(tt.args.input)
			if err != nil != tt.wantErr {
				t.Errorf("AccountService.CreateAccount() error = %v, wantErr %v", err, tt.wantErr)
			}

			if id != 1 {
				t.Errorf("AccountService.CreateAccount() = %v, want %v", id, 1)
			}
		})
	}
}
