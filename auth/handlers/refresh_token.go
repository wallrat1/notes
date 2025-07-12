package handlers

import (
	"auth/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken"`
}

// обновление и рефреш и асес токена в случае просрочки асес
func refreshToken(ctx *gin.Context) {
	var token RefreshTokenRequest
	if err := ctx.ShouldBindJSON(&token); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка при чтении запроса"})
		return
	}
	userId, err := utils.ValidateRefreshToken(token.RefreshToken)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Неверный refresh токен" + err.Error()})
		return
	}
	tokens, err := utils.GenerateTokens(userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "не удалось создать токены"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"tokens": tokens})

}
