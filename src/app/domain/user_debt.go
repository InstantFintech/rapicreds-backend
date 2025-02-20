package domain

type UserDebt struct {
	Status  int            `json:"status"`
	Results UserDebtResult `json:"results"`
}

type UserDebtResult struct {
	Identificacion int64                  `json:"identificacion"`
	Denominacion   string                 `json:"denominacion"`
	Periodos       []UserDebtResultPeriod `json:"periodos"`
}

type UserDebtResultPeriod struct {
	Periodo   string                     `json:"periodo"`
	Entidades []UserDebtResultPeriodData `json:"entidades"`
}

type UserDebtResultPeriodData struct {
	Entidad                 string  `json:"entidad"`
	Situacion               int     `json:"situacion"`
	FechaSit1               string  `json:"fechaSit1"`
	Monto                   float64 `json:"monto"`
	DiasAtrasoPago          int     `json:"diasAtrasoPago"`
	Refinanciaciones        bool    `json:"refinanciaciones"`
	RecategorizacionOblig   bool    `json:"recategorizacionOblig"`
	SituacionJuridica       bool    `json:"situacionJuridica"`
	IrrecDisposicionTecnica bool    `json:"irrecDisposicionTecnica"`
	EnRevision              bool    `json:"enRevision"`
	ProcesoJud              bool    `json:"procesoJud"`
}
