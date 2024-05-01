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
	"github.com/golang-jwt/jwt"
	conn "github.com/spf13/viper"
	"gorm.io/gorm"
)

type TokenRepository interface {
	CreateToken(makeClaims func(string) jwt.Claims) (string, error)
	VerifyToken(tokenString string) (jwt.MapClaims, error)
	StoreToken(key string, value interface{}, expTime time.Duration) error
	StoreTokenForever(key string, value interface{}) error
	FetchToken(key string) (string, error)

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

// StoreTokenForever implements TokenRepository
func (t *tokenRepository) StoreTokenForever(key string, value interface{}) error {
	result := t.redisClient.Set(context.Background(), key, value, 0)
	return result.Err()
}

func (t *tokenRepository) CreateToken(makeClaims func(string) jwt.Claims) (string, error) {
	claims := makeClaims(t.tokenConfig.ApplicationName)
	token := jwt.NewWithClaims(
		t.tokenConfig.JwtSigningMethod,
		claims,
	)

	newToken, err := token.SignedString([]byte(t.tokenConfig.JwtSignatureKey))
	if err != nil {
		return "", err
	}
	return newToken, nil
}

func (t *tokenRepository) VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		method, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("signing method invalid")
		} else if method != t.tokenConfig.JwtSigningMethod {
			return nil, fmt.Errorf("signing method invalid")
		}
		return []byte(t.tokenConfig.JwtSignatureKey), nil
	})
	if err != nil {
		fmt.Println("Parsing failed..")
		return nil, errors.New("tokenExpired")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid || claims["iss"] != t.tokenConfig.ApplicationName {
		fmt.Println("Token invalid..")
		return nil, err
	}
	return claims, nil
}

func (t *tokenRepository) StoreToken(key string, value interface{}, expTime time.Duration) error {
	result := t.redisClient.Set(context.Background(), key, value, expTime)
	return result.Err()
}

func (t *tokenRepository) FetchToken(key string) (string, error) {
	return t.redisClient.Get(context.Background(), key).Result()
}

func NewTokenRepository(redisClient *redis.Client, tokenConfig config.TokenConfig, dbClient *gorm.DB) TokenRepository {
	repo := new(tokenRepository)
	repo.redisClient = redisClient
	repo.tokenConfig = tokenConfig
	repo.db = dbClient
	return repo
}
