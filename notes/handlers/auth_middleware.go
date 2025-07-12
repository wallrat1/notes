package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"strings"
	"zametki/envs"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenValue := strings.Split(ctx.GetHeader("Authorization"), " ")
		if len(tokenValue) != 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Неверный формат заголовка Authorization"})
			return
		}
		accessToken, err := jwt.Parse(tokenValue[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return "", errors.New("неверный формат кодирования токина")
			}
			return []byte(envs.ServerEnvs.JWT_SECRET), nil
		})
		if err != nil || accessToken == nil || !accessToken.Valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "неверный токен"})
			return
		}
		ctx.Next()
	}
}
