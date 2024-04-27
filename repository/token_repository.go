package repository

import (
	"7Zero4/config"
	"context"
	"errors"
	"fmt"

	"time"

	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt"
)

type TokenRepository interface {
	CreateToken(makeClaims func(string) jwt.Claims) (string, error)
	VerifyToken(tokenString string) (jwt.MapClaims, error)
	StoreToken(key string, value interface{}, expTime time.Duration) error
	StoreTokenForever(key string, value interface{}) error
	FetchToken(key string) (string, error)
}

type tokenRepository struct {
	tokenConfig config.TokenConfig
	redisClient *redis.Client
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

// func StoreAccessToken(userName string, tokenDetail *model.TokenDetails) error {
// 	at := time.Unix(tokenDetail.AtExpires, 0)
// 	now := time.Now()
// 	err := config.TokenConfig.Client.Set(context.Background(), tokenDetail.AccessUuid, userName, at.Sub(now)).Err()
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
// func FetchAccessToken(accessDetail *model.AccessDetail) (string, error) {
// 	if accessDetail != nil {
// 		result, err := config.TokenConfig.Client.Get(context.Background(), accessDetail.AccessUiid).Result()
// 		if err != nil {
// 			return "", err
// 		}
// 		return result, nil
// 	} else {
// 		return "", errors.New("invalid access")
// 	}
// }

func NewTokenRepository(redisClient *redis.Client, tokenConfig config.TokenConfig) TokenRepository {
	repo := new(tokenRepository)
	repo.redisClient = redisClient
	repo.tokenConfig = tokenConfig
	return repo
}
