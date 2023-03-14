package main

import (
	"AbitService/app/models"
	"AbitService/router"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}
func main() {
	r := router.InitRouter()
	models.ConnectDatabase() // new
	err := r.Run()
	if err != nil {
		return
	}
}
