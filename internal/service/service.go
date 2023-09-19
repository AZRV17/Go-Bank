package service

type Service struct {
	AccountService *AccountService
}

func NewService() *Service {
	return &Service{
		AccountService: NewAccountService(),
	}
}
