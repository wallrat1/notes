


# Используем официальный образ Go как базовый
FROM golang:1.24.2-alpine

# Устанавливаем рабочий каталог внутри контейнера
WORKDIR /app

# Копируем файлы 'go.mod' и, если есть, 'go.sum' в рабочий каталог
COPY go.mod ./
COPY go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем исходный код приложения в рабочий каталог
COPY . .

# Собираем приложение в исполняемый файл 'main'
RUN go build -o main .


# Указываем команду для запуска приложения
CMD ["./main"]