package domain

import "rapicreds-backend/src/app/domain/constants"

type UserRisk struct {
	RiskLevel constants.RiskLevel `json:"risk_level"`
}
