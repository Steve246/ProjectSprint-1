package repository

import (
	"7Zero4/config"
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"math/rand"

	"time"

	"github.com/go-redis/redis/v8"
	conn "github.com/spf13/viper"
	"gorm.io/gorm"
)

type TokenRepository interface {
	FetchToken(key string) (string, error)

	VerifyTokenV2(tokenString string) (bool, error)
	CreateTokenV2(email string, length int) (string, error)
}

type tokenRepository struct {
	tokenConfig config.TokenConfig
	redisClient *redis.Client
	db          *gorm.DB
}

func (t *tokenRepository) CreateTokenV2(email string, length int) (string, error) {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	token := hex.EncodeToString(b)
	expire := conn.GetInt64("token_auth.expire")
	expireDate := time.Now().Add(time.Second * time.Duration(expire))
	result := t.db.Exec("INSERT INTO authentication(user_email,token_auth,expire) VALUES(?,?,?)", email, token, expireDate)
	if result.Error != nil {
		return "", result.Error
	}
	return token, nil
}

type UserData struct {
	Email  string `gorm:"column:user_email"`
	Expire string `gorm:"column:expire"`
}

func (t *tokenRepository) VerifyTokenV2(tokenString string) (bool, error) {
	var userData UserData

	fmt.Println("VerifyTokenV2")
	result := t.db.Raw(`SELECT user_email, expire FROM authentication WHERE token_auth = ? ORDER BY expire DESC LIMIT 1`, tokenString).Scan(&userData)
	fmt.Println(userData)
	if result.Error != nil {
		return false, result.Error
	}
	if result.RowsAffected == 0 {
		return false, errors.New("Unauthorized")
	}

	// check expire
	now := time.Now()
	date, _ := time.Parse(time.RFC3339, userData.Expire)
	if !now.Before(date) {
		return false, errors.New("Token Expired")
	}

	return true, nil
}

func (t *tokenRepository) FetchToken(key string) (string, error) {
	return t.redisClient.Get(context.Background(), key).Result()
}

func NewTokenRepository(tokenConfig config.TokenConfig, dbClient *gorm.DB) TokenRepository {
	repo := new(tokenRepository)
	repo.tokenConfig = tokenConfig
	repo.db = dbClient
	return repo
}
