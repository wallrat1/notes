package handlers

import (
	"auth/database"
	"auth/models"
	"auth/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func registerUserHandler(ctx *gin.Context) {
	var user models.User
	var registerData models.RegisterData
	if err := ctx.ShouldBindJSON(&registerData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Неверные данные": err.Error()})
		return
	}
	hashedPassword, err := utils.HashPassword(registerData.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"ошибка хэширования": err.Error()})
		return
	}
	user.Email = registerData.Email
	user.Hash = hashedPassword
	result := database.DB.Create(&user)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": "не удалось сохранить юзера"})
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": "юзер сохранен"})
	}
}
