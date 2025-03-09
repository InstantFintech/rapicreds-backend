package domain

type UserLoan struct {
	ID                uint    `json:"id" gorm:"primaryKey"`
	UserID            string  `json:"user_id"`
	LoanAmount        float64 `json:"loan_amount"`
	LoanInstallments  float64 `json:"loan_installments"`
	LoanPaymentAmount string  `json:"loan_payment_amount"`
	LoanPaymentDate   string  `json:"loan_payment_date"`
	LoanRequestedDate string  `json:"loan_requested_date"`
	Status            string  `json:"status"`
}
