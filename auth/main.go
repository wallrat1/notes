package main

import (
	"auth/internal/config"
	"auth/internal/server"
	"fmt"
	"log"
)

func main() {
	cfg := &config.Config{
		Port:         "8101",      // Порт, на котором будет запущен сервер
		Host:         "localhost", // Хост, на котором будет запущен сервер
		DBDSN:        "--",        // Строка подключения к базе данных, например, "postgres://user:password@localhost:5432/dbname"
		JWTSecretKey: "123",       // Секретный ключ для JWT токенов
	}
	server, err := server.NewServer(cfg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Сервер успешно создан\n")
	if err := server.Serve(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("сервер запущен")
}
