package service

import (
	"auth/internal/models"
	"context"
)

// Service - интерфейс для управления пользователями
// Он определяет методы согласно паттерна CRUD для создания, поиска, обновления и удаления пользователей
// Все методы принимают контекст для управления временем выполнения и отмены операций
// Это позволяет гибко управлять жизненным циклом операций с пользователями
type Service interface {
	Create(ctx context.Context, user *models.User) (*models.User, error)
	Read(ctx context.Context, id int) (*models.User, error)
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id int) error
	Authenticate(ctx context.Context, username, password string) (*models.User, error)
	ReadByUsername(ctx context.Context, username string) (*models.User, error)
	Close() error
}
