package config

import (
	"fmt"
	"os"
	"strconv"
)

// Config - структура для хранения конфигурации приложения
// Содержит параметры, необходимые для запуска сервера
// Параметры могут быть загружены из файла, переменных окружения или других источников
type Config struct {
	Port                   string // Порт, на котором будет запущен сервер
	Host                   string // Хост, на котором будет запущен сервер
	Timeout                int    // Таймаут для операций с сервером в секундах
	DBDSN                  string // Строка подключения к базе данных, например, "postgres://user:password@localhost:5432/dbname"
	DBSSL                  string
	DBTimeout              int    // Таймаут для операций с базой данных в секундах
	JWTSecretKey           string // Секретный ключ для JWT токенов
	AccessTokenExpiration  int    // Срок действия access токена в часах
	RefreshTokenExpiration int    // Срок действия refresh токена в часах
}

func NewConfig() *Config {
	port, err := getEnv("PORT")
	if err != nil {
		fmt.Println("Не удалось получить PORT из переменной окружения, используется порт по умолчанию")
	}
	host, err := getEnv("HOST")
	if err != nil {
		fmt.Println("Не удалось получить HOST из переменной окружения, используется порт по умолчанию")
	}
	timeout := 10
	if envValue, err := getEnv("SERVER_TIMEOUT"); err == nil {
		if parsed, parseErr := strconv.Atoi(envValue); parseErr == nil {
			timeout = parsed
		}
	} else {
		fmt.Println("Не удалось получить SERVER_TIMEOUT из переменной окружения, используется 10 секунд")
	}
	jwtSecretKey, err := getEnv("JWT_SECRET")
	if err != nil {
		fmt.Println("Не удалось получить JWT_SECRET из переменной окружения, используется порт по умолчанию")
	}
	accessTokenExpiration := 24
	if envValue, err := getEnv("ACCESS_TOKEN_EXPIRATION"); err == nil {
		if envValue, envParseError := strconv.Atoi(envValue); envParseError == nil {
			accessTokenExpiration = envValue
		}
	}
	refreshTokenExpiration := 24
	if envValue, err := getEnv("REFRESH_TOKEN_EXPIRATION"); err == nil {
		if envValue, envParseError := strconv.Atoi(envValue); envParseError == nil {
			refreshTokenExpiration = envValue
		}
	}
	dbHost, err := getEnv("POSTGRES_HOST")
	if err != nil {
		fmt.Println("Не удалось получить POSTGRES_HOST из переменной окружения")
	}
	dbTimeout := 10
	if envValue, err := getEnv("SERVER_TIMEOUT"); err == nil {
		if parsed, parseErr := strconv.Atoi(envValue); parseErr == nil {
			dbTimeout = parsed
		}
	}
	dbPort, err := getEnv("POSTGRES_PORT")
	if err != nil {
		fmt.Println("Не удалось получить POSTGRES_PORT из переменной окружения")
	}
	dbUser, err := getEnv("POSTGRES_USER")
	if err != nil {
		fmt.Println("Не удалось получить POSTGRES_USER из переменной окружения")
	}
	dbPassword, err := getEnv("POSTGRES_PASSWORD")
	if err != nil {
		fmt.Println("Не удалось получить POSTGRES_PASSWORD из переменной окружения")
	}
	dbName, err := getEnv("POSTGRES_DB")
	if err != nil {
		fmt.Println("Не удалось получить POSTGRES_DB из переменной окружения")
	}
	dbSSL, err := getEnv("POSTGRES_USE_SSL")
	if err != nil {
		fmt.Println("Не удалось получить POSTGRES_USE_SSL из переменной окружения")
	}
	// Формирование строки подключения к базе данных
	dbDSN := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		dbUser, dbPassword, dbHost, dbPort, dbName, dbSSL)
	return &Config{
		Port:                   port,
		Host:                   host,
		DBDSN:                  dbDSN, // Строка подключения к базе данных
		DBSSL:                  dbSSL,
		JWTSecretKey:           jwtSecretKey,
		AccessTokenExpiration:  accessTokenExpiration,
		RefreshTokenExpiration: refreshTokenExpiration,
		Timeout:                timeout,   // Таймаут для операций с сервером
		DBTimeout:              dbTimeout, // Таймаут для операций с базой данных
	}
}
func getEnv(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", fmt.Errorf("%s: %s", "не установлена переменная окружения", key)
	}
	return value, nil
}
