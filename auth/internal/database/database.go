package database

import (
	"auth/internal/config"
	"context"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

func NewDatabase(cfg *config.Config, models ...any) (*gorm.DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.Timeout)*time.Second)
	defer cancel()
	// Добавляем задержку в 1 секунду, // чтобы дать время на инициализацию других компонентов, если это необходимо
	// Это может быть полезно, если база данных запускается в контейнере или сервисе, который требует времени на инициализацию
	// Например, если база данных запускается в Docker-контейнере, то может потребоваться время на его запуск и готовность к соединению
	time.Sleep(1 * time.Second)
	db, err := gorm.Open(postgres.Open(cfg.DBDSN), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("%s: %v", "ошибка подключения к базе данных", err)
	}

	errMigration := runMigrations(db, models...) // Выполняем миграции
	if errMigration != nil {
		return nil, fmt.Errorf("%s: %v", "ошибка миграции базы данных", errMigration)
	}
	return db.WithContext(ctx), nil
}

func runMigrations(db *gorm.DB, models ...any) error {
	// Выполняем миграции для всех переданных моделей
	for _, model := range models {
		if err := db.AutoMigrate(model); err != nil {
			return fmt.Errorf("ошибка миграции модели %T: %w", model, err)
		}
	}
	// Если все миграции прошли успешно, возвращаем nil
	fmt.Println("Все миграции успешно выполнены")
	return nil
}
