package server

import (
	"log"
	"zametki/database"
	"zametki/envs"
)

func InitServer() {
	// Инициализация внешних значений ENV
	envs.LoadEnvs()
	// Инициализация базы данных
	errDatabase := database.InitDatabase()
	if errDatabase != nil {
		log.Fatal("ошибка подключения базы данных", errDatabase)
	}
	errRedis := database.InitRedis()
	if errRedis != nil {
		log.Fatal("ошибка редиса ", errRedis)
	}
}

func StartServer() {
	InitRotes()
	// Запуск сервера
}
