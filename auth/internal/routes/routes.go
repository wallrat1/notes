package routes

import (
	"auth/internal/handler"
	"github.com/gin-gonic/gin"
)

// SetupRouter настраивает все маршруты приложения
// Принимает обработчик, который содержит логику для работы с пользователями
// Возвращает настроенный роутер
func SetupRoutes(h *handler.Handler) *gin.Engine {
	router := gin.Default()
	auth := router.Group("/auth")
	{
		auth.POST("/register", h.RegisterUser)
		auth.POST("/login", h.LoginUser)
		// Защищенные endpoints
		auth.GET("/user", h.GetUserInfo)
		auth.PUT("/user", h.UpdateUser)
		auth.DELETE("/user", h.DeleteUser)
	}
	return router

}
