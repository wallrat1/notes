package handler

import (
	"auth/internal/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Handler содержит все обработчики для работы с пользователями
type Handler struct {
	cfg *config.Config // Конфигурация сервера
}

// NewHandler создает новый экземпляр обработчика пользователей
func NewHandler(cfg *config.Config) *Handler {
	return &Handler{
		cfg: cfg, // Сохраняем конфигурацию в обработчике
	}
}

func (h *Handler) RegisterUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"key": "RegisterUser"})
}

func (h *Handler) LoginUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"key": "LoginUser"})
}

// GetUserInfo обрабатывает запрос на получение информации о пользователе
func (h *Handler) GetUserInfo(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Информация о пользователе успешно получена",
	})
}

// UpdateUser обрабатывает запрос на обновление данных пользователя
func (h *Handler) UpdateUser(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "Данные пользователя успешно обновлены",
	})
}

// DeleteUser обрабатывает запрос на удаление пользователя
func (h *Handler) DeleteUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Пользователь успешно удален",
	})
}
