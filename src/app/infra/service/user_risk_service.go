package service

import (
	"rapicreds-backend/src/app/domain"
	"rapicreds-backend/src/app/infra/apierror"
)

// IUserRiskService this service receives User BCRA Information and check if user is candidate
// to a loan
type IUserRiskService interface {
	GetUserRisk(document uint64) (*domain.UserRisk, *apierror.ApiError)
}

type BaseUserRiskService struct {
	userDebtService           IUserDebtService
	userRiskCalculatorService IUserRiskCalculatorService
}

func NewUserRiskService(
	userDebtService IUserDebtService,
	userRiskCalculatorService IUserRiskCalculatorService,
) IUserRiskService {
	return &BaseUserRiskService{
		userDebtService:           userDebtService,
		userRiskCalculatorService: userRiskCalculatorService,
	}
}

func (b BaseUserRiskService) GetUserRisk(document uint64) (*domain.UserRisk, *apierror.ApiError) {
	userDebt, err := b.userDebtService.GetUserDebt(document)
	if err != nil {
		err := apierror.NewCustomErrorWithStatus(
			"BaseUserRiskService - error fetching userDebt from userDebtService",
			err.Err,
			err.Status,
		)
		return nil, err
	}

	userRisk, err := b.userRiskCalculatorService.GetCalculatedUserRisk(userDebt)
	if err != nil {
		err := apierror.NewCustomErrorWithStatus(
			"BaseUserRiskService - error fetching userRisk from userRiskCalculatorService",
			err.Err,
			err.Status,
		)
		return nil, err
	}

	return userRisk, nil
}
