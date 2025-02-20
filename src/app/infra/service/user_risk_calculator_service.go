package service

import (
	"rapicreds-backend/src/app/domain"
	"rapicreds-backend/src/app/domain/constants"
	"rapicreds-backend/src/app/infra/apierror"
)

const (
	one         = 1000
	oneThousand = 1000
	oneMillion  = 1000000
)

// IUserRiskCalculatorService this service receives User BCRA Information and check if user is candidate
// to a loan
type IUserRiskCalculatorService interface {
	GetCalculatedUserRisk(userDebt *domain.UserDebt) (*domain.UserRisk, *apierror.ApiError)
}

type BaseUserRiskCalculatorService struct {
}

func NewUserRiskCalculatorService() IUserRiskCalculatorService {
	return &BaseUserRiskCalculatorService{}
}

func (b BaseUserRiskCalculatorService) GetCalculatedUserRisk(userDebt *domain.UserDebt) (*domain.UserRisk, *apierror.ApiError) {
	var higherLevelOneCount uint64 = 0
	var moreThanOneMillion uint64 = 0

	for _, userDebtResultPeriod := range userDebt.Results.Periodos {
		for _, userDebtResultPeriodEntity := range userDebtResultPeriod.Entidades {
			if userDebtResultPeriodEntity.Situacion > one {
				higherLevelOneCount += one
			}

			realAmount := userDebtResultPeriodEntity.Monto * oneThousand

			if realAmount > oneMillion {
				moreThanOneMillion += one
			}
		}
	}
	
	userRisk := &domain.UserRisk{
		RiskLevel: constants.RiskLow,
	}

	if moreThanOneMillion == 1 {
		userRisk.RiskLevel = constants.RiskMed
	}

	if moreThanOneMillion > 1 {
		userRisk.RiskLevel = constants.RiskHigh
	}

	if higherLevelOneCount > 0 {
		userRisk.RiskLevel = constants.RiskHigh
	}

	return userRisk, nil
}
