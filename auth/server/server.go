package server

import (
	"auth/database"
	"auth/envs"
	"auth/models"
	"log"
)

func InitServer() {
	envs.LoadEnvs()
	errDb := database.InitDatabase()
	if errDb != nil {
		log.Fatal(errDb)
		return
	}
	database.DB.AutoMigrate(&models.User{})
}

func StartServer() {
	InitRotes()
	// Запуск сервера
}
