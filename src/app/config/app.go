package config

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"rapicreds-backend/src/app/infra/controller"
	"rapicreds-backend/src/app/infra/repository"
	"rapicreds-backend/src/app/infra/restclient"
	"rapicreds-backend/src/app/infra/service"
)

func InjectDependencies() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		// Permitir solo solicitudes desde el frontend en localhost:5173 (ajusta el origen según necesites)
		AllowOrigins: []string{"http://localhost:5173"},
		// Permitir los métodos HTTP que quieras (GET, POST, PUT, DELETE)
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		// Permitir las cabeceras necesarias
		AllowHeaders: []string{"Content-Type", "Authorization"},
		// Habilitar credenciales si es necesario
		AllowCredentials: true,
	}))

	userDebtRestClient := restclient.NewUserDebtRestClient()
	userDebtService := service.NewUserDebtService(userDebtRestClient)
	userRiskCalculatorService := service.NewUserRiskCalculatorService()
	userRiskService := service.NewUserRiskService(userDebtService, userRiskCalculatorService)
	userRiskController := controller.NewBaseUserRiskController(userRiskService)

	// Ruta para /ping
	r.GET("/user/risk/:cuil", userRiskController.GetUserRisk)

	db := repository.InitDB()

	controller.InitDB(db)

	r.GET("/auth/valid-session", controller.IsAuth)

	r.POST("/auth/signup", controller.Signup)
	r.POST("/auth/login", controller.Login)

	// Rutas para Google OAuth
	r.GET("/auth/google", controller.GoogleLogin)
	r.GET("/auth/google/callback", controller.GoogleCallback)

	return r
}
