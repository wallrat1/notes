package handlers

import (
	"auth/database"
	"auth/models"
	"auth/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getUser(ctx *gin.Context) {
	userId, err := utils.ExtractUserID(ctx.GetHeader("Authorization"))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный токен: " + err.Error()})
		return
	}
	var user models.User
	result := database.DB.Where("id = ?", userId).First(&user)
	if result.Error != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не найден " + err.Error()})
		return
	}
	userResponce := struct {
		ID    uint   `json:"id"`
		Email string `json:"email"`
	}{
		user.ID,
		user.Email,
	}
	ctx.JSON(http.StatusOK, userResponce)
}
