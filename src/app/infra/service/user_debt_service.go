package service

import (
	"rapicreds-backend/src/app/domain"
	"rapicreds-backend/src/app/infra/apierror"
	"rapicreds-backend/src/app/infra/restclient"
)

type IUserDebtService interface {
	GetUserDebt(document uint64) (*domain.UserDebt, *apierror.ApiError)
}

type BaseUserDebtService struct {
	userDebtRestClient restclient.IUserDebtRestClient
}

func NewUserDebtService(userDebtRestClient restclient.IUserDebtRestClient) IUserDebtService {
	return &BaseUserDebtService{
		userDebtRestClient: userDebtRestClient,
	}
}

func (b *BaseUserDebtService) GetUserDebt(document uint64) (*domain.UserDebt, *apierror.ApiError) {
	userDebt, err := b.userDebtRestClient.GetUserDebt(document)

	if err != nil {
		err := apierror.NewCustomErrorWithStatus(
			"BaseUserDebtService - error fetching userDebt from RestClient",
			err.Err,
			err.Status,
		)
		return nil, err
	}

	return userDebt, nil
}
