package database

import (
	"auth/envs"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB

func InitDatabase() error {
	env := envs.ServerEnvs
	uri := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", env.POSTGRES_HOST, env.POSTGRES_USER, env.POSTGRES_PASSWORD, env.POSTGRES_DB, env.POSTGRES_PORT, env.POSTGRES_USE_SSL)
	db, err := gorm.Open(postgres.Open(uri))
	for i := 0; i < 5; i++ {
		fmt.Println(uri)
		if err == nil {
			break
		}
		time.Sleep(1 * time.Second)
		db, err = gorm.Open(postgres.Open(uri))
	}
	if err != nil {
		return err
	}
	DB = db
	return nil
}
