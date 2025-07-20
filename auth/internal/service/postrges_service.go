package service

import (
	"auth/internal/config"
	"auth/internal/database"
	"auth/internal/models"
	"context"
	"gorm.io/gorm"
)

type DBService struct {
	db *gorm.DB
}

// Проверка, что DBService реализует интерфейс Service
// Это позволяет гарантировать, что DBService соответствует контракту интерфейса Service
var _ Service = (*DBService)(nil)

// NewService - конструктор для создания нового экземпляра Service
// Он принимает конфигурацию и возвращает указатель на текущею реализацию
// сервиса или ошибку, если она произошла
func NewService(cfg *config.Config) (Service, error) {
	db, err := database.NewDatabase(cfg, &models.User{})
	if err != nil {
		return nil, err
	}
	return &DBService{db: db}, nil
}

// Create создает нового пользователя в базе данных
// Он принимает контекст и указатель на модель User
// Возвращает созданного пользователя или ошибку, если она произошла
func (p *DBService) Create(ctx context.Context, user *models.User) (*models.User, error) {
	if user == nil {
		return nil, gorm.ErrInvalidData
	}
	hashedpassword, err := models.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashedpassword
	if err := p.db.WithContext(ctx).Create(user).Error; err != nil {
		return nil, err
	}

	// Возвращаем созданного пользователя
	return user, nil
}

// Delete Удаляет пользователя из базы данных по ID
func (p *DBService) Delete(ctx context.Context, id int) error {
	if id <= 0 {
		return gorm.ErrRecordNotFound
	}
	if err := p.db.WithContext(ctx).Delete(&models.User{ID: id}).Error; err != nil {
		return err
	}
	return nil
}
func (p *DBService) Read(ctx context.Context, id int) (*models.User, error) {
	if id <= 0 {
		return nil, gorm.ErrRecordNotFound
	}
	var user models.User
	if err := p.db.WithContext(ctx).First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
func (p *DBService) Update(ctx context.Context, user *models.User) error {
	if user == nil || user.ID <= 0 {
		return gorm.ErrInvalidData
	}
	// Если пароль не пустой, значит он был изменен и нужно его хешировать
	if user.Password != "" {
		hashedPassword, err := models.HashPassword(user.Password)
		if err != nil {
			return err
		}
		user.Password = hashedPassword
	}
	// Используем контекст для управления временем выполнения операции
	if err := p.db.WithContext(ctx).Save(user).Error; err != nil {
		return err
	}

	return nil
}
func (p *DBService) ReadByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	if err := p.db.WithContext(ctx).Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil

}
func (p *DBService) Authenticate(ctx context.Context, username, password string) (*models.User, error) {
	if username == "" || password == "" {
		return nil, gorm.ErrInvalidData
	}
	user, err := p.ReadByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	if !models.CheckPassword(password, user.Password) {
		return nil, gorm.ErrRecordNotFound
	}
	return user, nil
}
func (p *DBService) Close() error {
	db, err := p.db.DB()
	if err != nil {
		return err
	}
	return db.Close()
}
