package main

import (
	"log"
	"rapicreds-backend/src/app/config"
)

func main() {
	r := config.InjectDependencies()

	err := r.Run(":8080")
	if err != nil {
		log.Fatal("Error:", err)
		return
	}
}
