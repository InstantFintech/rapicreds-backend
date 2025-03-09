package controller

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func ViewContractModel(c *gin.Context) {
	// Ruta al archivo PDF en el servidor
	filePath := "src/app/infra/controller/contract-model.pdf"

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the current working directory
	fmt.Println("Current working directory:", dir)
	// Abrir el archivo PDF

	// Establecer el encabezado Content-Type para el PDF
	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", "inline; filename=\"contract-model.pdf\"") // inline muestra el PDF en el navegador
	fmt.Println("aaa")

	// Devolver el archivo PDF en la respuesta
	c.File(filePath)
	fmt.Println("bbb")

}
