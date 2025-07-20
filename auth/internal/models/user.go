package models

import "golang.org/x/crypto/bcrypt"

// User представляет модель пользователя в системе
// Он содержит ID, имя пользователя и пароль
// Используется для хранения и управления данными пользователей
// Важно отметить, что пароль должен храниться в зашифрованном виде
type User struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique;not null"`
	Password string `json:"password,omitempty" gorm:"not null"`
}

const bcryptCost = 12

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	return string(bytes), err
}
func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
