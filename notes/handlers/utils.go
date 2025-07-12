package handlers

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"strings"
	"zametki/envs"
)

func ExtractUserID(tokenString string) (uint, error) {
	str := strings.TrimSpace(strings.TrimPrefix(tokenString, "Bearer "))
	token, err := jwt.Parse(str, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("неожиданный алгоритм подписи: %v", token.Header["alg"])
		}
		return []byte(envs.ServerEnvs.JWT_SECRET), nil
	})
	if err != nil {
		return 0, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := claims["user_id"]

		if userIDFloat, ok := userID.(float64); ok {
			return uint(userIDFloat), nil // Преобразуем float64 в uint
		}
	}

	return 0, fmt.Errorf("невозможно извлечь user_id из токена")
}
