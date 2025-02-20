package config

import (
	"github.com/gin-gonic/gin"
	"rapicreds-backend/src/app/infra/controller"
	"rapicreds-backend/src/app/infra/restclient"
	"rapicreds-backend/src/app/infra/service"
)

func InjectDependencies() *gin.Engine {
	r := gin.Default()

	userDebtRestClient := restclient.NewUserDebtRestClient()
	userDebtService := service.NewUserDebtService(userDebtRestClient)
	userRiskCalculatorService := service.NewUserRiskCalculatorService()
	userRiskService := service.NewUserRiskService(userDebtService, userRiskCalculatorService)
	userRiskController := controller.NewBaseUserRiskController(userRiskService)

	// Ruta para /ping
	r.GET("/user/risk/:cuil", userRiskController.GetUserRisk)

	return r
}
