package handlers

import (
	"auth/database"
	"auth/models"
	"auth/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func signIn(ctx *gin.Context) {
	var registerData models.RegisterData
	if err := ctx.ShouldBind(&registerData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"неверные данные": err.Error()})
		return
	}
	var user models.User
	result := database.DB.Where("email = ?", registerData.Email).First(&user)
	if result.Error != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "неверный email"})
		return
	}
	if !utils.CheckPasswordHash(registerData.Password, user.Hash) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "неверный пароль"})
		return
	}
	tokens, err := utils.GenerateTokens(user.ID)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"не удалось создать токен ": err.Error()})
		return
	}
	// Создаем анонимную структуру с только id и email
	userResponse := struct {
		ID    uint   `json:"id"`
		Email string `json:"email"`
	}{
		ID:    user.ID,
		Email: user.Email,
	}
	ctx.JSON(http.StatusOK, gin.H{"tokens": tokens, "user": userResponse})
}
